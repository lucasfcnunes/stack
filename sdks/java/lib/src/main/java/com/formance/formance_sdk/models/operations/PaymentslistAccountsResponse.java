/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class PaymentslistAccountsResponse {
    /**
     * OK
     */
    
    public com.formance.formance_sdk.models.shared.AccountsCursor accountsCursor;
    public PaymentslistAccountsResponse withAccountsCursor(com.formance.formance_sdk.models.shared.AccountsCursor accountsCursor) {
        this.accountsCursor = accountsCursor;
        return this;
    }
    
    
    public String contentType;
    public PaymentslistAccountsResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    
    public Integer statusCode;
    public PaymentslistAccountsResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public PaymentslistAccountsResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
}
