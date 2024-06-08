// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/internal/user/dto/role.proto" (package "dto", syntax proto3)
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
import { Paginate } from "../../../pkg/paginate/paginate";
import { String$ } from "../../../pkg/pbtypes/string";
import { Permission } from "../role";
import { Role } from "../role";
/**
 * @generated from protobuf message dto.RolePermission
 */
export interface RolePermission {
    /**
     * @generated from protobuf field: user.Role Role = 1 [json_name = "Role"];
     */
    role?: Role;
    /**
     * @generated from protobuf field: repeated user.Permission Permissions = 2 [json_name = "Permissions"];
     */
    permissions: Permission[];
}
/**
 * @generated from protobuf message dto.CreateRoleReq
 */
export interface CreateRoleReq {
    /**
     * @generated from protobuf field: bytes EntityID = 1 [json_name = "EntityID"];
     */
    entityID: Uint8Array;
    /**
     * @generated from protobuf field: string Name = 2 [json_name = "Name"];
     */
    name: string;
    /**
     * @generated from protobuf field: repeated user.Permission Permissions = 3 [json_name = "Permissions"];
     */
    permissions: Permission[];
}
/**
 * @generated from protobuf message dto.UpdateRoleReq
 */
export interface UpdateRoleReq {
    /**
     * @generated from protobuf field: bytes ID = 1 [json_name = "ID"];
     */
    iD: Uint8Array;
    /**
     * @generated from protobuf field: pbtypes.String Name = 2 [json_name = "Name"];
     */
    name?: String$;
    /**
     * @generated from protobuf field: repeated user.Permission Permissions = 3 [json_name = "Permissions"];
     */
    permissions: Permission[];
}
/**
 * @generated from protobuf message dto.ListRoleReq
 */
export interface ListRoleReq {
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
     * @generated from protobuf field: bool OwnEntity = 4 [json_name = "OwnEntity"];
     */
    ownEntity: boolean;
    /**
     * @generated from protobuf field: repeated bytes IDs = 5 [json_name = "IDs"];
     */
    iDs: Uint8Array[];
    /**
     * @generated from protobuf field: repeated bytes EntityIDs = 6 [json_name = "EntityIDs"];
     */
    entityIDs: Uint8Array[];
}
/**
 * @generated from protobuf message dto.ListRoleResp
 */
export interface ListRoleResp {
    /**
     * @generated from protobuf field: repeated dto.RolePermission Roles = 1 [json_name = "Roles"];
     */
    roles: RolePermission[];
    /**
     * @generated from protobuf field: uint64 Total = 2 [json_name = "Total"];
     */
    total: bigint;
}
/**
 * @generated from protobuf message dto.CreateRoleUserReq
 */
