/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class ListPaymentsResponse {
    
    public String contentType;
    public ListPaymentsResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * OK
     */
    
    public com.formance.formance_sdk.models.shared.PaymentsCursor paymentsCursor;
    public ListPaymentsResponse withPaymentsCursor(com.formance.formance_sdk.models.shared.PaymentsCursor paymentsCursor) {
        this.paymentsCursor = paymentsCursor;
        return this;
    }
    
    
    public Integer statusCode;
    public ListPaymentsResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public ListPaymentsResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    

    public ListPaymentsResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
    this.contentType = contentType;
this.statusCode = statusCode;
  }
}
