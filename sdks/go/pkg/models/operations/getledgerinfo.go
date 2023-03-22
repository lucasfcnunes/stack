// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/sdks/go/pkg/models/shared"
	"net/http"
)

type GetLedgerInfoRequest struct {
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
}

type GetLedgerInfoResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	// OK
	LedgerInfoResponse interface{}
	StatusCode         int
	RawResponse        *http.Response
}
