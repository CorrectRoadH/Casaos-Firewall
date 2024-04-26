/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BaseResponse } from '../models/BaseResponse';
import type { Port } from '../models/Port';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class FirewallMethodsService {

    /**
     * Get State
     * Get firewall state
     * @returns any OK
     * @throws ApiError
     */
    public static getState(): CancelablePromise<(BaseResponse & {
        data?: boolean;
    })> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/firewall',
            errors: {
                503: `Service Unavailable`,
            },
        });
    }

    /**
     * Get Opened Ports
     * (TODO)
     * @returns any OK
     * @throws ApiError
     */
    public static getOpenedPorts(): CancelablePromise<(BaseResponse & {
        data?: Array<Port>;
    })> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/port',
            errors: {
                503: `Service Unavailable`,
            },
        });
    }

    /**
     * Open or Close a Port
     * (TODO)
     * @param requestBody
     * @returns any OK
     * @throws ApiError
     */
    public static openOrClosePort(
        requestBody?: Port,
    ): CancelablePromise<(BaseResponse & {
        data?: string;
    })> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/port',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                403: `Bad Request`,
                503: `Service Unavailable`,
            },
        });
    }

}
