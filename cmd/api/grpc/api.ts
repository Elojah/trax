// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/cmd/api/grpc/api.proto" (package "grpc", syntax proto3)
// tslint:disable
import { Profile } from "../../../internal/user/user";
import { Empty } from "../../../pkg/pbtypes/empty";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
/**
 * @generated ServiceType for protobuf service grpc.API
 */
export const API = new ServiceType("grpc.API", [
    { name: "Ping", options: {}, I: Empty, O: Empty },
    { name: "FetchProfile", options: {}, I: Empty, O: Profile }
]);
