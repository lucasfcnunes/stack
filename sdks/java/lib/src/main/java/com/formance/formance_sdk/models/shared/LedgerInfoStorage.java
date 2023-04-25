/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class LedgerInfoStorage {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("migrations")
    public MigrationInfo[] migrations;
    public LedgerInfoStorage withMigrations(MigrationInfo[] migrations) {
        this.migrations = migrations;
        return this;
    }
    

    public LedgerInfoStorage(){}
}
