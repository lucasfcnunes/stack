/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class CreateWalletRequest {
    /**
     * Custom metadata to attach to this wallet.
     */
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("metadata")
    public java.util.Map<String, Object> metadata;
    public CreateWalletRequest withMetadata(java.util.Map<String, Object> metadata) {
        this.metadata = metadata;
        return this;
    }
    
    @JsonProperty("name")
    public String name;
    public CreateWalletRequest withName(String name) {
        this.name = name;
        return this;
    }
    

    public CreateWalletRequest(@JsonProperty("name") String name) {
    this.name = name;
  }
}
