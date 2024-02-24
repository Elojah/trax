import * as jspb from 'google-protobuf'

import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb';


export class Follow extends jspb.Message {
  getFromid(): string;
  setFromid(value: string): Follow;

  getFromlogin(): string;
  setFromlogin(value: string): Follow;

  getFromname(): string;
  setFromname(value: string): Follow;

  getToid(): string;
  setToid(value: string): Follow;

  getTologin(): string;
  setTologin(value: string): Follow;

  getToname(): string;
  setToname(value: string): Follow;

  getFollowedat(): string;
  setFollowedat(value: string): Follow;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Follow.AsObject;
  static toObject(includeInstance: boolean, msg: Follow): Follow.AsObject;
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

