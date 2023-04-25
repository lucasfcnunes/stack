package client

import (
	sdk "github.com/formancehq/formance-sdk-go"
	"go.uber.org/fx"
)

func NewModule(clientID string, clientSecret string, tokenURL string, debug bool) fx.Option {
	return fx.Provide(func() (*sdk.Formance, error) {
		return NewStackClient(clientID, clientSecret, tokenURL, debug)
	})
}
