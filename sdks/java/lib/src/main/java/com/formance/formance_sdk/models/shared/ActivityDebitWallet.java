/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class ActivityDebitWallet {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("data")
    public DebitWalletRequest data;
    public ActivityDebitWallet withData(DebitWalletRequest data) {
        this.data = data;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("id")
    public String id;
    public ActivityDebitWallet withId(String id) {
        this.id = id;
        return this;
    }
    

    public ActivityDebitWallet(){}
}
