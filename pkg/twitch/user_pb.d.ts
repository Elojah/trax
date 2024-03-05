// package: twitch
// file: github.com/elojah/trax/pkg/twitch/user.proto

import * as jspb from "google-protobuf";
import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from "../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb";

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getLogin(): string;
  setLogin(value: string): void;

  getDisplayname(): string;
  setDisplayname(value: string): void;

  getBroadcastertype(): string;
  setBroadcastertype(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  getProfileimageurl(): string;
  setProfileimageurl(value: string): void;

  getOfflineimageurl(): string;
  setOfflineimageurl(value: string): void;

  getViewcount(): number;
  setViewcount(value: number): void;

  getEmail(): string;
  setEmail(value: string): void;

  getCreatedat(): string;
  setCreatedat(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    login: string,
    displayname: string,
    broadcastertype: string,
    description: string,
    profileimageurl: string,
    offlineimageurl: string,
    viewcount: number,
    email: string,
    createdat: string,
  }
}

