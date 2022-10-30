/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "smplchain.roles";

export interface MsgAdd {
  creator: string;
  rolename: string;
  address: string;
}

export interface MsgAddResponse {
}

export interface MsgRemove {
  creator: string;
  rolename: string;
  address: string;
}

export interface MsgRemoveResponse {
}

function createBaseMsgAdd(): MsgAdd {
  return { creator: "", rolename: "", address: "" };
}

export const MsgAdd = {
  encode(message: MsgAdd, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.rolename !== "") {
      writer.uint32(18).string(message.rolename);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAdd {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAdd();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.rolename = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAdd {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      rolename: isSet(object.rolename) ? String(object.rolename) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgAdd): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.rolename !== undefined && (obj.rolename = message.rolename);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAdd>, I>>(object: I): MsgAdd {
    const message = createBaseMsgAdd();
    message.creator = object.creator ?? "";
    message.rolename = object.rolename ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgAddResponse(): MsgAddResponse {
  return {};
}

export const MsgAddResponse = {
  encode(_: MsgAddResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddResponse {
    return {};
  },

  toJSON(_: MsgAddResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddResponse>, I>>(_: I): MsgAddResponse {
    const message = createBaseMsgAddResponse();
    return message;
  },
};

function createBaseMsgRemove(): MsgRemove {
  return { creator: "", rolename: "", address: "" };
}

export const MsgRemove = {
  encode(message: MsgRemove, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.rolename !== "") {
      writer.uint32(18).string(message.rolename);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemove {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemove();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.rolename = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemove {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      rolename: isSet(object.rolename) ? String(object.rolename) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgRemove): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.rolename !== undefined && (obj.rolename = message.rolename);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemove>, I>>(object: I): MsgRemove {
    const message = createBaseMsgRemove();
    message.creator = object.creator ?? "";
    message.rolename = object.rolename ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgRemoveResponse(): MsgRemoveResponse {
  return {};
}

export const MsgRemoveResponse = {
  encode(_: MsgRemoveResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRemoveResponse {
    return {};
  },

  toJSON(_: MsgRemoveResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveResponse>, I>>(_: I): MsgRemoveResponse {
    const message = createBaseMsgRemoveResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Add(request: MsgAdd): Promise<MsgAddResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Remove(request: MsgRemove): Promise<MsgRemoveResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Add = this.Add.bind(this);
    this.Remove = this.Remove.bind(this);
  }
  Add(request: MsgAdd): Promise<MsgAddResponse> {
    const data = MsgAdd.encode(request).finish();
    const promise = this.rpc.request("smplchain.roles.Msg", "Add", data);
    return promise.then((data) => MsgAddResponse.decode(new _m0.Reader(data)));
  }

  Remove(request: MsgRemove): Promise<MsgRemoveResponse> {
    const data = MsgRemove.encode(request).finish();
    const promise = this.rpc.request("smplchain.roles.Msg", "Remove", data);
    return promise.then((data) => MsgRemoveResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
