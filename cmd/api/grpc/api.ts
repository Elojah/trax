// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/cmd/api/grpc/api.proto" (package "grpc", syntax proto3)
// tslint:disable
import { UpdateRoleReq } from "../../../internal/user/dto/role";
import { RolePermission } from "../../../internal/user/dto/role";
import { CreateRoleReq } from "../../../internal/user/dto/role";
import { ListRoleResp } from "../../../internal/user/dto/role";
import { ListRoleReq } from "../../../internal/user/dto/role";
import { DeleteEntityResp } from "../../../internal/user/dto/entity";
import { DeleteEntityReq } from "../../../internal/user/dto/entity";
import { UpdateEntityReq } from "../../../internal/user/dto/entity";
import { CreateEntityReq } from "../../../internal/user/dto/entity";
import { Entity } from "../../../internal/user/entity";
import { FetchEntityReq } from "../../../internal/user/dto/entity";
import { ListEntityResp } from "../../../internal/user/dto/entity";
import { ListEntityReq } from "../../../internal/user/dto/entity";
import { UpdateUserReq } from "../../../internal/user/dto/user";
import { U } from "../../../internal/user/user";
import { FetchUserReq } from "../../../internal/user/dto/user";
import { ListUserResp } from "../../../internal/user/dto/user";
import { ListUserReq } from "../../../internal/user/dto/user";
import { Empty } from "../../../pkg/pbtypes/empty";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
/**
 * @generated ServiceType for protobuf service grpc.API
 */
export const API = new ServiceType("grpc.API", [
    { name: "Ping", options: {}, I: Empty, O: Empty },
    { name: "ListUser", options: {}, I: ListUserReq, O: ListUserResp },
    { name: "FetchUser", options: {}, I: FetchUserReq, O: U },
    { name: "UpdateUser", options: {}, I: UpdateUserReq, O: U },
    { name: "ListEntity", options: {}, I: ListEntityReq, O: ListEntityResp },
    { name: "FetchEntity", options: {}, I: FetchEntityReq, O: Entity },
    { name: "CreateEntity", options: {}, I: CreateEntityReq, O: Entity },
    { name: "UpdateEntity", options: {}, I: UpdateEntityReq, O: Entity },
    { name: "DeleteEntity", options: {}, I: DeleteEntityReq, O: DeleteEntityResp },
    { name: "ListRole", options: {}, I: ListRoleReq, O: ListRoleResp },
    { name: "CreateRole", options: {}, I: CreateRoleReq, O: RolePermission },
    { name: "UpdateRole", options: {}, I: UpdateRoleReq, O: RolePermission }
]);
