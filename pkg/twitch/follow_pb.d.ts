// package: twitch
// file: github.com/elojah/trax/pkg/twitch/follow.proto

import * as jspb from "google-protobuf";
import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from "../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb";

export class Follow extends jspb.Message {
  getFromid(): string;
  setFromid(value: string): void;

  getFromlogin(): string;
  setFromlogin(value: string): void;

  getFromname(): string;
  setFromname(value: string): void;

  getToid(): string;
  setToid(value: string): void;

  getTologin(): string;
  setTologin(value: string): void;

  getToname(): string;
  setToname(value: string): void;

  getFollowedat(): string;
  setFollowedat(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Follow.AsObject;
  static toObject(includeInstance: boolean, msg: Follow): Follow.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Follow, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Follow;
  static deserializeBinaryFromReader(message: Follow, reader: jspb.BinaryReader): Follow;
}

export namespace Follow {
  export type AsObject = {
    fromid: string,
    fromlogin: string,
    fromname: string,
    toid: string,
    tologin: string,
    toname: string,
    followedat: string,
  }
}

