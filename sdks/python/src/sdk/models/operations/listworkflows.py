"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
import requests as requests_http
from ..shared import error as shared_error
from ..shared import listworkflowsresponse as shared_listworkflowsresponse
from typing import Optional


@dataclasses.dataclass
class ListWorkflowsResponse:
    
    content_type: str = dataclasses.field()  
    status_code: int = dataclasses.field()  
    error: Optional[shared_error.Error] = dataclasses.field(default=None)
    r"""General error"""  
    list_workflows_response: Optional[shared_listworkflowsresponse.ListWorkflowsResponse] = dataclasses.field(default=None)
    r"""List of workflows"""  
    raw_response: Optional[requests_http.Response] = dataclasses.field(default=None)  
    