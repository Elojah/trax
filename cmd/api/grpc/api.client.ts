// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/cmd/api/grpc/api.proto" (package "grpc", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { API } from "./api";
import type { ListEntityResp } from "../../../internal/user/dto/entity";
import type { ListEntityReq } from "../../../internal/user/dto/entity";
import type { Entity } from "../../../internal/user/entity";
import type { CreateEntityReq } from "../../../internal/user/dto/entity";
import type { UpdateProfileReq } from "../../../internal/user/dto/profile";
import type { Profile } from "../../../internal/user/user";
import type { FetchProfileReq } from "../../../internal/user/dto/profile";
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
     * @generated from protobuf rpc: FetchProfile(dto.FetchProfileReq) returns (user.Profile);
     */
    fetchProfile(input: FetchProfileReq, options?: RpcOptions): UnaryCall<FetchProfileReq, Profile>;
    /**
     * @generated from protobuf rpc: UpdateProfile(dto.UpdateProfileReq) returns (user.Profile);
     */
    updateProfile(input: UpdateProfileReq, options?: RpcOptions): UnaryCall<UpdateProfileReq, Profile>;
    /**
     * Entity
     *
     * @generated from protobuf rpc: CreateEntity(dto.CreateEntityReq) returns (user.Entity);
     */
    createEntity(input: CreateEntityReq, options?: RpcOptions): UnaryCall<CreateEntityReq, Entity>;
    /**
     * rpc UpdateEntity(user.Entity) returns (user.Entity); // TODO: IMPLEMENT
     *
     * @generated from protobuf rpc: ListEntity(dto.ListEntityReq) returns (dto.ListEntityResp);
     */
    listEntity(input: ListEntityReq, options?: RpcOptions): UnaryCall<ListEntityReq, ListEntityResp>;
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
     * @generated from protobuf rpc: FetchProfile(dto.FetchProfileReq) returns (user.Profile);
     */
    fetchProfile(input: FetchProfileReq, options?: RpcOptions): UnaryCall<FetchProfileReq, Profile> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<FetchProfileReq, Profile>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: UpdateProfile(dto.UpdateProfileReq) returns (user.Profile);
     */
    updateProfile(input: UpdateProfileReq, options?: RpcOptions): UnaryCall<UpdateProfileReq, Profile> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<UpdateProfileReq, Profile>("unary", this._transport, method, opt, input);
    }
    /**
     * Entity
     *
     * @generated from protobuf rpc: CreateEntity(dto.CreateEntityReq) returns (user.Entity);
     */
    createEntity(input: CreateEntityReq, options?: RpcOptions): UnaryCall<CreateEntityReq, Entity> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<CreateEntityReq, Entity>("unary", this._transport, method, opt, input);
    }
    /**
     * rpc UpdateEntity(user.Entity) returns (user.Entity); // TODO: IMPLEMENT
     *
     * @generated from protobuf rpc: ListEntity(dto.ListEntityReq) returns (dto.ListEntityResp);
     */
    listEntity(input: ListEntityReq, options?: RpcOptions): UnaryCall<ListEntityReq, ListEntityResp> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListEntityReq, ListEntityResp>("unary", this._transport, method, opt, input);
    }
}
