/**
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * OpenAPI spec version: 1.0.0-20230225
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { Monetary } from '../models/Monetary';
import { StageSendDestination } from '../models/StageSendDestination';
import { StageSendSource } from '../models/StageSendSource';
import { HttpFile } from '../http/http';

export class StageSend {
    'amount'?: Monetary;
    'destination'?: StageSendDestination;
    'source'?: StageSendSource;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "amount",
            "baseName": "amount",
            "type": "Monetary",
            "format": ""
        },
        {
            "name": "destination",
            "baseName": "destination",
            "type": "StageSendDestination",
            "format": ""
        },
        {
            "name": "source",
            "baseName": "source",
            "type": "StageSendSource",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return StageSend.attributeTypeMap;
    }

    public constructor() {
    }
}

