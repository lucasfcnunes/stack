"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

from __future__ import annotations
import dataclasses
from ..shared import workflowinstancehistorystage as shared_workflowinstancehistorystage
from dataclasses_json import Undefined, dataclass_json
from sdk import utils


@dataclass_json(undefined=Undefined.EXCLUDE)
@dataclasses.dataclass
class GetWorkflowInstanceHistoryStageResponse:
    r"""The workflow instance stage history"""
    
    data: list[shared_workflowinstancehistorystage.WorkflowInstanceHistoryStage] = dataclasses.field(metadata={'dataclasses_json': { 'letter_case': utils.get_field_name('data') }})  
    