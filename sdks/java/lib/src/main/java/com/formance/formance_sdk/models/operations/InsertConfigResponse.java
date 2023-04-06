/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class InsertConfigResponse {
    /**
     * Config created successfully.
     */
    
    public com.formance.formance_sdk.models.shared.ConfigResponse configResponse;
    public InsertConfigResponse withConfigResponse(com.formance.formance_sdk.models.shared.ConfigResponse configResponse) {
        this.configResponse = configResponse;
        return this;
    }
    
    
    public String contentType;
    public InsertConfigResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    
    public Integer statusCode;
    public InsertConfigResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public InsertConfigResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    /**
     * Bad Request
     */
    
    public String insertConfig400TextPlainString;
    public InsertConfigResponse withInsertConfig400TextPlainString(String insertConfig400TextPlainString) {
        this.insertConfig400TextPlainString = insertConfig400TextPlainString;
        return this;
    }
    
}
