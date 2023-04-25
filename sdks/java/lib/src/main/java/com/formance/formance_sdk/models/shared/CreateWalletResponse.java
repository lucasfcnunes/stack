/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * CreateWalletResponse - Wallet created
 */
public class CreateWalletResponse {
    @JsonProperty("data")
    public Wallet data;
    public CreateWalletResponse withData(Wallet data) {
        this.data = data;
        return this;
    }
    

    public CreateWalletResponse(@JsonProperty("data") Wallet data) {
    this.data = data;
  }
}
