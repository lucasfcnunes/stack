/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";
import { Type } from "class-transformer";

export class ReadStatsRequest extends SpeakeasyBase {
  /**
   * name of the ledger
   */
  @SpeakeasyMetadata({
    data: "pathParam, style=simple;explode=false;name=ledger",
  })
  ledger: string;
}

export class ReadStatsResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  /**
   * Error
   */
  @SpeakeasyMetadata()
  errorResponse?: shared.ErrorResponse;

  /**
   * OK
   */
  @SpeakeasyMetadata()
  statsResponse?: shared.StatsResponse;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  rawResponse?: AxiosResponse;
}
