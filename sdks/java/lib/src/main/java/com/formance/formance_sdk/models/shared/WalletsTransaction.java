/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.formance.formance_sdk.utils.DateTimeDeserializer;
import com.formance.formance_sdk.utils.DateTimeSerializer;
import java.time.OffsetDateTime;

public class WalletsTransaction {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("ledger")
    public String ledger;
    public WalletsTransaction withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    
    /**
     * Metadata associated with the wallet.
     */
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("metadata")
    public java.util.Map<String, Object> metadata;
    public WalletsTransaction withMetadata(java.util.Map<String, Object> metadata) {
        this.metadata = metadata;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("postCommitVolumes")
    public java.util.Map<String, java.util.Map<String, WalletsVolume>> postCommitVolumes;
    public WalletsTransaction withPostCommitVolumes(java.util.Map<String, java.util.Map<String, WalletsVolume>> postCommitVolumes) {
        this.postCommitVolumes = postCommitVolumes;
        return this;
    }
    
    @JsonProperty("postings")
    public Posting[] postings;
    public WalletsTransaction withPostings(Posting[] postings) {
        this.postings = postings;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("preCommitVolumes")
    public java.util.Map<String, java.util.Map<String, WalletsVolume>> preCommitVolumes;
    public WalletsTransaction withPreCommitVolumes(java.util.Map<String, java.util.Map<String, WalletsVolume>> preCommitVolumes) {
        this.preCommitVolumes = preCommitVolumes;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("reference")
    public String reference;
    public WalletsTransaction withReference(String reference) {
        this.reference = reference;
        return this;
    }
    
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("timestamp")
    public OffsetDateTime timestamp;
    public WalletsTransaction withTimestamp(OffsetDateTime timestamp) {
        this.timestamp = timestamp;
        return this;
    }
    
    @JsonProperty("txid")
    public Long txid;
    public WalletsTransaction withTxid(Long txid) {
        this.txid = txid;
        return this;
    }
    

    public WalletsTransaction(@JsonProperty("postings") Posting[] postings, @JsonProperty("timestamp") OffsetDateTime timestamp, @JsonProperty("txid") Long txid) {
    this.postings = postings;
this.timestamp = timestamp;
this.txid = txid;
  }
}
