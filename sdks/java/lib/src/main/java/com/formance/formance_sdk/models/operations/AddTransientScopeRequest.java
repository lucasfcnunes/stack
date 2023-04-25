/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class AddTransientScopeRequest {
    /**
     * Scope ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=scopeId")
    public String scopeId;
    public AddTransientScopeRequest withScopeId(String scopeId) {
        this.scopeId = scopeId;
        return this;
    }
    
    /**
     * Transient scope ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=transientScopeId")
    public String transientScopeId;
    public AddTransientScopeRequest withTransientScopeId(String transientScopeId) {
        this.transientScopeId = transientScopeId;
        return this;
    }
    

    public AddTransientScopeRequest(@JsonProperty("scopeId") String scopeId, @JsonProperty("transientScopeId") String transientScopeId) {
    this.scopeId = scopeId;
this.transientScopeId = transientScopeId;
  }
}
