/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class Security {
    @SpeakeasyMetadata("security:scheme=true,type=oauth2,name=Authorization")
    public String authorization;
    public Security withAuthorization(String authorization) {
        this.authorization = authorization;
        return this;
    }
    
}
