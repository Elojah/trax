// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/cmd/admin/grpc/admin.proto" (package "grpc", syntax proto3)
// tslint:disable
import { Empty } from "../../../pkg/pbtypes/empty";
import { String$ } from "../../../pkg/pbtypes/string";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
/**
 * @generated ServiceType for protobuf service grpc.Admin
 */
export const Admin = new ServiceType("grpc.Admin", [
    { name: "MigrateUp", options: {}, I: String$, O: Empty },
    { name: "RotateCookieKeys", options: {}, I: Empty, O: Empty },
    { name: "Ping", options: {}, I: Empty, O: Empty }
]);
