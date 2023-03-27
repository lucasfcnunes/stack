/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";

export class ListInstancesRequest extends SpeakeasyBase {
  /**
   * Filter running instances
   */
  @SpeakeasyMetadata({
    data: "queryParam, style=form;explode=true;name=running",
  })
  running?: boolean;

  /**
   * A workflow id
   */
  @SpeakeasyMetadata({
    data: "queryParam, style=form;explode=true;name=workflowID",
  })
  workflowID?: string;
}

export class ListInstancesResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  /**
   * General error
   */
  @SpeakeasyMetadata()
  error?: shared.ErrorT;

  /**
   * List of workflow instances
   */
  @SpeakeasyMetadata()
  listRunsResponse?: any;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  rawResponse?: AxiosResponse;
}
