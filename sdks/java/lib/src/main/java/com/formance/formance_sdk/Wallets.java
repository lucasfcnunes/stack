/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.formance.formance_sdk.utils.HTTPClient;
import com.formance.formance_sdk.utils.HTTPRequest;
import com.formance.formance_sdk.utils.JSON;
import com.formance.formance_sdk.utils.SerializedBody;
import java.net.http.HttpResponse;
import java.nio.charset.StandardCharsets;
import org.apache.http.NameValuePair;

public class Wallets {
	
	private HTTPClient _defaultClient;
	private HTTPClient _securityClient;
	private String _serverUrl;
	private String _language;
	private String _sdkVersion;
	private String _genVersion;

	public Wallets(HTTPClient defaultClient, HTTPClient securityClient, String serverUrl, String language, String sdkVersion, String genVersion) {
		this._defaultClient = defaultClient;
		this._securityClient = securityClient;
		this._serverUrl = serverUrl;
		this._language = language;
		this._sdkVersion = sdkVersion;
		this._genVersion = genVersion;
	}

    /**
     * Confirm a hold
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ConfirmHoldResponse confirmHold(com.formance.formance_sdk.models.operations.ConfirmHoldRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.ConfirmHoldRequest.class, baseUrl, "/api/wallets/holds/{hold_id}/confirm", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "confirmHoldRequest", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ConfirmHoldResponse res = new com.formance.formance_sdk.models.operations.ConfirmHoldResponse(contentType, httpRes.statusCode()) {{
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Create a balance
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.CreateBalanceResponse createBalance(com.formance.formance_sdk.models.operations.CreateBalanceRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.CreateBalanceRequest.class, baseUrl, "/api/wallets/wallets/{id}/balances", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "createBalanceRequest", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.CreateBalanceResponse res = new com.formance.formance_sdk.models.operations.CreateBalanceResponse(contentType, httpRes.statusCode()) {{
            createBalanceResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 201) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.CreateBalanceResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.CreateBalanceResponse.class);
                res.createBalanceResponse = out;
            }
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Create a new wallet
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.CreateWalletResponse createWallet(com.formance.formance_sdk.models.shared.CreateWalletRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/wallets/wallets");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "request", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.CreateWalletResponse res = new com.formance.formance_sdk.models.operations.CreateWalletResponse(contentType, httpRes.statusCode()) {{
            createWalletResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 201) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.CreateWalletResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.CreateWalletResponse.class);
                res.createWalletResponse = out;
            }
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Credit a wallet
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.CreditWalletResponse creditWallet(com.formance.formance_sdk.models.operations.CreditWalletRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.CreditWalletRequest.class, baseUrl, "/api/wallets/wallets/{id}/credit", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "creditWalletRequest", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.CreditWalletResponse res = new com.formance.formance_sdk.models.operations.CreditWalletResponse(contentType, httpRes.statusCode()) {{
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Debit a wallet
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.DebitWalletResponse debitWallet(com.formance.formance_sdk.models.operations.DebitWalletRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.DebitWalletRequest.class, baseUrl, "/api/wallets/wallets/{id}/debit", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "debitWalletRequest", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.DebitWalletResponse res = new com.formance.formance_sdk.models.operations.DebitWalletResponse(contentType, httpRes.statusCode()) {{
            debitWalletResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 201) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.DebitWalletResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.DebitWalletResponse.class);
                res.debitWalletResponse = out;
            }
        }
        else if (httpRes.statusCode() == 204) {
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Get detailed balance
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.GetBalanceResponse getBalance(com.formance.formance_sdk.models.operations.GetBalanceRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.GetBalanceRequest.class, baseUrl, "/api/wallets/wallets/{id}/balances/{balanceName}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.GetBalanceResponse res = new com.formance.formance_sdk.models.operations.GetBalanceResponse(contentType, httpRes.statusCode()) {{
            getBalanceResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.GetBalanceResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.GetBalanceResponse.class);
                res.getBalanceResponse = out;
            }
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Get a hold
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.GetHoldResponse getHold(com.formance.formance_sdk.models.operations.GetHoldRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.GetHoldRequest.class, baseUrl, "/api/wallets/holds/{holdID}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.GetHoldResponse res = new com.formance.formance_sdk.models.operations.GetHoldResponse(contentType, httpRes.statusCode()) {{
            getHoldResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.GetHoldResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.GetHoldResponse.class);
                res.getHoldResponse = out;
            }
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Get all holds for a wallet
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.GetHoldsResponse getHolds(com.formance.formance_sdk.models.operations.GetHoldsRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/wallets/holds");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        java.util.List<NameValuePair> queryParams = com.formance.formance_sdk.utils.Utils.getQueryParams(com.formance.formance_sdk.models.operations.GetHoldsRequest.class, request, null);
        if (queryParams != null) {
            for (NameValuePair queryParam : queryParams) {
                req.addQueryParam(queryParam);
            }
        }
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.GetHoldsResponse res = new com.formance.formance_sdk.models.operations.GetHoldsResponse(contentType, httpRes.statusCode()) {{
            getHoldsResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.GetHoldsResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.GetHoldsResponse.class);
                res.getHoldsResponse = out;
            }
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    public com.formance.formance_sdk.models.operations.GetTransactionsResponse getTransactions(com.formance.formance_sdk.models.operations.GetTransactionsRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/wallets/transactions");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        java.util.List<NameValuePair> queryParams = com.formance.formance_sdk.utils.Utils.getQueryParams(com.formance.formance_sdk.models.operations.GetTransactionsRequest.class, request, null);
        if (queryParams != null) {
            for (NameValuePair queryParam : queryParams) {
                req.addQueryParam(queryParam);
            }
        }
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.GetTransactionsResponse res = new com.formance.formance_sdk.models.operations.GetTransactionsResponse(contentType, httpRes.statusCode()) {{
            getTransactionsResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.GetTransactionsResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.GetTransactionsResponse.class);
                res.getTransactionsResponse = out;
            }
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Get a wallet
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.GetWalletResponse getWallet(com.formance.formance_sdk.models.operations.GetWalletRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.GetWalletRequest.class, baseUrl, "/api/wallets/wallets/{id}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.GetWalletResponse res = new com.formance.formance_sdk.models.operations.GetWalletResponse(contentType, httpRes.statusCode()) {{
            getWalletResponse = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.GetWalletResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.GetWalletResponse.class);
                res.getWalletResponse = out;
            }
        }
        else if (httpRes.statusCode() == 404) {
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * List balances of a wallet
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ListBalancesResponse listBalances(com.formance.formance_sdk.models.operations.ListBalancesRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.ListBalancesRequest.class, baseUrl, "/api/wallets/wallets/{id}/balances", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ListBalancesResponse res = new com.formance.formance_sdk.models.operations.ListBalancesResponse(contentType, httpRes.statusCode()) {{
            listBalancesResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ListBalancesResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ListBalancesResponse.class);
                res.listBalancesResponse = out;
            }
        }

        return res;
    }

    /**
     * List all wallets
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.ListWalletsResponse listWallets(com.formance.formance_sdk.models.operations.ListWalletsRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/wallets/wallets");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        java.util.List<NameValuePair> queryParams = com.formance.formance_sdk.utils.Utils.getQueryParams(com.formance.formance_sdk.models.operations.ListWalletsRequest.class, request, null);
        if (queryParams != null) {
            for (NameValuePair queryParam : queryParams) {
                req.addQueryParam(queryParam);
            }
        }
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.ListWalletsResponse res = new com.formance.formance_sdk.models.operations.ListWalletsResponse(contentType, httpRes.statusCode()) {{
            listWalletsResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ListWalletsResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ListWalletsResponse.class);
                res.listWalletsResponse = out;
            }
        }

        return res;
    }

    /**
     * Update a wallet
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.UpdateWalletResponse updateWallet(com.formance.formance_sdk.models.operations.UpdateWalletRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.UpdateWalletRequest.class, baseUrl, "/api/wallets/wallets/{id}", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("PATCH");
        req.setURL(url);
        SerializedBody serializedRequestBody = com.formance.formance_sdk.utils.Utils.serializeRequestBody(request, "requestBody", "json");
        req.setBody(serializedRequestBody);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.UpdateWalletResponse res = new com.formance.formance_sdk.models.operations.UpdateWalletResponse(contentType, httpRes.statusCode()) {{
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Cancel a hold
     * @param request the request object containing all of the parameters for the API call
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.VoidHoldResponse voidHold(com.formance.formance_sdk.models.operations.VoidHoldRequest request) throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(com.formance.formance_sdk.models.operations.VoidHoldRequest.class, baseUrl, "/api/wallets/holds/{hold_id}/void", request, null);
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("POST");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.VoidHoldResponse res = new com.formance.formance_sdk.models.operations.VoidHoldResponse(contentType, httpRes.statusCode()) {{
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 204) {
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }

    /**
     * Get server info
     * @return the response from the API call
     * @throws Exception if the API call fails
     */
    public com.formance.formance_sdk.models.operations.WalletsgetServerInfoResponse walletsgetServerInfo() throws Exception {
        String baseUrl = this._serverUrl;
        String url = com.formance.formance_sdk.utils.Utils.generateURL(baseUrl, "/api/wallets/_info");
        
        HTTPRequest req = new HTTPRequest();
        req.setMethod("GET");
        req.setURL(url);
        
        
        HTTPClient client = this._securityClient;
        
        HttpResponse<byte[]> httpRes = client.send(req);

        String contentType = httpRes.headers().firstValue("Content-Type").orElse("application/octet-stream");

        com.formance.formance_sdk.models.operations.WalletsgetServerInfoResponse res = new com.formance.formance_sdk.models.operations.WalletsgetServerInfoResponse(contentType, httpRes.statusCode()) {{
            serverInfo = null;
            walletsErrorResponse = null;
        }};
        res.rawResponse = httpRes;
        
        if (httpRes.statusCode() == 200) {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.ServerInfo out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.ServerInfo.class);
                res.serverInfo = out;
            }
        }
        else {
            if (com.formance.formance_sdk.utils.Utils.matchContentType(contentType, "application/json")) {
                ObjectMapper mapper = JSON.getMapper();
                com.formance.formance_sdk.models.shared.WalletsErrorResponse out = mapper.readValue(new String(httpRes.body(), StandardCharsets.UTF_8), com.formance.formance_sdk.models.shared.WalletsErrorResponse.class);
                res.walletsErrorResponse = out;
            }
        }

        return res;
    }
}