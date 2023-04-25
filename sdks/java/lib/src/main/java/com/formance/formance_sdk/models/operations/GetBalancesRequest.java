/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class GetBalancesRequest {
    /**
     * Filter balances involving given account, either as source or destination.
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=address")
    public String address;
    public GetBalancesRequest withAddress(String address) {
        this.address = address;
        return this;
    }
    
    /**
     * Pagination cursor, will return accounts after given address, in descending order.
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=after")
    public String after;
    public GetBalancesRequest withAfter(String after) {
        this.after = after;
        return this;
    }
    
    /**
     * Parameter used in pagination requests. Maximum page size is set to 15.
     * Set to the value of next for the next page of results.
     * Set to the value of previous for the previous page of results.
     * No other parameters can be set when this parameter is set.
     * 
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=cursor")
    public String cursor;
    public GetBalancesRequest withCursor(String cursor) {
        this.cursor = cursor;
        return this;
    }
    
    /**
     * Name of the ledger.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=ledger")
    public String ledger;
    public GetBalancesRequest withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    
    /**
     * Parameter used in pagination requests.
     * Set to the value of next for the next page of results.
     * Set to the value of previous for the previous page of results.
     * Deprecated, please use `cursor` instead.
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=pagination_token")
    public String paginationToken;
    public GetBalancesRequest withPaginationToken(String paginationToken) {
        this.paginationToken = paginationToken;
        return this;
    }
    

    public GetBalancesRequest(@JsonProperty("ledger") String ledger) {
    this.ledger = ledger;
  }
}
