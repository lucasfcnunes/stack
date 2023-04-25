/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * AccountsCursorResponse - OK
 */
public class AccountsCursorResponse {
    @JsonProperty("cursor")
    public AccountsCursorResponseCursor cursor;
    public AccountsCursorResponse withCursor(AccountsCursorResponseCursor cursor) {
        this.cursor = cursor;
        return this;
    }
    

    public AccountsCursorResponse(@JsonProperty("cursor") AccountsCursorResponseCursor cursor) {
    this.cursor = cursor;
  }
}
