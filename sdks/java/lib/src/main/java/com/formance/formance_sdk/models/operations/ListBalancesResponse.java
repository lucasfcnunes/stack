/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class ListBalancesResponse {
    
    public String contentType;
    public ListBalancesResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * Balances list
     */
    
    public com.formance.formance_sdk.models.shared.ListBalancesResponse listBalancesResponse;
    public ListBalancesResponse withListBalancesResponse(com.formance.formance_sdk.models.shared.ListBalancesResponse listBalancesResponse) {
        this.listBalancesResponse = listBalancesResponse;
        return this;
    }
    
    
    public Integer statusCode;
    public ListBalancesResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public ListBalancesResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    

    public ListBalancesResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
    this.contentType = contentType;
this.statusCode = statusCode;
  }
}
