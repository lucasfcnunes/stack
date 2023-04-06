/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class CreateScopeResponse {
    
    public String contentType;
    public CreateScopeResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * Created scope
     */
    
    public com.formance.formance_sdk.models.shared.CreateScopeResponse createScopeResponse;
    public CreateScopeResponse withCreateScopeResponse(com.formance.formance_sdk.models.shared.CreateScopeResponse createScopeResponse) {
        this.createScopeResponse = createScopeResponse;
        return this;
    }
    
    
    public Integer statusCode;
    public CreateScopeResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public CreateScopeResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}
