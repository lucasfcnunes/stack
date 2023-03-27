/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import * as utils from "../internal/utils";
import * as operations from "./models/operations";
import * as shared from "./models/shared";
import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";

export class Orchestration {
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
   * Cancel a running workflow
   *
   * @remarks
   * Cancel a running workflow
   */
  cancelEvent(
    req: operations.CancelEventRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.CancelEventResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.CancelEventRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/orchestration/instances/{instanceID}/abort",
      req
    );

    const client: AxiosInstance = this._securityClient || this._defaultClient;

    const r = client.request({
      url: url,
      method: "put",
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.CancelEventResponse =
        new operations.CancelEventResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 204:
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Create workflow
   *
   * @remarks
   * Create a workflow
   */
  createWorkflow(
    req: shared.CreateWorkflowRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.CreateWorkflowResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new shared.CreateWorkflowRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string =
      baseURL.replace(/\/$/, "") + "/api/orchestration/workflows";

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
      const res: operations.CreateWorkflowResponse =
        new operations.CreateWorkflowResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 201:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.createWorkflowResponse = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.CreateWorkflowResponse
            );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Get a workflow instance by id
   *
   * @remarks
   * Get a workflow instance by id
   */
  getInstance(
    req: operations.GetInstanceRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.GetInstanceResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.GetInstanceRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/orchestration/instances/{instanceID}",
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
      const res: operations.GetInstanceResponse =
        new operations.GetInstanceResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.getWorkflowInstanceResponse = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.GetWorkflowInstanceResponse
            );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Get a workflow instance history by id
   *
   * @remarks
   * Get a workflow instance history by id
   */
  getInstanceHistory(
    req: operations.GetInstanceHistoryRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.GetInstanceHistoryResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.GetInstanceHistoryRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/orchestration/instances/{instanceID}/history",
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
      const res: operations.GetInstanceHistoryResponse =
        new operations.GetInstanceHistoryResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.getWorkflowInstanceHistoryResponse =
              utils.deserializeJSONResponse(
                httpRes?.data,
                shared.GetWorkflowInstanceHistoryResponse
              );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Get a workflow instance stage history
   *
   * @remarks
   * Get a workflow instance stage history
   */
  getInstanceStageHistory(
    req: operations.GetInstanceStageHistoryRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.GetInstanceStageHistoryResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.GetInstanceStageHistoryRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/orchestration/instances/{instanceID}/stages/{number}/history",
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
      const res: operations.GetInstanceStageHistoryResponse =
        new operations.GetInstanceStageHistoryResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.getWorkflowInstanceHistoryStageResponse =
              utils.deserializeJSONResponse(
                httpRes?.data,
                shared.GetWorkflowInstanceHistoryStageResponse
              );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Get a flow by id
   *
   * @remarks
   * Get a flow by id
   */
  getWorkflow(
    req: operations.GetWorkflowRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.GetWorkflowResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.GetWorkflowRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/orchestration/workflows/{flowId}",
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
      const res: operations.GetWorkflowResponse =
        new operations.GetWorkflowResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.getWorkflowResponse = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.GetWorkflowResponse
            );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * List instances of a workflow
   *
   * @remarks
   * List instances of a workflow
   */
  listInstances(
    req: operations.ListInstancesRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.ListInstancesResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.ListInstancesRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string =
      baseURL.replace(/\/$/, "") + "/api/orchestration/instances";

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
      const res: operations.ListInstancesResponse =
        new operations.ListInstancesResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.listRunsResponse = httpRes?.data;
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * List registered workflows
   *
   * @remarks
   * List registered workflows
   */
  listWorkflows(
    config?: AxiosRequestConfig
  ): Promise<operations.ListWorkflowsResponse> {
    const baseURL: string = this._serverURL;
    const url: string =
      baseURL.replace(/\/$/, "") + "/api/orchestration/workflows";

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
      const res: operations.ListWorkflowsResponse =
        new operations.ListWorkflowsResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.listWorkflowsResponse = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ListWorkflowsResponse
            );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
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
  orchestrationgetServerInfo(
    config?: AxiosRequestConfig
  ): Promise<operations.OrchestrationgetServerInfoResponse> {
    const baseURL: string = this._serverURL;
    const url: string = baseURL.replace(/\/$/, "") + "/api/orchestration/_info";

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
      const res: operations.OrchestrationgetServerInfoResponse =
        new operations.OrchestrationgetServerInfoResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 200:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.serverInfo = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ServerInfo
            );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Run workflow
   *
   * @remarks
   * Run workflow
   */
  runWorkflow(
    req: operations.RunWorkflowRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.RunWorkflowResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.RunWorkflowRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/orchestration/workflows/{workflowID}/instances",
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
    const queryParams: string = utils.serializeQueryParams(req);

    const r = client.request({
      url: url + queryParams,
      method: "post",
      headers: headers,
      data: reqBody,
      ...config,
    });

    return r.then((httpRes: AxiosResponse) => {
      const contentType: string = httpRes?.headers?.["content-type"] ?? "";

      if (httpRes?.status == null)
        throw new Error(`status code not found in response: ${httpRes}`);
      const res: operations.RunWorkflowResponse =
        new operations.RunWorkflowResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 201:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.runWorkflowResponse = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.RunWorkflowResponse
            );
          }
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }

  /**
   * Send an event to a running workflow
   *
   * @remarks
   * Send an event to a running workflow
   */
  sendEvent(
    req: operations.SendEventRequest,
    config?: AxiosRequestConfig
  ): Promise<operations.SendEventResponse> {
    if (!(req instanceof utils.SpeakeasyBase)) {
      req = new operations.SendEventRequest(req);
    }

    const baseURL: string = this._serverURL;
    const url: string = utils.generateURL(
      baseURL,
      "/api/orchestration/instances/{instanceID}/events",
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
      const res: operations.SendEventResponse =
        new operations.SendEventResponse({
          statusCode: httpRes.status,
          contentType: contentType,
          rawResponse: httpRes,
        });
      switch (true) {
        case httpRes?.status == 204:
          break;
        default:
          if (utils.matchContentType(contentType, `application/json`)) {
            res.error = utils.deserializeJSONResponse(
              httpRes?.data,
              shared.ErrorT
            );
          }
          break;
      }

      return res;
    });
  }
}
