/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import java.net.http.HttpResponse;

public class CreateWalletResponse {
    
    public String contentType;
    public CreateWalletResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * Wallet created
     */
    
    public com.formance.formance_sdk.models.shared.CreateWalletResponse createWalletResponse;
    public CreateWalletResponse withCreateWalletResponse(com.formance.formance_sdk.models.shared.CreateWalletResponse createWalletResponse) {
        this.createWalletResponse = createWalletResponse;
        return this;
    }
    
    
    public Integer statusCode;
    public CreateWalletResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;
    public CreateWalletResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    /**
     * Error
     */
    
    public com.formance.formance_sdk.models.shared.WalletsErrorResponse walletsErrorResponse;
    public CreateWalletResponse withWalletsErrorResponse(com.formance.formance_sdk.models.shared.WalletsErrorResponse walletsErrorResponse) {
        this.walletsErrorResponse = walletsErrorResponse;
        return this;
    }
    
}