export interface CreateRoleUserReq {
    /**
     * @generated from protobuf field: bytes RoleID = 1 [json_name = "RoleID"];
     */
    roleID: Uint8Array;
    /**
     * @generated from protobuf field: bytes UserID = 2 [json_name = "UserID"];
     */
    userID: Uint8Array;
}
// @generated message type with reflection information, may provide speed optimized methods
class RolePermission$Type extends MessageType<RolePermission> {
    constructor() {
        super("dto.RolePermission", [
            { no: 1, name: "Role", kind: "message", jsonName: "Role", T: () => Role, options: { "gogoproto.nullable": false } },
            { no: 2, name: "Permissions", kind: "message", jsonName: "Permissions", repeat: 1 /*RepeatType.PACKED*/, T: () => Permission, options: { "gogoproto.nullable": false } }
        ]);
    }
    create(value?: PartialMessage<RolePermission>): RolePermission {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.permissions = [];
        if (value !== undefined)
            reflectionMergePartial<RolePermission>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: RolePermission): RolePermission {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* user.Role Role = 1 [json_name = "Role"];*/ 1:
                    message.role = Role.internalBinaryRead(reader, reader.uint32(), options, message.role);
                    break;
                case /* repeated user.Permission Permissions = 2 [json_name = "Permissions"];*/ 2:
                    message.permissions.push(Permission.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: RolePermission, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* user.Role Role = 1 [json_name = "Role"]; */
        if (message.role)
            Role.internalBinaryWrite(message.role, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated user.Permission Permissions = 2 [json_name = "Permissions"]; */
        for (let i = 0; i < message.permissions.length; i++)
            Permission.internalBinaryWrite(message.permissions[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.RolePermission
 */
export const RolePermission = new RolePermission$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateRoleReq$Type extends MessageType<CreateRoleReq> {
    constructor() {
        super("dto.CreateRoleReq", [
            { no: 1, name: "EntityID", kind: "scalar", jsonName: "EntityID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 2, name: "Name", kind: "scalar", jsonName: "Name", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "Permissions", kind: "message", jsonName: "Permissions", repeat: 1 /*RepeatType.PACKED*/, T: () => Permission, options: { "gogoproto.nullable": false } }
        ]);
    }
    create(value?: PartialMessage<CreateRoleReq>): CreateRoleReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.entityID = new Uint8Array(0);
        message.name = "";
        message.permissions = [];
        if (value !== undefined)
            reflectionMergePartial<CreateRoleReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateRoleReq): CreateRoleReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes EntityID = 1 [json_name = "EntityID"];*/ 1:
                    message.entityID = reader.bytes();
                    break;
                case /* string Name = 2 [json_name = "Name"];*/ 2:
                    message.name = reader.string();
                    break;
                case /* repeated user.Permission Permissions = 3 [json_name = "Permissions"];*/ 3:
                    message.permissions.push(Permission.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: CreateRoleReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes EntityID = 1 [json_name = "EntityID"]; */
        if (message.entityID.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.entityID);
        /* string Name = 2 [json_name = "Name"]; */
        if (message.name !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.name);
        /* repeated user.Permission Permissions = 3 [json_name = "Permissions"]; */
        for (let i = 0; i < message.permissions.length; i++)
            Permission.internalBinaryWrite(message.permissions[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.CreateRoleReq
 */
export const CreateRoleReq = new CreateRoleReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateRoleReq$Type extends MessageType<UpdateRoleReq> {
    constructor() {
        super("dto.UpdateRoleReq", [
            { no: 1, name: "ID", kind: "scalar", jsonName: "ID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 2, name: "Name", kind: "message", jsonName: "Name", T: () => String$ },
            { no: 3, name: "Permissions", kind: "message", jsonName: "Permissions", repeat: 1 /*RepeatType.PACKED*/, T: () => Permission, options: { "gogoproto.nullable": false } }
        ]);
    }
    create(value?: PartialMessage<UpdateRoleReq>): UpdateRoleReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.iD = new Uint8Array(0);
        message.permissions = [];
        if (value !== undefined)
            reflectionMergePartial<UpdateRoleReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateRoleReq): UpdateRoleReq {
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
                case /* repeated user.Permission Permissions = 3 [json_name = "Permissions"];*/ 3:
                    message.permissions.push(Permission.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: UpdateRoleReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes ID = 1 [json_name = "ID"]; */
        if (message.iD.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.iD);
        /* pbtypes.String Name = 2 [json_name = "Name"]; */
        if (message.name)
            String$.internalBinaryWrite(message.name, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* repeated user.Permission Permissions = 3 [json_name = "Permissions"]; */
        for (let i = 0; i < message.permissions.length; i++)
            Permission.internalBinaryWrite(message.permissions[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.UpdateRoleReq
 */
export const UpdateRoleReq = new UpdateRoleReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListRoleReq$Type extends MessageType<ListRoleReq> {
    constructor() {
        super("dto.ListRoleReq", [
            { no: 1, name: "Paginate", kind: "message", jsonName: "Paginate", T: () => Paginate, options: { "gogoproto.nullable": true } },
            { no: 2, name: "Search", kind: "scalar", jsonName: "Search", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "Own", kind: "scalar", jsonName: "Own", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "OwnEntity", kind: "scalar", jsonName: "OwnEntity", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "IDs", kind: "scalar", jsonName: "IDs", repeat: 2 /*RepeatType.UNPACKED*/, T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 6, name: "EntityIDs", kind: "scalar", jsonName: "EntityIDs", repeat: 2 /*RepeatType.UNPACKED*/, T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<ListRoleReq>): ListRoleReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.search = "";
        message.own = false;
        message.ownEntity = false;
        message.iDs = [];
        message.entityIDs = [];
        if (value !== undefined)
            reflectionMergePartial<ListRoleReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListRoleReq): ListRoleReq {
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
                case /* bool OwnEntity = 4 [json_name = "OwnEntity"];*/ 4:
                    message.ownEntity = reader.bool();
                    break;
                case /* repeated bytes IDs = 5 [json_name = "IDs"];*/ 5:
                    message.iDs.push(reader.bytes());
                    break;
                case /* repeated bytes EntityIDs = 6 [json_name = "EntityIDs"];*/ 6:
                    message.entityIDs.push(reader.bytes());
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
    internalBinaryWrite(message: ListRoleReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* paginate.Paginate Paginate = 1 [json_name = "Paginate"]; */
        if (message.paginate)
            Paginate.internalBinaryWrite(message.paginate, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string Search = 2 [json_name = "Search"]; */
        if (message.search !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.search);
        /* bool Own = 3 [json_name = "Own"]; */
        if (message.own !== false)
            writer.tag(3, WireType.Varint).bool(message.own);
        /* bool OwnEntity = 4 [json_name = "OwnEntity"]; */
        if (message.ownEntity !== false)
            writer.tag(4, WireType.Varint).bool(message.ownEntity);
        /* repeated bytes IDs = 5 [json_name = "IDs"]; */
        for (let i = 0; i < message.iDs.length; i++)
            writer.tag(5, WireType.LengthDelimited).bytes(message.iDs[i]);
        /* repeated bytes EntityIDs = 6 [json_name = "EntityIDs"]; */
        for (let i = 0; i < message.entityIDs.length; i++)
            writer.tag(6, WireType.LengthDelimited).bytes(message.entityIDs[i]);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.ListRoleReq
 */
export const ListRoleReq = new ListRoleReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListRoleResp$Type extends MessageType<ListRoleResp> {
    constructor() {
        super("dto.ListRoleResp", [
            { no: 1, name: "Roles", kind: "message", jsonName: "Roles", repeat: 1 /*RepeatType.PACKED*/, T: () => RolePermission, options: { "gogoproto.nullable": false } },
            { no: 2, name: "Total", kind: "scalar", jsonName: "Total", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<ListRoleResp>): ListRoleResp {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.roles = [];
        message.total = 0n;
        if (value !== undefined)
            reflectionMergePartial<ListRoleResp>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListRoleResp): ListRoleResp {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated dto.RolePermission Roles = 1 [json_name = "Roles"];*/ 1:
                    message.roles.push(RolePermission.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: ListRoleResp, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated dto.RolePermission Roles = 1 [json_name = "Roles"]; */
        for (let i = 0; i < message.roles.length; i++)
            RolePermission.internalBinaryWrite(message.roles[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
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
 * @generated MessageType for protobuf message dto.ListRoleResp
 */
export const ListRoleResp = new ListRoleResp$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateRoleUserReq$Type extends MessageType<CreateRoleUserReq> {
    constructor() {
        super("dto.CreateRoleUserReq", [
            { no: 1, name: "RoleID", kind: "scalar", jsonName: "RoleID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 2, name: "UserID", kind: "scalar", jsonName: "UserID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<CreateRoleUserReq>): CreateRoleUserReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.roleID = new Uint8Array(0);
        message.userID = new Uint8Array(0);
        if (value !== undefined)
            reflectionMergePartial<CreateRoleUserReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateRoleUserReq): CreateRoleUserReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes RoleID = 1 [json_name = "RoleID"];*/ 1:
                    message.roleID = reader.bytes();
                    break;
                case /* bytes UserID = 2 [json_name = "UserID"];*/ 2:
                    message.userID = reader.bytes();
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
    internalBinaryWrite(message: CreateRoleUserReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes RoleID = 1 [json_name = "RoleID"]; */
        if (message.roleID.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.roleID);
        /* bytes UserID = 2 [json_name = "UserID"]; */
        if (message.userID.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.userID);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.CreateRoleUserReq
 */
export const CreateRoleUserReq = new CreateRoleUserReq$Type();
