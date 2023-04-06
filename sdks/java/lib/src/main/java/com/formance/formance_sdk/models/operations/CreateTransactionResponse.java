/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class CreateTransactionResponse {
    
    public String contentType;
    public CreateTransactionResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * Error
     */
    
    public com.formance.formance_sdk.models.shared.ErrorResponse errorResponse;
    public CreateTransactionResponse withErrorResponse(com.formance.formance_sdk.models.shared.ErrorResponse errorResponse) {
        this.errorResponse = errorResponse;
        return this;
    }
    
    
    public Integer statusCode;
    public CreateTransactionResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public CreateTransactionResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    /**
     * OK
     */
    
    public com.formance.formance_sdk.models.shared.TransactionsResponse transactionsResponse;
    public CreateTransactionResponse withTransactionsResponse(com.formance.formance_sdk.models.shared.TransactionsResponse transactionsResponse) {
        this.transactionsResponse = transactionsResponse;
        return this;
    }
    
}
