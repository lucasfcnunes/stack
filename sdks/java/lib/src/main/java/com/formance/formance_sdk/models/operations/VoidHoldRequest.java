/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class VoidHoldRequest {
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=hold_id")public String holdId;
    public VoidHoldRequest withHoldId(String holdId) {
        this.holdId = holdId;
        return this;
    }
    
}
