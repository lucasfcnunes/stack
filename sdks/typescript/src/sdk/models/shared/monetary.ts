/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export class Monetary extends SpeakeasyBase {
  /**
   * The amount of the monetary value.
   */
  @SpeakeasyMetadata()
  @Expose({ name: "amount" })
  amount: number;

  /**
   * The asset of the monetary value.
   */
  @SpeakeasyMetadata()
  @Expose({ name: "asset" })
  asset: string;
}
