module github.com/formancehq/auth

go 1.19

require (
	github.com/formancehq/stack/libs/go-libs v0.0.0-20230221161632-e6dc6a89a85e
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/oauth2-proxy/mockoidc v0.0.0-20220308204021-b9169deeb282
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.1
	github.com/zitadel/logging v0.3.4
	github.com/zitadel/oidc v1.13.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.39.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.39.0
	go.opentelemetry.io/otel v1.13.0
	go.opentelemetry.io/otel/trace v1.13.0
	go.uber.org/fx v1.19.1
	golang.org/x/oauth2 v0.5.0
	golang.org/x/text v0.14.0
	gopkg.in/square/go-jose.v2 v2.6.0
	gorm.io/driver/postgres v1.4.8
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.24.5
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.4 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/uptrace/opentelemetry-go-extra/otellogrus v0.1.21 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.1.21 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.14.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.13.0 // indirect
	go.opentelemetry.io/otel/metric v0.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.13.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230221151758-ace64dc21148 // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/formancehq/stack/libs/go-libs => ../../libs/go-libs

replace github.com/zitadel/oidc => github.com/formancehq/oidc v0.0.0-20220923202448-e2960a99b71c
