/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class Scope {
    @JsonProperty("id")
    public String id;
    public Scope withId(String id) {
        this.id = id;
        return this;
    }
    
    @JsonProperty("label")
    public String label;
    public Scope withLabel(String label) {
        this.label = label;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("metadata")
    public java.util.Map<String, Object> metadata;
    public Scope withMetadata(java.util.Map<String, Object> metadata) {
        this.metadata = metadata;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("transient")
    public String[] transient_;
    public Scope withTransient(String[] transient_) {
        this.transient_ = transient_;
        return this;
    }
    
}
