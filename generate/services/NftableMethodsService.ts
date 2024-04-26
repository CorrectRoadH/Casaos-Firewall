/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { BaseResponse } from '../models/BaseResponse';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class NftableMethodsService {

    /**
     * Get Version
     * (TODO)
     * @returns any OK
     * @throws ApiError
     */
    public static getVersion(): CancelablePromise<(BaseResponse & {
        data?: string;
    })> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/version',
            errors: {
                503: `Service Unavailable`,
            },
        });
    }

}
