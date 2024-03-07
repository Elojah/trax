// @generated by protobuf-ts 2.9.3 with parameter optimize_speed
// @generated from protobuf file "github.com/elojah/trax/internal/user/dto/user.proto" (package "dto", syntax proto3)
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
 * @generated from protobuf message dto.SigninResp
 */
export interface SigninResp {
    /**
     * @generated from protobuf field: string AccessToken = 1 [json_name = "AccessToken"];
     */
    accessToken: string;
    /**
     * @generated from protobuf field: string RefreshToken = 2 [json_name = "RefreshToken"];
     */
    refreshToken: string;
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
                case /* string AccessToken = 1 [json_name = "AccessToken"];*/ 1:
                    message.accessToken = reader.string();
                    break;
                case /* string RefreshToken = 2 [json_name = "RefreshToken"];*/ 2:
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
        /* string AccessToken = 1 [json_name = "AccessToken"]; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* string RefreshToken = 2 [json_name = "RefreshToken"]; */
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