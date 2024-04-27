// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/cmd/api/grpc/api.proto" (package "grpc", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { API } from "./api";
import type { ListRoleResp } from "../../../internal/user/dto/role";
import type { ListRoleReq } from "../../../internal/user/dto/role";
import type { Role } from "../../../internal/user/role";
import type { CreateRoleReq } from "../../../internal/user/dto/role";
import type { UpdateEntityReq } from "../../../internal/user/dto/entity";
import type { FetchEntityReq } from "../../../internal/user/dto/entity";
import type { ListEntityResp } from "../../../internal/user/dto/entity";
import type { ListEntityReq } from "../../../internal/user/dto/entity";
import type { Entity } from "../../../internal/user/entity";
import type { CreateEntityReq } from "../../../internal/user/dto/entity";
import type { ListUserResp } from "../../../internal/user/dto/user";
import type { ListUserReq } from "../../../internal/user/dto/user";
import type { UpdateUserReq } from "../../../internal/user/dto/user";
import type { U } from "../../../internal/user/user";
import type { FetchUserReq } from "../../../internal/user/dto/user";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { Empty } from "../../../pkg/pbtypes/empty";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service grpc.API
 */
export interface IAPIClient {
    /**
     * Ping
     *
     * @generated from protobuf rpc: Ping(pbtypes.Empty) returns (pbtypes.Empty);
     */
    ping(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty>;
    /**
     * User
     *
     * @generated from protobuf rpc: FetchUser(dto.FetchUserReq) returns (user.U);
     */
    fetchUser(input: FetchUserReq, options?: RpcOptions): UnaryCall<FetchUserReq, U>;
    /**
     * @generated from protobuf rpc: UpdateUser(dto.UpdateUserReq) returns (user.U);
     */
    updateUser(input: UpdateUserReq, options?: RpcOptions): UnaryCall<UpdateUserReq, U>;
    /**
     * @generated from protobuf rpc: ListUser(dto.ListUserReq) returns (dto.ListUserResp);
     */
    listUser(input: ListUserReq, options?: RpcOptions): UnaryCall<ListUserReq, ListUserResp>;
    /**
     * Entity
     *
     * @generated from protobuf rpc: CreateEntity(dto.CreateEntityReq) returns (user.Entity);
     */
    createEntity(input: CreateEntityReq, options?: RpcOptions): UnaryCall<CreateEntityReq, Entity>;
    /**
     * @generated from protobuf rpc: ListEntity(dto.ListEntityReq) returns (dto.ListEntityResp);
     */
    listEntity(input: ListEntityReq, options?: RpcOptions): UnaryCall<ListEntityReq, ListEntityResp>;
    /**
     * @generated from protobuf rpc: FetchEntity(dto.FetchEntityReq) returns (user.Entity);
     */
    fetchEntity(input: FetchEntityReq, options?: RpcOptions): UnaryCall<FetchEntityReq, Entity>;
    /**
     * @generated from protobuf rpc: UpdateEntity(dto.UpdateEntityReq) returns (user.Entity);
     */
    updateEntity(input: UpdateEntityReq, options?: RpcOptions): UnaryCall<UpdateEntityReq, Entity>;
    /**
     * Roles
     *
     * @generated from protobuf rpc: CreateRole(dto.CreateRoleReq) returns (user.Role);
     */
    createRole(input: CreateRoleReq, options?: RpcOptions): UnaryCall<CreateRoleReq, Role>;
    /**
     * @generated from protobuf rpc: ListRole(dto.ListRoleReq) returns (dto.ListRoleResp);
     */
    listRole(input: ListRoleReq, options?: RpcOptions): UnaryCall<ListRoleReq, ListRoleResp>;
}
/**
 * @generated from protobuf service grpc.API
 */
export class APIClient implements IAPIClient, ServiceInfo {
    typeName = API.typeName;
    methods = API.methods;
    options = API.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * Ping
     *
     * @generated from protobuf rpc: Ping(pbtypes.Empty) returns (pbtypes.Empty);
     */
    ping(input: Empty, options?: RpcOptions): UnaryCall<Empty, Empty> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<Empty, Empty>("unary", this._transport, method, opt, input);
    }
    /**
     * User
     *
     * @generated from protobuf rpc: FetchUser(dto.FetchUserReq) returns (user.U);
     */
    fetchUser(input: FetchUserReq, options?: RpcOptions): UnaryCall<FetchUserReq, U> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<FetchUserReq, U>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateUser(dto.UpdateUserReq) returns (user.U);
     */
    updateUser(input: UpdateUserReq, options?: RpcOptions): UnaryCall<UpdateUserReq, U> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateUserReq, U>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ListUser(dto.ListUserReq) returns (dto.ListUserResp);
     */
    listUser(input: ListUserReq, options?: RpcOptions): UnaryCall<ListUserReq, ListUserResp> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListUserReq, ListUserResp>("unary", this._transport, method, opt, input);
    }
    /**
     * Entity
     *
     * @generated from protobuf rpc: CreateEntity(dto.CreateEntityReq) returns (user.Entity);
     */
    createEntity(input: CreateEntityReq, options?: RpcOptions): UnaryCall<CreateEntityReq, Entity> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateEntityReq, Entity>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ListEntity(dto.ListEntityReq) returns (dto.ListEntityResp);
     */
    listEntity(input: ListEntityReq, options?: RpcOptions): UnaryCall<ListEntityReq, ListEntityResp> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListEntityReq, ListEntityResp>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: FetchEntity(dto.FetchEntityReq) returns (user.Entity);
     */
    fetchEntity(input: FetchEntityReq, options?: RpcOptions): UnaryCall<FetchEntityReq, Entity> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<FetchEntityReq, Entity>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateEntity(dto.UpdateEntityReq) returns (user.Entity);
     */
    updateEntity(input: UpdateEntityReq, options?: RpcOptions): UnaryCall<UpdateEntityReq, Entity> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateEntityReq, Entity>("unary", this._transport, method, opt, input);
    }
    /**
     * Roles
     *
     * @generated from protobuf rpc: CreateRole(dto.CreateRoleReq) returns (user.Role);
     */
    createRole(input: CreateRoleReq, options?: RpcOptions): UnaryCall<CreateRoleReq, Role> {
        const method = this.methods[8], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateRoleReq, Role>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ListRole(dto.ListRoleReq) returns (dto.ListRoleResp);
     */
    listRole(input: ListRoleReq, options?: RpcOptions): UnaryCall<ListRoleReq, ListRoleResp> {
        const method = this.methods[9], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListRoleReq, ListRoleResp>("unary", this._transport, method, opt, input);
    }
}
