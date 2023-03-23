/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class GetInstanceHistoryRequest {
    /**
     * The instance id
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=instanceID")public String instanceID;
    public GetInstanceHistoryRequest withInstanceID(String instanceID) {
        this.instanceID = instanceID;
        return this;
    }
    
}
