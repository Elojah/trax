// package: user
// file: github.com/elojah/trax/internal/user/user.proto

import * as jspb from "google-protobuf";
import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from "../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb";

export class U extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getEmail(): string;
  setEmail(value: string): void;

  getPassword(): string;
  setPassword(value: string): void;

  getTwitchid(): string;
  setTwitchid(value: string): void;

  getGoogleid(): string;
  setGoogleid(value: string): void;

  getCreatedat(): number;
  setCreatedat(value: number): void;

  getUpdatedat(): number;
  setUpdatedat(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): U.AsObject;
  static toObject(includeInstance: boolean, msg: U): U.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: U, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): U;
  static deserializeBinaryFromReader(message: U, reader: jspb.BinaryReader): U;
}

export namespace U {
  export type AsObject = {
    id: Uint8Array | string,
    email: string,
    password: string,
    twitchid: string,
    googleid: string,
    createdat: number,
    updatedat: number,
  }
}

export class Profile extends jspb.Message {
  getUserid(): Uint8Array | string;
  getUserid_asU8(): Uint8Array;
  getUserid_asB64(): string;
  setUserid(value: Uint8Array | string): void;

  getFirstname(): string;
  setFirstname(value: string): void;

  getLastname(): string;
  setLastname(value: string): void;

  getCreatedat(): number;
  setCreatedat(value: number): void;

  getUpdatedat(): number;
  setUpdatedat(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Profile.AsObject;
  static toObject(includeInstance: boolean, msg: Profile): Profile.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Profile, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Profile;
  static deserializeBinaryFromReader(message: Profile, reader: jspb.BinaryReader): Profile;
}

export namespace Profile {
  export type AsObject = {
    userid: Uint8Array | string,
    firstname: string,
    lastname: string,
    createdat: number,
    updatedat: number,
  }
}

