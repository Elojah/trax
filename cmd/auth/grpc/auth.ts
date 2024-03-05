// @generated by protobuf-ts 2.9.3 with parameter optimize_code_size
// @generated from protobuf file "github.com/elojah/trax/cmd/auth/grpc/auth.proto" (package "grpc", syntax proto3)
// tslint:disable
import { Empty } from "../../../pkg/pbtypes/empty";
import { SigninResp } from "../../../internal/user/dto/user";
import { String$ } from "../../../pkg/pbtypes/string";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
/**
 * @generated ServiceType for protobuf service grpc.Auth
 */
export const Auth = new ServiceType("grpc.Auth", [
    { name: "SigninGoogle", options: {}, I: String$, O: SigninResp },
    { name: "SigninTwitch", options: {}, I: String$, O: SigninResp },
    { name: "RefreshToken", options: {}, I: String$, O: SigninResp },
    { name: "Ping", options: {}, I: Empty, O: Empty }
]);
