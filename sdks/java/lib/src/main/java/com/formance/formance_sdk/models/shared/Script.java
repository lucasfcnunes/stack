/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class Script {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("metadata")
    public java.util.Map<String, Object> metadata;
    public Script withMetadata(java.util.Map<String, Object> metadata) {
        this.metadata = metadata;
        return this;
    }
    
    @JsonProperty("plain")
    public String plain;
    public Script withPlain(String plain) {
        this.plain = plain;
        return this;
    }
    
    /**
     * Reference to attach to the generated transaction
     */
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("reference")
    public String reference;
    public Script withReference(String reference) {
        this.reference = reference;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("vars")
    public java.util.Map<String, Object> vars;
    public Script withVars(java.util.Map<String, Object> vars) {
        this.vars = vars;
        return this;
    }
    

    public Script(@JsonProperty("plain") String plain) {
    this.plain = plain;
  }
}
