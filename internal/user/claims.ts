// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/internal/user/claims.proto" (package "user", syntax proto3)
// tslint:disable
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Command } from "./role";
/**
 * @generated from protobuf message user.ClaimCommands
 */
export interface ClaimCommands {
    /**
     * @generated from protobuf field: repeated user.Command Commands = 1 [json_name = "Commands"];
     */
    commands: Command[];
}
/**
 * @generated from protobuf message user.ClaimEntity
 */
export interface ClaimEntity {
    /**
     * @generated from protobuf field: bytes ID = 1 [json_name = "ID"];
     */
    iD: Uint8Array;
    /**
     * Permissions key must be a user.Resource
     *
     * @generated from protobuf field: map<string, user.ClaimCommands> Commands = 2 [json_name = "Commands"];
     */
    commands: {
        [key: string]: ClaimCommands;
    };
}
/**
 * @generated from protobuf message user.ClaimAuth
 */
export interface ClaimAuth {
    /**
     * Permissions key must be a entity.ID
     *
     * @generated from protobuf field: map<string, user.ClaimEntity> Entities = 1 [json_name = "Entities"];
     */
    entities: {
        [key: string]: ClaimEntity;
    };
}
// @generated message type with reflection information, may provide speed optimized methods
class ClaimCommands$Type extends MessageType<ClaimCommands> {
    constructor() {
        super("user.ClaimCommands", [
            { no: 1, name: "Commands", kind: "enum", jsonName: "Commands", repeat: 1 /*RepeatType.PACKED*/, T: () => ["user.Command", Command] }
        ]);
    }
    create(value?: PartialMessage<ClaimCommands>): ClaimCommands {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.commands = [];
        if (value !== undefined)
            reflectionMergePartial<ClaimCommands>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ClaimCommands): ClaimCommands {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated user.Command Commands = 1 [json_name = "Commands"];*/ 1:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.commands.push(reader.int32());
                    else
                        message.commands.push(reader.int32());
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
    internalBinaryWrite(message: ClaimCommands, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated user.Command Commands = 1 [json_name = "Commands"]; */
        if (message.commands.length) {
            writer.tag(1, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.commands.length; i++)
                writer.int32(message.commands[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message user.ClaimCommands
 */
export const ClaimCommands = new ClaimCommands$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ClaimEntity$Type extends MessageType<ClaimEntity> {
    constructor() {
        super("user.ClaimEntity", [
            { no: 1, name: "ID", kind: "scalar", jsonName: "ID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 2, name: "Commands", kind: "map", jsonName: "Commands", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => ClaimCommands }, options: { "gogoproto.nullable": false } }
        ]);
    }
    create(value?: PartialMessage<ClaimEntity>): ClaimEntity {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.iD = new Uint8Array(0);
        message.commands = {};
        if (value !== undefined)
            reflectionMergePartial<ClaimEntity>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ClaimEntity): ClaimEntity {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes ID = 1 [json_name = "ID"];*/ 1:
                    message.iD = reader.bytes();
                    break;
                case /* map<string, user.ClaimCommands> Commands = 2 [json_name = "Commands"];*/ 2:
                    this.binaryReadMap2(message.commands, reader, options);
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
    private binaryReadMap2(map: ClaimEntity["commands"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof ClaimEntity["commands"] | undefined, val: ClaimEntity["commands"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = ClaimCommands.internalBinaryRead(reader, reader.uint32(), options);
                    break;
                default: throw new globalThis.Error("unknown map entry field for field user.ClaimEntity.Commands");
            }
        }
        map[key ?? ""] = val ?? ClaimCommands.create();
    }
    internalBinaryWrite(message: ClaimEntity, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes ID = 1 [json_name = "ID"]; */
        if (message.iD.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.iD);
        /* map<string, user.ClaimCommands> Commands = 2 [json_name = "Commands"]; */
        for (let k of globalThis.Object.keys(message.commands)) {
            writer.tag(2, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k);
            writer.tag(2, WireType.LengthDelimited).fork();
            ClaimCommands.internalBinaryWrite(message.commands[k], writer, options);
            writer.join().join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message user.ClaimEntity
 */
export const ClaimEntity = new ClaimEntity$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ClaimAuth$Type extends MessageType<ClaimAuth> {
    constructor() {
        super("user.ClaimAuth", [
            { no: 1, name: "Entities", kind: "map", jsonName: "Entities", K: 9 /*ScalarType.STRING*/, V: { kind: "message", T: () => ClaimEntity }, options: { "gogoproto.nullable": false } }
        ]);
    }
    create(value?: PartialMessage<ClaimAuth>): ClaimAuth {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.entities = {};
        if (value !== undefined)
            reflectionMergePartial<ClaimAuth>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ClaimAuth): ClaimAuth {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* map<string, user.ClaimEntity> Entities = 1 [json_name = "Entities"];*/ 1:
                    this.binaryReadMap1(message.entities, reader, options);
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
    private binaryReadMap1(map: ClaimAuth["entities"], reader: IBinaryReader, options: BinaryReadOptions): void {
        let len = reader.uint32(), end = reader.pos + len, key: keyof ClaimAuth["entities"] | undefined, val: ClaimAuth["entities"][any] | undefined;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case 1:
                    key = reader.string();
                    break;
                case 2:
                    val = ClaimEntity.internalBinaryRead(reader, reader.uint32(), options);
                    break;
                default: throw new globalThis.Error("unknown map entry field for field user.ClaimAuth.Entities");
            }
        }
        map[key ?? ""] = val ?? ClaimEntity.create();
    }
    internalBinaryWrite(message: ClaimAuth, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* map<string, user.ClaimEntity> Entities = 1 [json_name = "Entities"]; */
        for (let k of globalThis.Object.keys(message.entities)) {
            writer.tag(1, WireType.LengthDelimited).fork().tag(1, WireType.LengthDelimited).string(k);
            writer.tag(2, WireType.LengthDelimited).fork();
            ClaimEntity.internalBinaryWrite(message.entities[k], writer, options);
            writer.join().join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message user.ClaimAuth
 */
export const ClaimAuth = new ClaimAuth$Type();