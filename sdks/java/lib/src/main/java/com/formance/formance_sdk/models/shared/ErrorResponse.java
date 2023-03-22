/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * ErrorResponse - Error
 */
public class ErrorResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("details")public String details;
    public ErrorResponse withDetails(String details) {
        this.details = details;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("errorCode")public ErrorsEnumEnum errorCode;
    public ErrorResponse withErrorCode(ErrorsEnumEnum errorCode) {
        this.errorCode = errorCode;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("errorMessage")public String errorMessage;
    public ErrorResponse withErrorMessage(String errorMessage) {
        this.errorMessage = errorMessage;
        return this;
    }
    
}
