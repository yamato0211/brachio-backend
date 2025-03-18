// @generated by protoc-gen-es v2.0.0
// @generated from file messages/pack_power.proto (package messages.pack_power, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file messages/pack_power.proto.
 */
export declare const file_messages_pack_power: GenFile;

/**
 * @generated from message messages.pack_power.PackPower
 */
export declare type PackPower = Message<"messages.pack_power.PackPower"> & {
  /**
   * 次のパックが貯まるまでの秒数
   *
   * @generated from field: int32 next = 1;
   */
  next: number;

  /**
   * 現在溜まっているパックの数
   *
   * @generated from field: int32 charged = 2;
   */
  charged: number;
};

/**
 * Describes the message messages.pack_power.PackPower.
 * Use `create(PackPowerSchema)` to create a new message.
 */
export declare const PackPowerSchema: GenMessage<PackPower>;

