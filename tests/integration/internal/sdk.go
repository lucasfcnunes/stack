package internal

import (
	"net/http"
	"net/url"

	"github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/stack/libs/go-libs/httpclient"
)

var sdkClient *formance.Formance

func configureSDK() {
	gatewayUrl, err := url.Parse(gatewayServer.URL)
	if err != nil {
		panic(err)
	}

	sdkClient = formance.New(
		formance.WithClient(&http.Client{
			Transport: httpclient.NewDebugHTTPTransport(http.DefaultTransport),
		}),
		formance.WithServerURL(gatewayUrl.Host),
	)
}

func Client() *formance.Formance {
	return sdkClient
}
