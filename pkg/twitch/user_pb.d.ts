import * as jspb from 'google-protobuf'

import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb';


export class User extends jspb.Message {
  getId(): string;
  setId(value: string): User;

  getLogin(): string;
  setLogin(value: string): User;

  getDisplayname(): string;
  setDisplayname(value: string): User;

  getBroadcastertype(): string;
  setBroadcastertype(value: string): User;

  getDescription(): string;
  setDescription(value: string): User;

  getProfileimageurl(): string;
  setProfileimageurl(value: string): User;

  getOfflineimageurl(): string;
  setOfflineimageurl(value: string): User;

  getViewcount(): number;
  setViewcount(value: number): User;

  getEmail(): string;
  setEmail(value: string): User;

  getCreatedat(): string;
  setCreatedat(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
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

