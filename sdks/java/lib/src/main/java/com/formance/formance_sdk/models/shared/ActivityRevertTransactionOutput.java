/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ActivityRevertTransactionOutput {
    @JsonProperty("data")
    public Transaction data;
    public ActivityRevertTransactionOutput withData(Transaction data) {
        this.data = data;
        return this;
    }
    
}
