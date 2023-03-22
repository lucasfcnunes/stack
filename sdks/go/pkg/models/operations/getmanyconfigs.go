// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type GetManyConfigsRequest struct {
	// Optional filter by endpoint URL
	Endpoint *string `queryParam:"style=form,explode=true,name=endpoint"`
	// Optional filter by Config ID
	ID *string `queryParam:"style=form,explode=true,name=id"`
}

type GetManyConfigsResponse struct {
	// OK
	ConfigsResponse *shared.ConfigsResponse
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
}
