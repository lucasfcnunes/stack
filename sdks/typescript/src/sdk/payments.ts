/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as utils from "../internal/utils";
import * as operations from "./models/operations";
import * as shared from "./models/shared";
import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";

export class Payments {
  _defaultClient: AxiosInstance;
  _securityClient: AxiosInstance;
  _serverURL: string;
  _language: string;
  _sdkVersion: string;
  _genVersion: string;

  constructor(
    defaultClient: AxiosInstance,
    securityClient: AxiosInstance,
    serverURL: string,
    language: string,
    sdkVersion: string,
    genVersion: string
  ) {
    this._defaultClient = defaultClient;
    this._securityClient = securityClient;
    this._serverURL = serverURL;
    this._language = language;
    this._sdkVersion = sdkVersion;
    this._genVersion = genVersion;
  }

  /**
   * Transfer funds between Stripe accounts
   *
   * @remarks
   * Execute a transfer between two Stripe accounts.
   */
  connectorsStripeTransfer(
    req: shared.StripeTransferRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ConnectorsStripeTransferResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new shared.StripeTransferRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string =
      baseURL.replace(/\/$/, "") + "/api/payments/connectors/stripe/transfers";

    let [reqBodyHeaders, reqBody]: [object, any] = [{}, {}];

    try {
      [reqBodyHeaders, reqBody] = utils.serializeRequestBody(
        req,
        "request",
        "json"
      );
    } catch (e: unknown) {
      if (e instanceof Error) {
        throw new Error(`Error serializing request body, cause: ${e.message}`);
      }
    }

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const headers = { ...reqBodyHeaders, ...config?.headers };
    if (reqBody == null || Object.keys(reqBody).length === 0)
      throw new Error("request body is required");

    const r = client.request({
      url: url,
      method: "post",
      headers: headers,
      data: reqBody,
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ConnectorsStripeTransferResponse =
        new operations.ConnectorsStripeTransferResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.stripeTransferResponse = utils.objectToClass(httpRes?.data);
          }
          break;
      }

      return res;
    });
  }

  /**
   * Transfer funds between Connector accounts
   *
   * @remarks
   * Execute a transfer between two accounts.
   */
  connectorsTransfer(
    req: operations.ConnectorsTransferRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ConnectorsTransferResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.ConnectorsTransferRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}/transfers",
      req
    );

    let [reqBodyHeaders, reqBody]: [object, any] = [{}, {}];

    try {
      [reqBodyHeaders, reqBody] = utils.serializeRequestBody(
        req,
        "transferRequest",
        "json"
      );
    } catch (e: unknown) {
      if (e instanceof Error) {
        throw new Error(`Error serializing request body, cause: ${e.message}`);
      }
    }

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const headers = { ...reqBodyHeaders, ...config?.headers };
    if (reqBody == null || Object.keys(reqBody).length === 0)
      throw new Error("request body is required");

    const r = client.request({
      url: url,
      method: "post",
      headers: headers,
      data: reqBody,
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ConnectorsTransferResponse =
        new operations.ConnectorsTransferResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.transferResponse = utils.objectToClass(
              httpRes?.data,
              shared.TransferResponse
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Read a specific task of the connector
   *
   * @remarks
   * Get a specific task associated to the connector.
   */
  getConnectorTask(
    req: operations.GetConnectorTaskRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.GetConnectorTaskResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.GetConnectorTaskRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}/tasks/{taskId}",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.GetConnectorTaskResponse =
        new operations.GetConnectorTaskResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.taskResponse = utils.objectToClass(
              httpRes?.data,
              shared.TaskResponse
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Get a payment
   */
  getPayment(
    req: operations.GetPaymentRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.GetPaymentResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.GetPaymentRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/payments/{paymentId}",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.GetPaymentResponse =
        new operations.GetPaymentResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.paymentResponse = utils.objectToClass(
              httpRes?.data,
              shared.PaymentResponse
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Install a connector
   *
   * @remarks
   * Install a connector by its name and config.
   */
  installConnector(
    req: operations.InstallConnectorRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.InstallConnectorResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.InstallConnectorRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}",
      req
    );

    let [reqBodyHeaders, reqBody]: [object, any] = [{}, {}];

    try {
      [reqBodyHeaders, reqBody] = utils.serializeRequestBody(
        req,
        "requestBody",
        "json"
      );
    } catch (e: unknown) {
      if (e instanceof Error) {
        throw new Error(`Error serializing request body, cause: ${e.message}`);
      }
    }

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const headers = { ...reqBodyHeaders, ...config?.headers };
    if (reqBody == null || Object.keys(reqBody).length === 0)
      throw new Error("request body is required");

    const r = client.request({
      url: url,
      method: "post",
      headers: headers,
      data: reqBody,
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.InstallConnectorResponse =
        new operations.InstallConnectorResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 204:
          break;
      }

      return res;
    });
  }

  /**
   * List all installed connectors
   *
   * @remarks
   * List all installed connectors.
   */
  listAllConnectors(
    config?: AxiosRequestConfig
  ): Promise<operations.ListAllConnectorsResponse> {
    const baseURL: string = this._serverURL;
    const url: string = baseURL.replace(/\/$/, "") + "/api/payments/connectors";

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ListAllConnectorsResponse =
        new operations.ListAllConnectorsResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.connectorsResponse = utils.objectToClass(
              httpRes?.data,
              shared.ConnectorsResponse
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * List the configs of each available connector
   *
   * @remarks
   * List the configs of each available connector.
   */
  listConfigsAvailableConnectors(
    config?: AxiosRequestConfig
  ): Promise<operations.ListConfigsAvailableConnectorsResponse> {
    const baseURL: string = this._serverURL;
    const url: string =
      baseURL.replace(/\/$/, "") + "/api/payments/connectors/configs";

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ListConfigsAvailableConnectorsResponse =
        new operations.ListConfigsAvailableConnectorsResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.connectorsConfigsResponse = utils.objectToClass(
              httpRes?.data,
              shared.ConnectorsConfigsResponse
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * List tasks from a connector
   *
   * @remarks
   * List all tasks associated with this connector.
   */
  listConnectorTasks(
    req: operations.ListConnectorTasksRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ListConnectorTasksResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.ListConnectorTasksRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}/tasks",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const queryParams: string = utils.serializeQueryParams(req);

    const r = client.request({
      url: url + queryParams,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ListConnectorTasksResponse =
        new operations.ListConnectorTasksResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.tasksCursor = utils.objectToClass(
              httpRes?.data,
              shared.TasksCursor
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * List transfers and their statuses
   *
   * @remarks
   * List transfers
   */
  listConnectorsTransfers(
    req: operations.ListConnectorsTransfersRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ListConnectorsTransfersResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.ListConnectorsTransfersRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}/transfers",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ListConnectorsTransfersResponse =
        new operations.ListConnectorsTransfersResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.transfersResponse = utils.objectToClass(
              httpRes?.data,
              shared.TransfersResponse
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * List payments
   */
  listPayments(
    req: operations.ListPaymentsRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ListPaymentsResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.ListPaymentsRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = baseURL.replace(/\/$/, "") + "/api/payments/payments";

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const queryParams: string = utils.serializeQueryParams(req);

    const r = client.request({
      url: url + queryParams,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ListPaymentsResponse =
        new operations.ListPaymentsResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.paymentsCursor = utils.objectToClass(
              httpRes?.data,
              shared.PaymentsCursor
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Get server info
   */
  paymentsgetServerInfo(
    config?: AxiosRequestConfig
  ): Promise<operations.PaymentsgetServerInfoResponse> {
    const baseURL: string = this._serverURL;
    const url: string = baseURL.replace(/\/$/, "") + "/api/payments/_info";

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.PaymentsgetServerInfoResponse =
        new operations.PaymentsgetServerInfoResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.serverInfo = utils.objectToClass(
              httpRes?.data,
              shared.ServerInfo
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * List accounts
   */
  paymentslistAccounts(
    req: operations.PaymentslistAccountsRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.PaymentslistAccountsResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.PaymentslistAccountsRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = baseURL.replace(/\/$/, "") + "/api/payments/accounts";

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const queryParams: string = utils.serializeQueryParams(req);

    const r = client.request({
      url: url + queryParams,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.PaymentslistAccountsResponse =
        new operations.PaymentslistAccountsResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.accountsCursor = utils.objectToClass(
              httpRes?.data,
              shared.AccountsCursor
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Read the config of a connector
   *
   * @remarks
   * Read connector config
   */
  readConnectorConfig(
    req: operations.ReadConnectorConfigRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ReadConnectorConfigResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.ReadConnectorConfigRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}/config",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "get",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ReadConnectorConfigResponse =
        new operations.ReadConnectorConfigResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.connectorConfigResponse = utils.objectToClass(
              httpRes?.data,
              shared.ConnectorConfigResponse
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Reset a connector
   *
   * @remarks
   * Reset a connector by its name.
   * It will remove the connector and ALL PAYMENTS generated with it.
   *
   */
  resetConnector(
    req: operations.ResetConnectorRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ResetConnectorResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.ResetConnectorRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}/reset",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "post",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.ResetConnectorResponse =
        new operations.ResetConnectorResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 204:
          break;
      }

      return res;
    });
  }

  /**
   * Uninstall a connector
   *
   * @remarks
   * Uninstall a connector by its name.
   */
  uninstallConnector(
    req: operations.UninstallConnectorRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.UninstallConnectorResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.UninstallConnectorRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/connectors/{connector}",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "delete",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.UninstallConnectorResponse =
        new operations.UninstallConnectorResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 204:
          break;
      }

      return res;
    });
  }

  /**
   * Update metadata
   */
  updateMetadata(
    req: operations.UpdateMetadataRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.UpdateMetadataResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.UpdateMetadataRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/payments/payments/{paymentId}/metadata",
      req
    );

    let [reqBodyHeaders, reqBody]: [object, any] = [{}, {}];

    try {
      [reqBodyHeaders, reqBody] = utils.serializeRequestBody(
        req,
        "paymentMetadata",
        "json"
      );
    } catch (e: unknown) {
      if (e instanceof Error) {
        throw new Error(`Error serializing request body, cause: ${e.message}`);
      }
    }

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const headers = { ...reqBodyHeaders, ...config?.headers };
    if (reqBody == null || Object.keys(reqBody).length === 0)
      throw new Error("request body is required");

    const r = client.request({
      url: url,
      method: "patch",
      headers: headers,
      data: reqBody,
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.UpdateMetadataResponse =
        new operations.UpdateMetadataResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 204:
          break;
      }

      return res;
    });
  }
}
