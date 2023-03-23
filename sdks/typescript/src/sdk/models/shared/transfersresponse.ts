/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose, Type } from "class-transformer";

export class TransfersResponseData extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "amount" })
  amount?: number;

  @SpeakeasyMetadata()
  @Expose({ name: "asset" })
  asset?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "currency" })
  currency?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "destination" })
  destination?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "error" })
  error?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "id" })
  id?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "source" })
  source?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "status" })
  status?: string;
}

/**
 * OK
 */
export class TransfersResponse extends SpeakeasyBase {
  @SpeakeasyMetadata({ elemType: TransfersResponseData })
  @Expose({ name: "data" })
  @Type(() => TransfersResponseData)
  data?: TransfersResponseData[];
}
