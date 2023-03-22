// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/sdks/go/pkg/models/shared"
	"net/http"
	"time"
)

type ListTransactionsRequest struct {
	// Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $).
	Account *string `queryParam:"style=form,explode=true,name=account"`
	// Pagination cursor, will return transactions after given txid (in descending order).
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Filter transactions with postings involving given account at destination (regular expression placed between ^ and $).
	Destination *string `queryParam:"style=form,explode=true,name=destination"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	//
	EndTime *time.Time `queryParam:"style=form,explode=true,name=endTime"`
	// Filter transactions that occurred before this timestamp.
	// The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
	// Deprecated, please use `endTime` instead.
	//
	EndTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=end_time"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata map[string]interface{} `queryParam:"style=deepObject,explode=true,name=metadata"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
	// The maximum number of results to return per page.
	// Deprecated, please use `pageSize` instead.
	//
	PageSizeDeprecated *int64 `queryParam:"style=form,explode=true,name=page_size"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	// Deprecated, please use `cursor` instead.
	//
	PaginationToken *string `queryParam:"style=form,explode=true,name=pagination_token"`
	// Find transactions by reference field.
	Reference *string `queryParam:"style=form,explode=true,name=reference"`
	// Filter transactions with postings involving given account at source (regular expression placed between ^ and $).
	Source *string `queryParam:"style=form,explode=true,name=source"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	//
	StartTime *time.Time `queryParam:"style=form,explode=true,name=startTime"`
	// Filter transactions that occurred after this timestamp.
	// The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
	// Deprecated, please use `startTime` instead.
	//
	StartTimeDeprecated *time.Time `queryParam:"style=form,explode=true,name=start_time"`
}

type ListTransactionsResponse struct {
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	StatusCode    int
	RawResponse   *http.Response
	// OK
	TransactionsCursorResponse *shared.TransactionsCursorResponse
}
