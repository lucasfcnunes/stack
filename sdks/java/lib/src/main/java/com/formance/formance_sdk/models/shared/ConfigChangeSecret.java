/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class ConfigChangeSecret {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("secret")public String secret;
    public ConfigChangeSecret withSecret(String secret) {
        this.secret = secret;
        return this;
    }
    
}
