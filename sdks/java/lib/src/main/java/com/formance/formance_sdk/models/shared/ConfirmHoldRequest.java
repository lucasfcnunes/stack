/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class ConfirmHoldRequest {
    /**
     * Define the amount to transfer.
     */
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("amount")
    public Long amount;
    public ConfirmHoldRequest withAmount(Long amount) {
        this.amount = amount;
        return this;
    }
    
    /**
     * Define a final confirmation. Remaining funds will be returned to the wallet.
     */
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("final")
    public Boolean final_;
    public ConfirmHoldRequest withFinal(Boolean final_) {
        this.final_ = final_;
        return this;
    }
    
}
