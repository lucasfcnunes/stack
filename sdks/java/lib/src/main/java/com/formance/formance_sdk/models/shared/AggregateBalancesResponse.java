/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * AggregateBalancesResponse - OK
 */
public class AggregateBalancesResponse {
    @JsonProperty("data")
    public java.util.Map<String, Long> data;
    public AggregateBalancesResponse withData(java.util.Map<String, Long> data) {
        this.data = data;
        return this;
    }
    

    public AggregateBalancesResponse(@JsonProperty("data") java.util.Map<String, Long> data) {
    this.data = data;
  }
}
