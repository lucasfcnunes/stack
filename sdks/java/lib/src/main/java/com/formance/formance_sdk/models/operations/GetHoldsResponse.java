/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class GetHoldsResponse {
    public String contentType;
    public GetHoldsResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * Holds
     */
    public com.formance.formance_sdk.models.shared.GetHoldsResponse getHoldsResponse;
    public GetHoldsResponse withGetHoldsResponse(com.formance.formance_sdk.models.shared.GetHoldsResponse getHoldsResponse) {
        this.getHoldsResponse = getHoldsResponse;
        return this;
    }
    
    public Integer statusCode;
    public GetHoldsResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    public HttpResponse<byte[]> rawResponse;
    public GetHoldsResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    /**
     * Error
     */
    public com.formance.formance_sdk.models.shared.WalletsErrorResponse walletsErrorResponse;
    public GetHoldsResponse withWalletsErrorResponse(com.formance.formance_sdk.models.shared.WalletsErrorResponse walletsErrorResponse) {
        this.walletsErrorResponse = walletsErrorResponse;
        return this;
    }
    
}
