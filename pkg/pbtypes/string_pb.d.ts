import * as jspb from 'google-protobuf'

import * as github_com_elojah_trax_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/trax/pkg/gogoproto/gogo_pb';


export class String extends jspb.Message {
  getValue(): string;
  setValue(value: string): String;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): String.AsObject;
  static toObject(includeInstance: boolean, msg: String): String.AsObject;
  static serializeBinaryToWriter(message: String, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): String;
  static deserializeBinaryFromReader(message: String, reader: jspb.BinaryReader): String;
}

export namespace String {
  export type AsObject = {
    value: string,
  }
}

