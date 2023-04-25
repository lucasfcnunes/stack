/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class AddMetadataToAccountRequest {
    /**
     * metadata
     */
    @SpeakeasyMetadata("request:mediaType=application/json")
    public java.util.Map<String, Object> requestBody;
    public AddMetadataToAccountRequest withRequestBody(java.util.Map<String, Object> requestBody) {
        this.requestBody = requestBody;
        return this;
    }
    
    /**
     * Exact address of the account. It must match the following regular expressions pattern:
     * ```
     * ^\w+(:\w+)*$
     * ```
     * 
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=address")
    public String address;
    public AddMetadataToAccountRequest withAddress(String address) {
        this.address = address;
        return this;
    }
    
    /**
     * Name of the ledger.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=ledger")
    public String ledger;
    public AddMetadataToAccountRequest withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    

    public AddMetadataToAccountRequest(@JsonProperty("RequestBody") java.util.Map<String, Object> requestBody, @JsonProperty("address") String address, @JsonProperty("ledger") String ledger) {
    this.requestBody = requestBody;
this.address = address;
this.ledger = ledger;
  }
}
