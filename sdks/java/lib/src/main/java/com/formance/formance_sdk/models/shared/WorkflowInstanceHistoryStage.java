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

public class WorkflowInstanceHistoryStage {
    @JsonProperty("attempt")
    public Long attempt;
    public WorkflowInstanceHistoryStage withAttempt(Long attempt) {
        this.attempt = attempt;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("error")
    public String error;
    public WorkflowInstanceHistoryStage withError(String error) {
        this.error = error;
        return this;
    }
    
    @JsonProperty("input")
    public WorkflowInstanceHistoryStageInput input;
    public WorkflowInstanceHistoryStage withInput(WorkflowInstanceHistoryStageInput input) {
        this.input = input;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("lastFailure")
    public String lastFailure;
    public WorkflowInstanceHistoryStage withLastFailure(String lastFailure) {
        this.lastFailure = lastFailure;
        return this;
    }
    
    @JsonProperty("name")
    public String name;
    public WorkflowInstanceHistoryStage withName(String name) {
        this.name = name;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("nextExecution")
    public OffsetDateTime nextExecution;
    public WorkflowInstanceHistoryStage withNextExecution(OffsetDateTime nextExecution) {
        this.nextExecution = nextExecution;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("output")
    public WorkflowInstanceHistoryStageOutput output;
    public WorkflowInstanceHistoryStage withOutput(WorkflowInstanceHistoryStageOutput output) {
        this.output = output;
        return this;
    }
    
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("startedAt")
    public OffsetDateTime startedAt;
    public WorkflowInstanceHistoryStage withStartedAt(OffsetDateTime startedAt) {
        this.startedAt = startedAt;
        return this;
    }
    
    @JsonProperty("terminated")
    public Boolean terminated;
    public WorkflowInstanceHistoryStage withTerminated(Boolean terminated) {
        this.terminated = terminated;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("terminatedAt")
    public OffsetDateTime terminatedAt;
    public WorkflowInstanceHistoryStage withTerminatedAt(OffsetDateTime terminatedAt) {
        this.terminatedAt = terminatedAt;
        return this;
    }
    
}
