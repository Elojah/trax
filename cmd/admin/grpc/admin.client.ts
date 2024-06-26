// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/cmd/admin/grpc/admin.proto" (package "grpc", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { Admin } from "./admin";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { Empty } from "../../../pkg/pbtypes/empty";
import type { String$ } from "../../../pkg/pbtypes/string";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * #MARK:## TECHNICAL ADMIN ###
 *
 * @generated from protobuf service grpc.Admin
 */
export interface IAdminClient {
    /**
     * DB migrations
     *
     * @generated from protobuf rpc: MigrateUp(pbtypes.String) returns (pbtypes.Empty);
     */
    migrateUp(input: String$, options?: RpcOptions): UnaryCall<String$, Empty>;
    /**
     * Cookie secure management
     *
     * @generated from protobuf rpc: RotateCookieKeys(pbtypes.Empty) returns (pbtypes.Empty);
     */
    rotateCookieKeys(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty>;
    /**
     * Ping
     *
     * @generated from protobuf rpc: Ping(pbtypes.Empty) returns (pbtypes.Empty);
     */
    ping(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty>;
}
/**
 * #MARK:## TECHNICAL ADMIN ###
 *
 * @generated from protobuf service grpc.Admin
 */
export class AdminClient implements IAdminClient, ServiceInfo {
    typeName = Admin.typeName;
    methods = Admin.methods;
    options = Admin.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * DB migrations
     *
     * @generated from protobuf rpc: MigrateUp(pbtypes.String) returns (pbtypes.Empty);
     */
    migrateUp(input: String$, options?: RpcOptions): UnaryCall<String$, Empty> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<String$, Empty>("unary", this._transport, method, opt, input);
    }
    /**
     * Cookie secure management
     *
     * @generated from protobuf rpc: RotateCookieKeys(pbtypes.Empty) returns (pbtypes.Empty);
     */
    rotateCookieKeys(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<Empty, Empty>("unary", this._transport, method, opt, input);
    }
    /**
     * Ping
     *
     * @generated from protobuf rpc: Ping(pbtypes.Empty) returns (pbtypes.Empty);
     */
    ping(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<Empty, Empty>("unary", this._transport, method, opt, input);
    }
}
