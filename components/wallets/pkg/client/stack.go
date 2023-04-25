package client

import (
	"context"
	"net/http"

	sdk "github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/stack/libs/go-libs/otlp"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func GetAuthenticatedClient(ctx context.Context, clientID, clientSecret, stackURL string, debug bool) (*http.Client, error) {
	if clientID == "" || clientSecret == "" {
		return nil, errors.New("STACK_CLIENT_ID and STACK_CLIENT_SECRET must be set")
	}

	clientCredentialsConfig := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     stackURL + "/api/auth/oauth/token",
	}
	underlyingHTTPClient := &http.Client{
		Transport: otlp.NewRoundTripper(debug),
	}

	return clientCredentialsConfig.Client(context.WithValue(ctx, oauth2.HTTPClient, underlyingHTTPClient)), nil
}

func NewStackClient(clientID, clientSecret, stackURL string, debug bool) (*sdk.Formance, error) {

	httpClient, err := GetAuthenticatedClient(context.Background(), clientID, clientSecret, stackURL, debug)
	if err != nil {
		return nil, err
	}
	config := sdk.New(
		sdk.WithServerURL(stackURL),
		sdk.WithClient(httpClient),
	)

	return config, nil
}
