/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Balance {
    @JsonProperty("name")
    public String name;
    public Balance withName(String name) {
        this.name = name;
        return this;
    }
    
}
