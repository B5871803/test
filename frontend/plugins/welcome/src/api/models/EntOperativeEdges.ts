/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntOperativerecord,
    EntOperativerecordFromJSON,
    EntOperativerecordFromJSONTyped,
    EntOperativerecordToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOperativeEdges
 */
export interface EntOperativeEdges {
    /**
     * Fromoperative holds the value of the fromoperative edge.
     * @type {Array<EntOperativerecord>}
     * @memberof EntOperativeEdges
     */
    fromoperative?: Array<EntOperativerecord>;
}

export function EntOperativeEdgesFromJSON(json: any): EntOperativeEdges {
    return EntOperativeEdgesFromJSONTyped(json, false);
}

export function EntOperativeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOperativeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fromoperative': !exists(json, 'fromoperative') ? undefined : ((json['fromoperative'] as Array<any>).map(EntOperativerecordFromJSON)),
    };
}

export function EntOperativeEdgesToJSON(value?: EntOperativeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fromoperative': value.fromoperative === undefined ? undefined : ((value.fromoperative as Array<any>).map(EntOperativerecordToJSON)),
    };
}

