"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

import requests as requests_http
from . import utils
from sdk.models import operations, shared
from typing import Any, Optional

class Orchestration:
    _client: requests_http.Session
    _security_client: requests_http.Session
    _server_url: str
    _language: str
    _sdk_version: str
    _gen_version: str

    def __init__(self, client: requests_http.Session, security_client: requests_http.Session, server_url: str, language: str, sdk_version: str, gen_version: str) -> None:
        self._client = client
        self._security_client = security_client
        self._server_url = server_url
        self._language = language
        self._sdk_version = sdk_version
        self._gen_version = gen_version
        
    def cancel_event(self, request: operations.CancelEventRequest) -> operations.CancelEventResponse:
        r"""Cancel a running workflow
        Cancel a running workflow
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.CancelEventRequest, base_url, '/api/orchestration/instances/{instanceID}/abort', request)
        
        
        client = self._security_client
        
        http_res = client.request('PUT', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.CancelEventResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 204:
            pass
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def create_workflow(self, request: shared.CreateWorkflowRequest) -> operations.CreateWorkflowResponse:
        r"""Create workflow
        Create a workflow
        """
        base_url = self._server_url
        
        url = base_url.removesuffix('/') + '/api/orchestration/workflows'
        
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "request", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        
        client = self._security_client
        
        http_res = client.request('POST', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.CreateWorkflowResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 201:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.CreateWorkflowResponse])
                res.create_workflow_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def get_instance(self, request: operations.GetInstanceRequest) -> operations.GetInstanceResponse:
        r"""Get a workflow instance by id
        Get a workflow instance by id
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.GetInstanceRequest, base_url, '/api/orchestration/instances/{instanceID}', request)
        
        
        client = self._security_client
        
        http_res = client.request('GET', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetInstanceResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.GetWorkflowInstanceResponse])
                res.get_workflow_instance_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def get_instance_history(self, request: operations.GetInstanceHistoryRequest) -> operations.GetInstanceHistoryResponse:
        r"""Get a workflow instance history by id
        Get a workflow instance history by id
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.GetInstanceHistoryRequest, base_url, '/api/orchestration/instances/{instanceID}/history', request)
        
        
        client = self._security_client
        
        http_res = client.request('GET', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetInstanceHistoryResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.GetWorkflowInstanceHistoryResponse])
                res.get_workflow_instance_history_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def get_instance_stage_history(self, request: operations.GetInstanceStageHistoryRequest) -> operations.GetInstanceStageHistoryResponse:
        r"""Get a workflow instance stage history
        Get a workflow instance stage history
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.GetInstanceStageHistoryRequest, base_url, '/api/orchestration/instances/{instanceID}/stages/{number}/history', request)
        
        
        client = self._security_client
        
        http_res = client.request('GET', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetInstanceStageHistoryResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.GetWorkflowInstanceHistoryStageResponse])
                res.get_workflow_instance_history_stage_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def get_workflow(self, request: operations.GetWorkflowRequest) -> operations.GetWorkflowResponse:
        r"""Get a flow by id
        Get a flow by id
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.GetWorkflowRequest, base_url, '/api/orchestration/workflows/{flowId}', request)
        
        
        client = self._security_client
        
        http_res = client.request('GET', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetWorkflowResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.GetWorkflowResponse])
                res.get_workflow_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def list_instances(self, request: operations.ListInstancesRequest) -> operations.ListInstancesResponse:
        r"""List instances of a workflow
        List instances of a workflow
        """
        base_url = self._server_url
        
        url = base_url.removesuffix('/') + '/api/orchestration/instances'
        
        query_params = utils.get_query_params(operations.ListInstancesRequest, request)
        
        client = self._security_client
        
        http_res = client.request('GET', url, params=query_params)
        content_type = http_res.headers.get('Content-Type')

        res = operations.ListInstancesResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[Any])
                res.list_runs_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def list_workflows(self) -> operations.ListWorkflowsResponse:
        r"""List registered workflows
        List registered workflows
        """
        base_url = self._server_url
        
        url = base_url.removesuffix('/') + '/api/orchestration/workflows'
        
        
        client = self._security_client
        
        http_res = client.request('GET', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.ListWorkflowsResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ListWorkflowsResponse])
                res.list_workflows_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def orchestrationget_server_info(self) -> operations.OrchestrationgetServerInfoResponse:
        r"""Get server info"""
        base_url = self._server_url
        
        url = base_url.removesuffix('/') + '/api/orchestration/_info'
        
        
        client = self._security_client
        
        http_res = client.request('GET', url)
        content_type = http_res.headers.get('Content-Type')

        res = operations.OrchestrationgetServerInfoResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ServerInfo])
                res.server_info = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def run_workflow(self, request: operations.RunWorkflowRequest) -> operations.RunWorkflowResponse:
        r"""Run workflow
        Run workflow
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.RunWorkflowRequest, base_url, '/api/orchestration/workflows/{workflowID}/instances', request)
        
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "request_body", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        query_params = utils.get_query_params(operations.RunWorkflowRequest, request)
        
        client = self._security_client
        
        http_res = client.request('POST', url, params=query_params, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.RunWorkflowResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 201:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.RunWorkflowResponse])
                res.run_workflow_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    def send_event(self, request: operations.SendEventRequest) -> operations.SendEventResponse:
        r"""Send an event to a running workflow
        Send an event to a running workflow
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.SendEventRequest, base_url, '/api/orchestration/instances/{instanceID}/events', request)
        
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "request_body", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        
        client = self._security_client
        
        http_res = client.request('POST', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.SendEventResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 204:
            pass
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.Error])
                res.error = out

        return res

    