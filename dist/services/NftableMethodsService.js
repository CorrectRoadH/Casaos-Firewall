"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.NftableMethodsService = void 0;
const OpenAPI_1 = require("../core/OpenAPI");
const request_1 = require("../core/request");
class NftableMethodsService {
    /**
     * Get Rules
     * (TODO)
     * @returns any OK
     * @throws ApiError
     */
    static getRules() {
        return (0, request_1.request)(OpenAPI_1.OpenAPI, {
            method: 'GET',
            url: '/nftables',
            errors: {
                503: `Service Unavailable`,
            },
        });
    }
}
exports.NftableMethodsService = NftableMethodsService;
