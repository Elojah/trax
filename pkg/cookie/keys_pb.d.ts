// package: cookie
// file: github.com/elojah/trax/pkg/cookie/keys.proto

import * as jspb from "google-protobuf";
import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from "../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb";

export class Keys extends jspb.Message {
  getHash(): Uint8Array | string;
  getHash_asU8(): Uint8Array;
  getHash_asB64(): string;
  setHash(value: Uint8Array | string): void;

  getBlock(): Uint8Array | string;
  getBlock_asU8(): Uint8Array;
  getBlock_asB64(): string;
  setBlock(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Keys.AsObject;
  static toObject(includeInstance: boolean, msg: Keys): Keys.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Keys, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Keys;
  static deserializeBinaryFromReader(message: Keys, reader: jspb.BinaryReader): Keys;
}

export namespace Keys {
  export type AsObject = {
    hash: Uint8Array | string,
    block: Uint8Array | string,
  }
}

