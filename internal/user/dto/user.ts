// @generated by protobuf-ts 2.11.1 with parameter generate_dependencies,optimize_speed,long_type_string
// @generated from protobuf file "internal/user/dto/user.proto" (package "dto", syntax proto3)
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
import { U } from "../user";
import { Paginate } from "../../../pkg/paginate/paginate";
import { String$ } from "../../../pkg/pbtypes/string";
/**
 * @generated from protobuf message dto.SigninResp
 */
export interface SigninResp {
    /**
     * @generated from protobuf field: string AccessToken = 1
     */
    accessToken: string;
    /**
     * @generated from protobuf field: string RefreshToken = 2
     */
    refreshToken: string;
}
/**
 * @generated from protobuf message dto.SignupReq
 */
export interface SignupReq {
    /**
     * @generated from protobuf field: string Firstname = 1
     */
    firstname: string;
    /**
     * @generated from protobuf field: string Lastname = 2
     */
    lastname: string;
    /**
     * @generated from protobuf field: string Email = 3
     */
    email: string;
    /**
     * @generated from protobuf field: string Password = 4
     */
    password: string;
}
/**
 * @generated from protobuf message dto.SigninReq
 */
export interface SigninReq {
    /**
     * @generated from protobuf field: string Email = 1
     */
    email: string;
    /**
     * @generated from protobuf field: string Password = 2
     */
    password: string;
}
/**
 * @generated from protobuf message dto.FetchUserReq
 */
export interface FetchUserReq {
    /**
     * @generated from protobuf field: bool Me = 1
     */
    me: boolean;
    /**
     * @generated from protobuf field: bytes ID = 2
     */
    iD: Uint8Array;
    /**
     * @generated from protobuf field: bytes EntityID = 3
     */
    entityID: Uint8Array;
}
/**
 * @generated from protobuf message dto.UpdateUserReq
 */
export interface UpdateUserReq {
    /**
     * @generated from protobuf field: pbtypes.String Firstname = 1
     */
    firstname?: String$;
    /**
     * @generated from protobuf field: pbtypes.String Lastname = 2
     */
    lastname?: String$;
    /**
     * @generated from protobuf field: pbtypes.String AvatarURL = 3
     */
    avatarURL?: String$;
}
/**
 * @generated from protobuf message dto.ListUserReq
 */
export interface ListUserReq {
    /**
     * @generated from protobuf field: paginate.Paginate Paginate = 1
     */
    paginate?: Paginate;
    /**
     * @generated from protobuf field: string Search = 2
     */
    search: string;
    /**
     * @generated from protobuf field: bool OwnEntity = 3
     */
    ownEntity: boolean;
    /**
     * @generated from protobuf field: repeated bytes IDs = 4
     */
    iDs: Uint8Array[];
    /**
     * @generated from protobuf field: repeated bytes EntityIDs = 5
     */
    entityIDs: Uint8Array[];
    /**
     * @generated from protobuf field: bytes RoleID = 6
     */
    roleID: Uint8Array;
}
/**
 * @generated from protobuf message dto.ListUserResp
 */
export interface ListUserResp {
    /**
     * @generated from protobuf field: repeated user.U Users = 1
     */
    users: U[];
    /**
     * @generated from protobuf field: uint64 Total = 2
     */
    total: string;
}
/**
 * @generated from protobuf message dto.InviteUserReq
 */
