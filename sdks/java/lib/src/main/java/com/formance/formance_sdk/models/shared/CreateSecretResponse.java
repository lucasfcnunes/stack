/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * CreateSecretResponse - Created secret
 */
public class CreateSecretResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("data")
    public Secret data;
    public CreateSecretResponse withData(Secret data) {
        this.data = data;
        return this;
    }
    

    public CreateSecretResponse(){}
}
