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

import { StageStatus } from '../models/StageStatus';
import { HttpFile } from '../http/http';

export class WorkflowInstance {
    'workflowID': string;
    'id': string;
    'createdAt': Date;
    'updatedAt': Date;
    'status'?: Array<StageStatus>;
    'terminated': boolean;
    'terminatedAt'?: Date;
    'error'?: string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "workflowID",
            "baseName": "workflowID",
            "type": "string",
            "format": ""
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string",
            "format": ""
        },
        {
            "name": "createdAt",
            "baseName": "createdAt",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "updatedAt",
            "baseName": "updatedAt",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "status",
            "baseName": "status",
            "type": "Array<StageStatus>",
            "format": ""
        },
        {
            "name": "terminated",
            "baseName": "terminated",
            "type": "boolean",
            "format": ""
        },
        {
            "name": "terminatedAt",
            "baseName": "terminatedAt",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "error",
            "baseName": "error",
            "type": "string",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WorkflowInstance.attributeTypeMap;
    }

    public constructor() {
    }
}

