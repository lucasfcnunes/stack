/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class DeleteTransientScopeRequest {
    /**
     * Scope ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=scopeId")
    public String scopeId;
    public DeleteTransientScopeRequest withScopeId(String scopeId) {
        this.scopeId = scopeId;
        return this;
    }
    
    /**
     * Transient scope ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=transientScopeId")
    public String transientScopeId;
    public DeleteTransientScopeRequest withTransientScopeId(String transientScopeId) {
        this.transientScopeId = transientScopeId;
        return this;
    }
    
}
