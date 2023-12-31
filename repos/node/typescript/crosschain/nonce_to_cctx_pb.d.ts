// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file crosschain/nonce_to_cctx.proto (package zetachain.zetacore.crosschain, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * store key is tss+chainid+nonce
 *
 * @generated from message zetachain.zetacore.crosschain.NonceToCctx
 */
export declare class NonceToCctx extends Message<NonceToCctx> {
  /**
   * @generated from field: int64 chain_id = 1;
   */
  chainId: bigint;

  /**
   * @generated from field: int64 nonce = 2;
   */
  nonce: bigint;

  /**
   * @generated from field: string cctxIndex = 3;
   */
  cctxIndex: string;

  /**
   * @generated from field: string tss = 4;
   */
  tss: string;

  constructor(data?: PartialMessage<NonceToCctx>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.NonceToCctx";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NonceToCctx;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NonceToCctx;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NonceToCctx;

  static equals(a: NonceToCctx | PlainMessage<NonceToCctx> | undefined, b: NonceToCctx | PlainMessage<NonceToCctx> | undefined): boolean;
}

/**
 * store key is tss+chainid
 *
 * @generated from message zetachain.zetacore.crosschain.PendingNonces
 */
export declare class PendingNonces extends Message<PendingNonces> {
  /**
   * @generated from field: int64 nonce_low = 1;
   */
  nonceLow: bigint;

  /**
   * @generated from field: int64 nonce_high = 2;
   */
  nonceHigh: bigint;

  /**
   * @generated from field: int64 chain_id = 3;
   */
  chainId: bigint;

  /**
   * @generated from field: string tss = 4;
   */
  tss: string;

  constructor(data?: PartialMessage<PendingNonces>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.PendingNonces";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PendingNonces;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PendingNonces;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PendingNonces;

  static equals(a: PendingNonces | PlainMessage<PendingNonces> | undefined, b: PendingNonces | PlainMessage<PendingNonces> | undefined): boolean;
}

