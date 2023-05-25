import type { BaseResponse } from '../models/BaseResponse';
import type { Rule } from '../models/Rule';
import type { CancelablePromise } from '../core/CancelablePromise';
export declare class NftableMethodsService {
    /**
     * Get Rules
     * (TODO)
     * @returns any OK
     * @throws ApiError
     */
    static getRules(): CancelablePromise<(BaseResponse & {
        data?: Array<Rule>;
    })>;
}
