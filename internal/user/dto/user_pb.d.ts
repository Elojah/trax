// package: dto
// file: github.com/elojah/trax/internal/user/dto/user.proto

import * as jspb from "google-protobuf";
import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from "../../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb";

export class SigninResp extends jspb.Message {
  getAccesstoken(): string;
  setAccesstoken(value: string): void;

  getRefreshtoken(): string;
  setRefreshtoken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SigninResp.AsObject;
  static toObject(includeInstance: boolean, msg: SigninResp): SigninResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SigninResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SigninResp;
  static deserializeBinaryFromReader(message: SigninResp, reader: jspb.BinaryReader): SigninResp;
}

export namespace SigninResp {
  export type AsObject = {
    accesstoken: string,
    refreshtoken: string,
  }
}