export interface InviteUserReq {
    /**
     * @generated from protobuf field: string Email = 1
     */
    email: string;
    /**
     * @generated from protobuf field: bytes EntityID = 2
     */
    entityID: Uint8Array;
    /**
     * @generated from protobuf field: repeated bytes RoleIDs = 3
     */
    roleIDs: Uint8Array[];
}
// @generated message type with reflection information, may provide speed optimized methods
class SigninResp$Type extends MessageType<SigninResp> {
    constructor() {
        super("dto.SigninResp", [
            { no: 1, name: "AccessToken", kind: "scalar", jsonName: "AccessToken", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "RefreshToken", kind: "scalar", jsonName: "RefreshToken", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SigninResp>): SigninResp {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.accessToken = "";
        message.refreshToken = "";
        if (value !== undefined)
            reflectionMergePartial<SigninResp>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SigninResp): SigninResp {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string AccessToken */ 1:
                    message.accessToken = reader.string();
                    break;
                case /* string RefreshToken */ 2:
                    message.refreshToken = reader.string();
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
    internalBinaryWrite(message: SigninResp, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string AccessToken = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* string RefreshToken = 2; */
        if (message.refreshToken !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.refreshToken);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.SigninResp
 */
export const SigninResp = new SigninResp$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SignupReq$Type extends MessageType<SignupReq> {
    constructor() {
        super("dto.SignupReq", [
            { no: 1, name: "Firstname", kind: "scalar", jsonName: "Firstname", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "Lastname", kind: "scalar", jsonName: "Lastname", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "Email", kind: "scalar", jsonName: "Email", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "Password", kind: "scalar", jsonName: "Password", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SignupReq>): SignupReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.firstname = "";
        message.lastname = "";
        message.email = "";
        message.password = "";
        if (value !== undefined)
            reflectionMergePartial<SignupReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SignupReq): SignupReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string Firstname */ 1:
                    message.firstname = reader.string();
                    break;
                case /* string Lastname */ 2:
                    message.lastname = reader.string();
                    break;
                case /* string Email */ 3:
                    message.email = reader.string();
                    break;
                case /* string Password */ 4:
                    message.password = reader.string();
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
    internalBinaryWrite(message: SignupReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string Firstname = 1; */
        if (message.firstname !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.firstname);
        /* string Lastname = 2; */
        if (message.lastname !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.lastname);
        /* string Email = 3; */
        if (message.email !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.email);
        /* string Password = 4; */
        if (message.password !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.password);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.SignupReq
 */
export const SignupReq = new SignupReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SigninReq$Type extends MessageType<SigninReq> {
    constructor() {
        super("dto.SigninReq", [
            { no: 1, name: "Email", kind: "scalar", jsonName: "Email", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "Password", kind: "scalar", jsonName: "Password", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SigninReq>): SigninReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.email = "";
        message.password = "";
        if (value !== undefined)
            reflectionMergePartial<SigninReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SigninReq): SigninReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string Email */ 1:
                    message.email = reader.string();
                    break;
                case /* string Password */ 2:
                    message.password = reader.string();
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
    internalBinaryWrite(message: SigninReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string Email = 1; */
        if (message.email !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.email);
        /* string Password = 2; */
        if (message.password !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.password);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.SigninReq
 */
export const SigninReq = new SigninReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class FetchUserReq$Type extends MessageType<FetchUserReq> {
    constructor() {
        super("dto.FetchUserReq", [
            { no: 1, name: "Me", kind: "scalar", jsonName: "Me", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "ID", kind: "scalar", jsonName: "ID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 3, name: "EntityID", kind: "scalar", jsonName: "EntityID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<FetchUserReq>): FetchUserReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.me = false;
        message.iD = new Uint8Array(0);
        message.entityID = new Uint8Array(0);
        if (value !== undefined)
            reflectionMergePartial<FetchUserReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: FetchUserReq): FetchUserReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool Me */ 1:
                    message.me = reader.bool();
                    break;
                case /* bytes ID */ 2:
                    message.iD = reader.bytes();
                    break;
                case /* bytes EntityID */ 3:
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
    internalBinaryWrite(message: FetchUserReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool Me = 1; */
        if (message.me !== false)
            writer.tag(1, WireType.Varint).bool(message.me);
        /* bytes ID = 2; */
        if (message.iD.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.iD);
        /* bytes EntityID = 3; */
        if (message.entityID.length)
            writer.tag(3, WireType.LengthDelimited).bytes(message.entityID);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.FetchUserReq
 */
export const FetchUserReq = new FetchUserReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UpdateUserReq$Type extends MessageType<UpdateUserReq> {
    constructor() {
        super("dto.UpdateUserReq", [
            { no: 1, name: "Firstname", kind: "message", jsonName: "Firstname", T: () => String$ },
            { no: 2, name: "Lastname", kind: "message", jsonName: "Lastname", T: () => String$ },
            { no: 3, name: "AvatarURL", kind: "message", jsonName: "AvatarURL", T: () => String$ }
        ]);
    }
    create(value?: PartialMessage<UpdateUserReq>): UpdateUserReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        if (value !== undefined)
            reflectionMergePartial<UpdateUserReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UpdateUserReq): UpdateUserReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* pbtypes.String Firstname */ 1:
                    message.firstname = String$.internalBinaryRead(reader, reader.uint32(), options, message.firstname);
                    break;
                case /* pbtypes.String Lastname */ 2:
                    message.lastname = String$.internalBinaryRead(reader, reader.uint32(), options, message.lastname);
                    break;
                case /* pbtypes.String AvatarURL */ 3:
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
    internalBinaryWrite(message: UpdateUserReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* pbtypes.String Firstname = 1; */
        if (message.firstname)
            String$.internalBinaryWrite(message.firstname, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* pbtypes.String Lastname = 2; */
        if (message.lastname)
            String$.internalBinaryWrite(message.lastname, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* pbtypes.String AvatarURL = 3; */
        if (message.avatarURL)
            String$.internalBinaryWrite(message.avatarURL, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.UpdateUserReq
 */
export const UpdateUserReq = new UpdateUserReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListUserReq$Type extends MessageType<ListUserReq> {
    constructor() {
        super("dto.ListUserReq", [
            { no: 1, name: "Paginate", kind: "message", jsonName: "Paginate", T: () => Paginate, options: { "gogoproto.nullable": true } },
            { no: 2, name: "Search", kind: "scalar", jsonName: "Search", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "OwnEntity", kind: "scalar", jsonName: "OwnEntity", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "IDs", kind: "scalar", jsonName: "IDs", repeat: 2 /*RepeatType.UNPACKED*/, T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 5, name: "EntityIDs", kind: "scalar", jsonName: "EntityIDs", repeat: 2 /*RepeatType.UNPACKED*/, T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 6, name: "RoleID", kind: "scalar", jsonName: "RoleID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<ListUserReq>): ListUserReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.search = "";
        message.ownEntity = false;
        message.iDs = [];
        message.entityIDs = [];
        message.roleID = new Uint8Array(0);
        if (value !== undefined)
            reflectionMergePartial<ListUserReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListUserReq): ListUserReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* paginate.Paginate Paginate */ 1:
                    message.paginate = Paginate.internalBinaryRead(reader, reader.uint32(), options, message.paginate);
                    break;
                case /* string Search */ 2:
                    message.search = reader.string();
                    break;
                case /* bool OwnEntity */ 3:
                    message.ownEntity = reader.bool();
                    break;
                case /* repeated bytes IDs */ 4:
                    message.iDs.push(reader.bytes());
                    break;
                case /* repeated bytes EntityIDs */ 5:
                    message.entityIDs.push(reader.bytes());
                    break;
                case /* bytes RoleID */ 6:
                    message.roleID = reader.bytes();
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
    internalBinaryWrite(message: ListUserReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* paginate.Paginate Paginate = 1; */
        if (message.paginate)
            Paginate.internalBinaryWrite(message.paginate, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* string Search = 2; */
        if (message.search !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.search);
        /* bool OwnEntity = 3; */
        if (message.ownEntity !== false)
            writer.tag(3, WireType.Varint).bool(message.ownEntity);
        /* repeated bytes IDs = 4; */
        for (let i = 0; i < message.iDs.length; i++)
            writer.tag(4, WireType.LengthDelimited).bytes(message.iDs[i]);
        /* repeated bytes EntityIDs = 5; */
        for (let i = 0; i < message.entityIDs.length; i++)
            writer.tag(5, WireType.LengthDelimited).bytes(message.entityIDs[i]);
        /* bytes RoleID = 6; */
        if (message.roleID.length)
            writer.tag(6, WireType.LengthDelimited).bytes(message.roleID);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.ListUserReq
 */
export const ListUserReq = new ListUserReq$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListUserResp$Type extends MessageType<ListUserResp> {
    constructor() {
        super("dto.ListUserResp", [
            { no: 1, name: "Users", kind: "message", jsonName: "Users", repeat: 2 /*RepeatType.UNPACKED*/, T: () => U, options: { "gogoproto.nullable": false } },
            { no: 2, name: "Total", kind: "scalar", jsonName: "Total", T: 4 /*ScalarType.UINT64*/ }
        ]);
    }
    create(value?: PartialMessage<ListUserResp>): ListUserResp {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.users = [];
        message.total = "0";
        if (value !== undefined)
            reflectionMergePartial<ListUserResp>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListUserResp): ListUserResp {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated user.U Users */ 1:
                    message.users.push(U.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* uint64 Total */ 2:
                    message.total = reader.uint64().toString();
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
    internalBinaryWrite(message: ListUserResp, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated user.U Users = 1; */
        for (let i = 0; i < message.users.length; i++)
            U.internalBinaryWrite(message.users[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* uint64 Total = 2; */
        if (message.total !== "0")
            writer.tag(2, WireType.Varint).uint64(message.total);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.ListUserResp
 */
export const ListUserResp = new ListUserResp$Type();
// @generated message type with reflection information, may provide speed optimized methods
class InviteUserReq$Type extends MessageType<InviteUserReq> {
    constructor() {
        super("dto.InviteUserReq", [
            { no: 1, name: "Email", kind: "scalar", jsonName: "Email", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "EntityID", kind: "scalar", jsonName: "EntityID", T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } },
            { no: 3, name: "RoleIDs", kind: "scalar", jsonName: "RoleIDs", repeat: 2 /*RepeatType.UNPACKED*/, T: 12 /*ScalarType.BYTES*/, options: { "gogoproto.nullable": false, "gogoproto.customtype": "github.com/elojah/trax/pkg/ulid.ID" } }
        ]);
    }
    create(value?: PartialMessage<InviteUserReq>): InviteUserReq {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.email = "";
        message.entityID = new Uint8Array(0);
        message.roleIDs = [];
        if (value !== undefined)
            reflectionMergePartial<InviteUserReq>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: InviteUserReq): InviteUserReq {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string Email */ 1:
                    message.email = reader.string();
                    break;
                case /* bytes EntityID */ 2:
                    message.entityID = reader.bytes();
                    break;
                case /* repeated bytes RoleIDs */ 3:
                    message.roleIDs.push(reader.bytes());
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
    internalBinaryWrite(message: InviteUserReq, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string Email = 1; */
        if (message.email !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.email);
        /* bytes EntityID = 2; */
        if (message.entityID.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.entityID);
        /* repeated bytes RoleIDs = 3; */
        for (let i = 0; i < message.roleIDs.length; i++)
            writer.tag(3, WireType.LengthDelimited).bytes(message.roleIDs[i]);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message dto.InviteUserReq
 */
export const InviteUserReq = new InviteUserReq$Type();
