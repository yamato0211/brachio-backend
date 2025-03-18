// @generated by protoc-gen-es v2.0.0
// @generated from file messages/item.proto (package messages.item, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file messages/item.proto.
 */
export declare const file_messages_item: GenFile;

/**
 * アイテム
 *
 * @generated from message messages.item.Item
 */
export declare type Item = Message<"messages.item.Item"> & {
  /**
   * アイテムID（各アイテムごとに一意）
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * アイテム名
   *
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * 所持数
   *
   * @generated from field: int32 count = 3;
   */
  count: number;
};

/**
 * Describes the message messages.item.Item.
 * Use `create(ItemSchema)` to create a new message.
 */
export declare const ItemSchema: GenMessage<Item>;

