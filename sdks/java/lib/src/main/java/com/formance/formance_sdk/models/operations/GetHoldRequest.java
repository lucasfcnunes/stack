/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class GetHoldRequest {
    /**
     * The hold ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=holdID")
    public String holdID;
    public GetHoldRequest withHoldID(String holdID) {
        this.holdID = holdID;
        return this;
    }
    
}
