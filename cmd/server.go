package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/formancehq/go-libs/oauth2/oauth2introspect"
	"github.com/formancehq/go-libs/sharedauth"
	sharedhealth "github.com/formancehq/go-libs/sharedhealth/pkg"
	"github.com/formancehq/go-libs/sharedlogging"
	"github.com/formancehq/go-libs/sharedotlp/pkg/sharedotlptraces"
	"github.com/formancehq/search/pkg/searchengine"
	"github.com/formancehq/search/pkg/searchhttp"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/opensearch-project/opensearch-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

const (
	openSearchServiceFlag = "open-search-service"
	openSearchSchemeFlag  = "open-search-scheme"
	esIndicesFlag         = "es-indices"
	bindFlag              = "bind"

	authBasicEnabledFlag        = "auth-basic-enabled"
	authBasicCredentialsFlag    = "auth-basic-credentials"
	authBearerEnabledFlag       = "auth-bearer-enabled"
	authBearerIntrospectUrlFlag = "auth-bearer-introspect-url"
	authBearerAudienceFlag      = "auth-bearer-audience"

	defaultBind = ":8080"

	healthCheckPath = "/_healthcheck"
)

func NewServer() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "serve",
		Short:        "Launch the search server",
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			debug := viper.GetBool(debugFlag)

			logger := logrus.New()
			logger.SetFormatter(&logrus.JSONFormatter{})
			if debug {
				logger.SetLevel(logrus.DebugLevel)
				logger.Debugln("Debug mode enabled")
			}
			logger.Debugf("Starting with config: %s", viper.AllSettings())

			openSearchServiceHost := viper.GetString(openSearchServiceFlag)
			if openSearchServiceHost == "" {
				exitWithError(logger, "missing open search service host")
			}

			esIndices := viper.GetStringSlice(esIndicesFlag)
			if len(esIndices) == 0 {
				esIndices = searchengine.DefaultEsIndices
			}

			bind := viper.GetString(bindFlag)
			if bind == "" {
				bind = defaultBind
			}

			options := make([]fx.Option, 0)
			options = append(options, opensearchClientModule(openSearchServiceHost, esIndices...))
			options = append(options,
				sharedhealth.Module(),
				sharedhealth.ProvideHealthCheck(func(client *opensearch.Client) sharedhealth.NamedCheck {
					return sharedhealth.NewNamedCheck("elasticsearch connection", sharedhealth.CheckFn(func(ctx context.Context) error {
						_, err := client.Ping()
						return err
					}))
				}),
			)

			options = append(options, sharedotlptraces.CLITracesModule(viper.GetViper()))
			options = append(options, apiModule("search", bind, esIndices...))

			app := fx.New(options...)

			err := app.Start(cmd.Context())
			if err != nil {
				return err
			}

			<-app.Done()

			return nil
		},
	}

	cmd.Flags().String(bindFlag, defaultBind, "http server address")
	cmd.Flags().StringSlice(esIndicesFlag, searchengine.DefaultEsIndices, "ES indices to look")
	cmd.Flags().String(openSearchServiceFlag, "", "Open search service hostname")
	cmd.Flags().String(openSearchSchemeFlag, "https", "OpenSearch scheme")
	cmd.Flags().Bool(authBasicEnabledFlag, false, "Enable basic auth")
	cmd.Flags().StringSlice(authBasicCredentialsFlag, []string{}, "HTTP basic auth credentials (<username>:<password>)")
	cmd.Flags().Bool(authBearerEnabledFlag, false, "Enable bearer auth")
	cmd.Flags().String(authBearerIntrospectUrlFlag, "", "OAuth2 introspect URL")
	cmd.Flags().String(authBearerAudienceFlag, "", "OAuth2 audience template")
	sharedotlptraces.InitOTLPTracesFlags(cmd.Flags())

	return cmd
}

func exitWithError(logger *logrus.Logger, msg string) {
	logger.Error(msg)
	os.Exit(1)
}

func opensearchClientModule(openSearchServiceHost string, esIndices ...string) fx.Option {
	return fx.Options(
		fx.Provide(func() (*opensearch.Client, error) {
			httpTransport := http.DefaultTransport
			httpTransport.(*http.Transport).TLSClientConfig = &tls.Config{
				InsecureSkipVerify: true,
			}

			return opensearch.NewClient(opensearch.Config{
				Addresses: []string{viper.GetString(openSearchSchemeFlag) + "://" + openSearchServiceHost},
				Transport: otelhttp.NewTransport(httpTransport),
			})
		}),
		fx.Invoke(func(lc fx.Lifecycle, client *opensearch.Client) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return searchengine.LoadDefaultMapping(context.TODO(), client, esIndices...)
				},
			})
		}),
	)
}

func apiModule(serviceName, bind string, esIndices ...string) fx.Option {
	return fx.Options(
		fx.Provide(fx.Annotate(func(openSearchClient *opensearch.Client, tp trace.TracerProvider, healthController *sharedhealth.HealthController) (http.Handler, error) {
			router := mux.NewRouter()

			router.Use(handlers.RecoveryHandler())
			router.Handle(healthCheckPath, http.HandlerFunc(healthController.Check))

			routerWithTraces := router.PathPrefix("/").Subrouter()
			if viper.GetBool(sharedotlptraces.OtelTracesFlag) {
				routerWithTraces.Use(otelmux.Middleware(serviceName, otelmux.WithTracerProvider(tp)))
			}

			protected := routerWithTraces.PathPrefix("/").Subrouter()

			methods := make([]sharedauth.Method, 0)
			if viper.GetBool(authBasicEnabledFlag) {
				credentials := sharedauth.Credentials{}
				for _, kv := range viper.GetStringSlice(authBasicCredentialsFlag) {
					parts := strings.SplitN(kv, ":", 2)
					credentials[parts[0]] = sharedauth.Credential{
						Password: parts[1],
						Scopes:   []string{"search"},
					}
				}
				methods = append(methods, sharedauth.NewHTTPBasicMethod(credentials))
			}
			if viper.GetBool(authBearerEnabledFlag) {
				methods = append(methods, sharedauth.NewHttpBearerMethod(
					sharedauth.NewIntrospectionValidator(
						oauth2introspect.NewIntrospecter(viper.GetString(authBearerIntrospectUrlFlag)),
						false,
						sharedauth.AudienceIn(viper.GetString(authBearerAudienceFlag)),
					),
				))
			}

			if len(methods) > 0 {
				protected.Use(sharedauth.Middleware(methods...))
			}
			routerWithTraces.PathPrefix("/").Handler(searchhttp.Handler(searchengine.NewDefaultEngine(
				openSearchClient,
				searchengine.WithESIndices(esIndices...),
			)))

			return router, nil
		}, fx.ParamTags(``, `optional:"true"`))),
		fx.Invoke(func(lc fx.Lifecycle, handler http.Handler) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					sharedlogging.GetLogger(ctx).Infof("Starting http server on %s", bind)
					go func() {
						err := http.ListenAndServe(bind, handler)
						if err != nil {
							fmt.Fprintln(os.Stderr, err)
							os.Exit(1)
						}
					}()
					return nil
				},
			})
		}),
	)
}
