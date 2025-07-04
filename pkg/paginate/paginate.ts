// @generated by protobuf-ts 2.11.1 with parameter generate_dependencies,optimize_speed,long_type_string
// @generated from protobuf file "pkg/paginate/paginate.proto" (package "paginate", syntax proto3)
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
 * @generated from protobuf message paginate.Paginate
 */
export interface Paginate {
    /**
     * @generated from protobuf field: int64 Start = 1
     */
    start: string;
    /**
     * @generated from protobuf field: int64 End = 2
     */
    end: string;
    /**
     * @generated from protobuf field: string Sort = 3
     */
    sort: string;
    /**
     * @generated from protobuf field: bool Order = 4
     */
    order: boolean;
}
// @generated message type with reflection information, may provide speed optimized methods
class Paginate$Type extends MessageType<Paginate> {
    constructor() {
        super("paginate.Paginate", [
            { no: 1, name: "Start", kind: "scalar", jsonName: "Start", T: 3 /*ScalarType.INT64*/ },
            { no: 2, name: "End", kind: "scalar", jsonName: "End", T: 3 /*ScalarType.INT64*/ },
            { no: 3, name: "Sort", kind: "scalar", jsonName: "Sort", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "Order", kind: "scalar", jsonName: "Order", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value?: PartialMessage<Paginate>): Paginate {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.start = "0";
        message.end = "0";
        message.sort = "";
        message.order = false;
        if (value !== undefined)
            reflectionMergePartial<Paginate>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Paginate): Paginate {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int64 Start */ 1:
                    message.start = reader.int64().toString();
                    break;
                case /* int64 End */ 2:
                    message.end = reader.int64().toString();
                    break;
                case /* string Sort */ 3:
                    message.sort = reader.string();
                    break;
                case /* bool Order */ 4:
                    message.order = reader.bool();
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
    internalBinaryWrite(message: Paginate, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int64 Start = 1; */
        if (message.start !== "0")
            writer.tag(1, WireType.Varint).int64(message.start);
        /* int64 End = 2; */
        if (message.end !== "0")
            writer.tag(2, WireType.Varint).int64(message.end);
        /* string Sort = 3; */
        if (message.sort !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.sort);
        /* bool Order = 4; */
        if (message.order !== false)
            writer.tag(4, WireType.Varint).bool(message.order);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message paginate.Paginate
 */
export const Paginate = new Paginate$Type();
