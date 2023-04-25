/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class ResetConnectorRequest {
    /**
     * The name of the connector.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=connector")
    public com.formance.formance_sdk.models.shared.ConnectorEnum connector;
    public ResetConnectorRequest withConnector(com.formance.formance_sdk.models.shared.ConnectorEnum connector) {
        this.connector = connector;
        return this;
    }
    

    public ResetConnectorRequest(@JsonProperty("connector") com.formance.formance_sdk.models.shared.ConnectorEnum connector) {
    this.connector = connector;
  }
}
