/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class WalletSubject {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("balance")
    public String balance;
    public WalletSubject withBalance(String balance) {
        this.balance = balance;
        return this;
    }
    
    @JsonProperty("identifier")
    public String identifier;
    public WalletSubject withIdentifier(String identifier) {
        this.identifier = identifier;
        return this;
    }
    
    @JsonProperty("type")
    public String type;
    public WalletSubject withType(String type) {
        this.type = type;
        return this;
    }
    

    public WalletSubject(@JsonProperty("identifier") String identifier, @JsonProperty("type") String type) {
    this.identifier = identifier;
this.type = type;
  }
}
