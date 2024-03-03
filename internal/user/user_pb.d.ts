import * as jspb from 'google-protobuf'

import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb';


export class U extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): U;

  getEmail(): string;
  setEmail(value: string): U;

  getPassword(): string;
  setPassword(value: string): U;

  getTwitchid(): string;
  setTwitchid(value: string): U;

  getGoogleid(): string;
  setGoogleid(value: string): U;

  getCreatedat(): number;
  setCreatedat(value: number): U;

  getUpdatedat(): number;
  setUpdatedat(value: number): U;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): U.AsObject;
  static toObject(includeInstance: boolean, msg: U): U.AsObject;
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
  setUserid(value: Uint8Array | string): Profile;

  getFirstname(): string;
  setFirstname(value: string): Profile;

  getLastname(): string;
  setLastname(value: string): Profile;

  getCreatedat(): number;
  setCreatedat(value: number): Profile;

  getUpdatedat(): number;
  setUpdatedat(value: number): Profile;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Profile.AsObject;
  static toObject(includeInstance: boolean, msg: Profile): Profile.AsObject;
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

