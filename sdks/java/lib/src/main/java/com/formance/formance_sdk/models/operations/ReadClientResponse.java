/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class ReadClientResponse {
    
    public String contentType;
    public ReadClientResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * Retrieved client
     */
    
    public com.formance.formance_sdk.models.shared.ReadClientResponse readClientResponse;
    public ReadClientResponse withReadClientResponse(com.formance.formance_sdk.models.shared.ReadClientResponse readClientResponse) {
        this.readClientResponse = readClientResponse;
        return this;
    }
    
    
    public Integer statusCode;
    public ReadClientResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public ReadClientResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}
