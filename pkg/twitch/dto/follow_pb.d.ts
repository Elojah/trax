import * as jspb from 'google-protobuf'

import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_trax_pkg_twitch_follow_pb from '../../../../../../github.com/elojah/trax/pkg/twitch/follow_pb';


export class ListFollowReq extends jspb.Message {
  getAfter(): string;
  setAfter(value: string): ListFollowReq;

  getFirst(): string;
  setFirst(value: string): ListFollowReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFollowReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListFollowReq): ListFollowReq.AsObject;
  static serializeBinaryToWriter(message: ListFollowReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFollowReq;
  static deserializeBinaryFromReader(message: ListFollowReq, reader: jspb.BinaryReader): ListFollowReq;
}

export namespace ListFollowReq {
  export type AsObject = {
    after: string,
    first: string,
  }
}

export class ListFollowResp extends jspb.Message {
  getFollowsList(): Array<github_com_elojah_trax_pkg_twitch_follow_pb.Follow>;
  setFollowsList(value: Array<github_com_elojah_trax_pkg_twitch_follow_pb.Follow>): ListFollowResp;
  clearFollowsList(): ListFollowResp;
  addFollows(value?: github_com_elojah_trax_pkg_twitch_follow_pb.Follow, index?: number): github_com_elojah_trax_pkg_twitch_follow_pb.Follow;

  getTotal(): number;
  setTotal(value: number): ListFollowResp;

  getCursor(): string;
  setCursor(value: string): ListFollowResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFollowResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListFollowResp): ListFollowResp.AsObject;
  static serializeBinaryToWriter(message: ListFollowResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFollowResp;
  static deserializeBinaryFromReader(message: ListFollowResp, reader: jspb.BinaryReader): ListFollowResp;
}

export namespace ListFollowResp {
  export type AsObject = {
    followsList: Array<github_com_elojah_trax_pkg_twitch_follow_pb.Follow.AsObject>,
    total: number,
    cursor: string,
  }
}

