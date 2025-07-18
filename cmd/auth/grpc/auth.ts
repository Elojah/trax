// @generated by protobuf-ts 2.11.1 with parameter generate_dependencies,optimize_speed,long_type_string
// @generated from protobuf file "cmd/auth/grpc/auth.proto" (package "grpc", syntax proto3)
// tslint:disable
import { String$ } from "../../../pkg/pbtypes/string";
import { SigninResp } from "../../../internal/user/dto/user";
import { SigninReq } from "../../../internal/user/dto/user";
import { Empty } from "../../../pkg/pbtypes/empty";
import { SignupReq } from "../../../internal/user/dto/user";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
/**
 * @generated ServiceType for protobuf service grpc.Auth
 */
export const Auth = new ServiceType("grpc.Auth", [
    { name: "Signup", options: {}, I: SignupReq, O: Empty },
    { name: "Signin", options: {}, I: SigninReq, O: SigninResp },
    { name: "SigninGoogle", options: {}, I: String$, O: SigninResp },
    { name: "RefreshToken", options: {}, I: String$, O: SigninResp },
    { name: "Ping", options: {}, I: Empty, O: Empty }
]);
