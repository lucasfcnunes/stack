/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * ScriptResponse - On success, it will return a 200 status code, and the resulting transaction under the `transaction` field.
 * 
 * On failure, it will also return a 200 status code, and the following fields:
 *   - `details`: contains a URL. When there is an error parsing Numscript, the result can be difficult to read—the provided URL will render the error in an easy-to-read format.
 *   - `errorCode` and `error_code` (deprecated): contains the string code of the error
 *   - `errorMessage` and `error_message` (deprecated): contains a human-readable indication of what went wrong, for example that an account had insufficient funds, or that there was an error in the provided Numscript.
 * 
 */
public class ScriptResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("details")public String details;
    public ScriptResponse withDetails(String details) {
        this.details = details;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("errorCode")public ErrorsEnumEnum errorCode;
    public ScriptResponse withErrorCode(ErrorsEnumEnum errorCode) {
        this.errorCode = errorCode;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("errorMessage")public String errorMessage;
    public ScriptResponse withErrorMessage(String errorMessage) {
        this.errorMessage = errorMessage;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("transaction")public Transaction transaction;
    public ScriptResponse withTransaction(Transaction transaction) {
        this.transaction = transaction;
        return this;
    }
    
}
