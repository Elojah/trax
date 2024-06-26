// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/internal/user/dto/entity.proto" (package "dto", syntax proto3)
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
import { Entity } from "../entity";
import { Paginate } from "../../../pkg/paginate/paginate";
import { String$ } from "../../../pkg/pbtypes/string";
/**
 * @generated from protobuf message dto.CreateEntityReq
 */
export interface CreateEntityReq {
    /**
     * @generated from protobuf field: string Name = 1 [json_name = "Name"];
     */
    name: string;
    /**
     * @generated from protobuf field: string Description = 2 [json_name = "Description"];
     */
    description: string;
    /**
     * @generated from protobuf field: string AvatarURL = 3 [json_name = "AvatarURL"];
     */
    avatarURL: string;
}
/**
 * @generated from protobuf message dto.UpdateEntityReq
 */
export interface UpdateEntityReq {
    /**
     * @generated from protobuf field: bytes ID = 1 [json_name = "ID"];
     */
    iD: Uint8Array;
    /**
     * @generated from protobuf field: pbtypes.String Name = 2 [json_name = "Name"];
     */
    name?: String$;
    /**
     * @generated from protobuf field: pbtypes.String Description = 3 [json_name = "Description"];
     */
    description?: String$;
    /**
     * @generated from protobuf field: pbtypes.String AvatarURL = 4 [json_name = "AvatarURL"];
     */
    avatarURL?: String$;
}
/**
 * @generated from protobuf message dto.FetchEntityReq
 */
export interface FetchEntityReq {
    /**
     * @generated from protobuf field: bytes ID = 1 [json_name = "ID"];
     */
    iD: Uint8Array;
}
/**
 * @generated from protobuf message dto.ListEntityReq
 */
export interface ListEntityReq {
    /**
     * @generated from protobuf field: paginate.Paginate Paginate = 1 [json_name = "Paginate"];
     */
    paginate?: Paginate;
    /**
     * @generated from protobuf field: string Search = 2 [json_name = "Search"];
     */
    search: string;
    /**
     * @generated from protobuf field: bool Own = 3 [json_name = "Own"];
     */
    own: boolean;
    /**
     * @generated from protobuf field: repeated bytes IDs = 4 [json_name = "IDs"];
     */
    iDs: Uint8Array[];
}
/**
 * @generated from protobuf message dto.ListEntityResp
 */
export interface ListEntityResp {
    /**
     * @generated from protobuf field: repeated user.Entity Entities = 1 [json_name = "Entities"];
     */
    entities: Entity[];
    /**
     * @generated from protobuf field: uint64 Total = 2 [json_name = "Total"];
     */
    total: bigint;
}
/**
 * @generated from protobuf message dto.DeleteEntityReq
 */
export interface DeleteEntityReq {
    /**
     * @generated from protobuf field: bytes ID = 1 [json_name = "ID"];
     */
    iD: Uint8Array;
}
/**
 * @generated from protobuf message dto.DeleteEntityResp
 */
