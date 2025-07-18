// @generated by protobuf-ts 2.11.1 with parameter generate_dependencies,optimize_speed,long_type_string
// @generated from protobuf file "internal/user/entity.proto" (package "user", syntax proto3)
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
/**
 * @generated from protobuf message user.Entity
 */
export interface Entity {
    /**
     * @generated from protobuf field: bytes ID = 1
     */
    iD: Uint8Array;
    /**
     * @generated from protobuf field: string Name = 2
     */
    name: string;
    /**
     * @generated from protobuf field: string AvatarURL = 3
     */
    avatarURL: string;
    /**
     * @generated from protobuf field: string Description = 4
     */
    description: string;
    /**
     * @generated from protobuf field: int64 CreatedAt = 5
     */
    createdAt: string;
    /**
     * @generated from protobuf field: int64 UpdatedAt = 6
     */
    updatedAt: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Entity$Type extends MessageType<Entity> {
    constructor() {
        super("user.Entity", [
            { no: 1, name: "ID", kind: "scalar", jsonName: "ID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 2, name: "Name", kind: "scalar", jsonName: "Name", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "AvatarURL", kind: "scalar", jsonName: "AvatarURL", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "Description", kind: "scalar", jsonName: "Description", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "CreatedAt", kind: "scalar", jsonName: "CreatedAt", T: 3 /*ScalarType.INT64*/ },
            { no: 6, name: "UpdatedAt", kind: "scalar", jsonName: "UpdatedAt", T: 3 /*ScalarType.INT64*/ }
        ]);
    }
    create(value?: PartialMessage<Entity>): Entity {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.iD = new Uint8Array(0);
        message.name = "";
        message.avatarURL = "";
        message.description = "";
        message.createdAt = "0";
        message.updatedAt = "0";
        if (value !== undefined)
            reflectionMergePartial<Entity>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Entity): Entity {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes ID */ 1:
                    message.iD = reader.bytes();
                    break;
                case /* string Name */ 2:
                    message.name = reader.string();
                    break;
                case /* string AvatarURL */ 3:
                    message.avatarURL = reader.string();
                    break;
                case /* string Description */ 4:
                    message.description = reader.string();
                    break;
                case /* int64 CreatedAt */ 5:
                    message.createdAt = reader.int64().toString();
                    break;
                case /* int64 UpdatedAt */ 6:
                    message.updatedAt = reader.int64().toString();
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
    internalBinaryWrite(message: Entity, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes ID = 1; */
        if (message.iD.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.iD);
        /* string Name = 2; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* string AvatarURL = 3; */
        if (message.avatarURL !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.avatarURL);
        /* string Description = 4; */
        if (message.description !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.description);
        /* int64 CreatedAt = 5; */
        if (message.createdAt !== "0")
            writer.tag(5, WireType.Varint).int64(message.createdAt);
        /* int64 UpdatedAt = 6; */
        if (message.updatedAt !== "0")
            writer.tag(6, WireType.Varint).int64(message.updatedAt);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message user.Entity
 */
export const Entity = new Entity$Type();
