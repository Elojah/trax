// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/internal/user/dto/profile.proto" (package "dto", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { String$ } from "../../../pkg/pbtypes/string";
/**
 * @generated from protobuf message dto.FetchProfileReq
 */
export interface FetchProfileReq {
    /**
     * @generated from protobuf field: bool Me = 1 [json_name = "Me"];
     */
    me: boolean;
    /**
     * @generated from protobuf field: bytes UserID = 2 [json_name = "UserID"];
     */
    userID: Uint8Array;
    /**
     * @generated from protobuf field: bytes EntityID = 3 [json_name = "EntityID"];
     */
    entityID: Uint8Array;
}
/**
 * @generated from protobuf message dto.UpdateProfileReq
 */
export interface UpdateProfileReq {
    /**
     * @generated from protobuf field: pbtypes.String Firstname = 1 [json_name = "Firstname"];
     */
    firstname?: String$;
    /**
     * @generated from protobuf field: pbtypes.String Lastname = 2 [json_name = "Lastname"];
     */
    lastname?: String$;
    /**
     * @generated from protobuf field: pbtypes.String AvatarURL = 3 [json_name = "AvatarURL"];
     */
    avatarURL?: String$;
}
// @generated message type with reflection information, may provide speed optimized methods
class FetchProfileReq$Type extends MessageType<FetchProfileReq> {
    constructor() {
        super("dto.FetchProfileReq", [
            { no: 1, name: "Me", kind: "scalar", jsonName: "Me", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "UserID", kind: "scalar", jsonName: "UserID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 3, name: "EntityID", kind: "scalar", jsonName: "EntityID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<FetchProfileReq>): FetchProfileReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.me = false;
        message.userID = new Uint8Array(0);
        message.entityID = new Uint8Array(0);
        if (value !== undefined)
            reflectionMergePartial<FetchProfileReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: FetchProfileReq): FetchProfileReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool Me = 1 [json_name = "Me"];*/ 1:
                    message.me = reader.bool();
                    break;
                case /* bytes UserID = 2 [json_name = "UserID"];*/ 2:
                    message.userID = reader.bytes();
                    break;
                case /* bytes EntityID = 3 [json_name = "EntityID"];*/ 3:
                    message.entityID = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: FetchProfileReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool Me = 1 [json_name = "Me"]; */
        if (message.me !== false)
            writer.tag(1, WireType.Varint).bool(message.me);
        /* bytes UserID = 2 [json_name = "UserID"]; */
        if (message.userID.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.userID);
        /* bytes EntityID = 3 [json_name = "EntityID"]; */
        if (message.entityID.length)
            writer.tag(3, WireType.LengthDelimited).bytes(message.entityID);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.FetchProfileReq
 */
export const FetchProfileReq = new FetchProfileReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateProfileReq$Type extends MessageType<UpdateProfileReq> {
    constructor() {
        super("dto.UpdateProfileReq", [
            { no: 1, name: "Firstname", kind: "message", jsonName: "Firstname", T: () => String$ },
            { no: 2, name: "Lastname", kind: "message", jsonName: "Lastname", T: () => String$ },
            { no: 3, name: "AvatarURL", kind: "message", jsonName: "AvatarURL", T: () => String$ }
        ]);
    }
    create(value?: PartialMessage<UpdateProfileReq>): UpdateProfileReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpdateProfileReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateProfileReq): UpdateProfileReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* pbtypes.String Firstname = 1 [json_name = "Firstname"];*/ 1:
                    message.firstname = String$.internalBinaryRead(reader, reader.uint32(), options, message.firstname);
                    break;
                case /* pbtypes.String Lastname = 2 [json_name = "Lastname"];*/ 2:
                    message.lastname = String$.internalBinaryRead(reader, reader.uint32(), options, message.lastname);
                    break;
                case /* pbtypes.String AvatarURL = 3 [json_name = "AvatarURL"];*/ 3:
                    message.avatarURL = String$.internalBinaryRead(reader, reader.uint32(), options, message.avatarURL);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UpdateProfileReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* pbtypes.String Firstname = 1 [json_name = "Firstname"]; */
        if (message.firstname)
            String$.internalBinaryWrite(message.firstname, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* pbtypes.String Lastname = 2 [json_name = "Lastname"]; */
        if (message.lastname)
            String$.internalBinaryWrite(message.lastname, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* pbtypes.String AvatarURL = 3 [json_name = "AvatarURL"]; */
        if (message.avatarURL)
            String$.internalBinaryWrite(message.avatarURL, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.UpdateProfileReq
 */
export const UpdateProfileReq = new UpdateProfileReq$Type();
