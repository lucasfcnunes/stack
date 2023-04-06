/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class ActivityRevertTransaction {
    @JsonProperty("id")
    public String id;
    public ActivityRevertTransaction withId(String id) {
        this.id = id;
        return this;
    }
    
    @JsonProperty("ledger")
    public String ledger;
    public ActivityRevertTransaction withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    
}
