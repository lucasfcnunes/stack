/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class GetInstanceResponse {
    public String contentType;
    public GetInstanceResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * General error
     */
    public com.formance.formance_sdk.models.shared.Error error;
    public GetInstanceResponse withError(com.formance.formance_sdk.models.shared.Error error) {
        this.error = error;
        return this;
    }
    
    /**
     * The workflow instance
     */
    public com.formance.formance_sdk.models.shared.GetWorkflowInstanceResponse getWorkflowInstanceResponse;
    public GetInstanceResponse withGetWorkflowInstanceResponse(com.formance.formance_sdk.models.shared.GetWorkflowInstanceResponse getWorkflowInstanceResponse) {
        this.getWorkflowInstanceResponse = getWorkflowInstanceResponse;
        return this;
    }
    
    public Integer statusCode;
    public GetInstanceResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    public HttpResponse<byte[]> rawResponse;
    public GetInstanceResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}