export interface DeleteEntityResp {
    /**
     * @generated from protobuf field: user.Entity Entity = 1 [json_name = "Entity"];
     */
    entity?: Entity;
}
// @generated message type with reflection information, may provide speed optimized methods
class CreateEntityReq$Type extends MessageType<CreateEntityReq> {
    constructor() {
        super("dto.CreateEntityReq", [
            { no: 1, name: "Name", kind: "scalar", jsonName: "Name", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "Description", kind: "scalar", jsonName: "Description", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "AvatarURL", kind: "scalar", jsonName: "AvatarURL", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<CreateEntityReq>): CreateEntityReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.name = "";
        message.description = "";
        message.avatarURL = "";
        if (value !== undefined)
            reflectionMergePartial<CreateEntityReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateEntityReq): CreateEntityReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string Name = 1 [json_name = "Name"];*/ 1:
                    message.name = reader.string();
                    break;
                case /* string Description = 2 [json_name = "Description"];*/ 2:
                    message.description = reader.string();
                    break;
                case /* string AvatarURL = 3 [json_name = "AvatarURL"];*/ 3:
                    message.avatarURL = reader.string();
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
    internalBinaryWrite(message: CreateEntityReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string Name = 1 [json_name = "Name"]; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string Description = 2 [json_name = "Description"]; */
        if (message.description !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.description);
        /* string AvatarURL = 3 [json_name = "AvatarURL"]; */
        if (message.avatarURL !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.avatarURL);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.CreateEntityReq
 */
export const CreateEntityReq = new CreateEntityReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateEntityReq$Type extends MessageType<UpdateEntityReq> {
    constructor() {
        super("dto.UpdateEntityReq", [
            { no: 1, name: "ID", kind: "scalar", jsonName: "ID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 2, name: "Name", kind: "message", jsonName: "Name", T: () => String$ },
            { no: 3, name: "Description", kind: "message", jsonName: "Description", T: () => String$ },
            { no: 4, name: "AvatarURL", kind: "message", jsonName: "AvatarURL", T: () => String$ }
        ]);
    }
    create(value?: PartialMessage<UpdateEntityReq>): UpdateEntityReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.iD = new Uint8Array(0);
        if (value !== undefined)
            reflectionMergePartial<UpdateEntityReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateEntityReq): UpdateEntityReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes ID = 1 [json_name = "ID"];*/ 1:
                    message.iD = reader.bytes();
                    break;
                case /* pbtypes.String Name = 2 [json_name = "Name"];*/ 2:
                    message.name = String$.internalBinaryRead(reader, reader.uint32(), options, message.name);
                    break;
                case /* pbtypes.String Description = 3 [json_name = "Description"];*/ 3:
                    message.description = String$.internalBinaryRead(reader, reader.uint32(), options, message.description);
                    break;
                case /* pbtypes.String AvatarURL = 4 [json_name = "AvatarURL"];*/ 4:
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
    internalBinaryWrite(message: UpdateEntityReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes ID = 1 [json_name = "ID"]; */
        if (message.iD.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.iD);
        /* pbtypes.String Name = 2 [json_name = "Name"]; */
        if (message.name)
            String$.internalBinaryWrite(message.name, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* pbtypes.String Description = 3 [json_name = "Description"]; */
        if (message.description)
            String$.internalBinaryWrite(message.description, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* pbtypes.String AvatarURL = 4 [json_name = "AvatarURL"]; */
        if (message.avatarURL)
            String$.internalBinaryWrite(message.avatarURL, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.UpdateEntityReq
 */
export const UpdateEntityReq = new UpdateEntityReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FetchEntityReq$Type extends MessageType<FetchEntityReq> {
    constructor() {
        super("dto.FetchEntityReq", [
            { no: 1, name: "ID", kind: "scalar", jsonName: "ID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<FetchEntityReq>): FetchEntityReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.iD = new Uint8Array(0);
        if (value !== undefined)
            reflectionMergePartial<FetchEntityReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: FetchEntityReq): FetchEntityReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes ID = 1 [json_name = "ID"];*/ 1:
                    message.iD = reader.bytes();
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
    internalBinaryWrite(message: FetchEntityReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes ID = 1 [json_name = "ID"]; */
        if (message.iD.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.iD);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.FetchEntityReq
 */
export const FetchEntityReq = new FetchEntityReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListEntityReq$Type extends MessageType<ListEntityReq> {
    constructor() {
        super("dto.ListEntityReq", [
            { no: 1, name: "Paginate", kind: "message", jsonName: "Paginate", T: () => Paginate, options: { "gogoproto.nullable": true } },
            { no: 2, name: "Search", kind: "scalar", jsonName: "Search", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "Own", kind: "scalar", jsonName: "Own", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "IDs", kind: "scalar", jsonName: "IDs", repeat: 2 /*RepeatType.UNPACKED*/, T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<ListEntityReq>): ListEntityReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.search = "";
        message.own = false;
        message.iDs = [];
        if (value !== undefined)
            reflectionMergePartial<ListEntityReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListEntityReq): ListEntityReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* paginate.Paginate Paginate = 1 [json_name = "Paginate"];*/ 1:
                    message.paginate = Paginate.internalBinaryRead(reader, reader.uint32(), options, message.paginate);
                    break;
                case /* string Search = 2 [json_name = "Search"];*/ 2:
                    message.search = reader.string();
                    break;
                case /* bool Own = 3 [json_name = "Own"];*/ 3:
                    message.own = reader.bool();
                    break;
                case /* repeated bytes IDs = 4 [json_name = "IDs"];*/ 4:
                    message.iDs.push(reader.bytes());
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
    internalBinaryWrite(message: ListEntityReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* paginate.Paginate Paginate = 1 [json_name = "Paginate"]; */
        if (message.paginate)
            Paginate.internalBinaryWrite(message.paginate, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string Search = 2 [json_name = "Search"]; */
        if (message.search !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.search);
        /* bool Own = 3 [json_name = "Own"]; */
        if (message.own !== false)
            writer.tag(3, WireType.Varint).bool(message.own);
        /* repeated bytes IDs = 4 [json_name = "IDs"]; */
        for (let i = 0; i < message.iDs.length; i++)
            writer.tag(4, WireType.LengthDelimited).bytes(message.iDs[i]);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.ListEntityReq
 */
export const ListEntityReq = new ListEntityReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListEntityResp$Type extends MessageType<ListEntityResp> {
    constructor() {
        super("dto.ListEntityResp", [
            { no: 1, name: "Entities", kind: "message", jsonName: "Entities", repeat: 1 /*RepeatType.PACKED*/, T: () => Entity, options: { "gogoproto.nullable": false } },
            { no: 2, name: "Total", kind: "scalar", jsonName: "Total", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<ListEntityResp>): ListEntityResp {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.entities = [];
        message.total = 0n;
        if (value !== undefined)
            reflectionMergePartial<ListEntityResp>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListEntityResp): ListEntityResp {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated user.Entity Entities = 1 [json_name = "Entities"];*/ 1:
                    message.entities.push(Entity.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* uint64 Total = 2 [json_name = "Total"];*/ 2:
                    message.total = reader.uint64().toBigInt();
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
    internalBinaryWrite(message: ListEntityResp, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated user.Entity Entities = 1 [json_name = "Entities"]; */
        for (let i = 0; i < message.entities.length; i++)
            Entity.internalBinaryWrite(message.entities[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* uint64 Total = 2 [json_name = "Total"]; */
        if (message.total !== 0n)
            writer.tag(2, WireType.Varint).uint64(message.total);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.ListEntityResp
 */
export const ListEntityResp = new ListEntityResp$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteEntityReq$Type extends MessageType<DeleteEntityReq> {
    constructor() {
        super("dto.DeleteEntityReq", [
            { no: 1, name: "ID", kind: "scalar", jsonName: "ID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<DeleteEntityReq>): DeleteEntityReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.iD = new Uint8Array(0);
        if (value !== undefined)
            reflectionMergePartial<DeleteEntityReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteEntityReq): DeleteEntityReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes ID = 1 [json_name = "ID"];*/ 1:
                    message.iD = reader.bytes();
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
    internalBinaryWrite(message: DeleteEntityReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes ID = 1 [json_name = "ID"]; */
        if (message.iD.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.iD);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.DeleteEntityReq
 */
export const DeleteEntityReq = new DeleteEntityReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteEntityResp$Type extends MessageType<DeleteEntityResp> {
    constructor() {
        super("dto.DeleteEntityResp", [
            { no: 1, name: "Entity", kind: "message", jsonName: "Entity", T: () => Entity, options: { "gogoproto.nullable": false } }
        ]);
    }
    create(value?: PartialMessage<DeleteEntityResp>): DeleteEntityResp {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<DeleteEntityResp>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteEntityResp): DeleteEntityResp {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* user.Entity Entity = 1 [json_name = "Entity"];*/ 1:
                    message.entity = Entity.internalBinaryRead(reader, reader.uint32(), options, message.entity);
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
    internalBinaryWrite(message: DeleteEntityResp, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* user.Entity Entity = 1 [json_name = "Entity"]; */
        if (message.entity)
            Entity.internalBinaryWrite(message.entity, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.DeleteEntityResp
 */
export const DeleteEntityResp = new DeleteEntityResp$Type();
