// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type ListConfigsAvailableConnectorsResponse struct {
	// OK
	ConnectorsConfigsResponse *shared.ConnectorsConfigsResponse
	ContentType               string
	StatusCode                int
	RawResponse               *http.Response
}
