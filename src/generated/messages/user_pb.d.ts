// @generated by protoc-gen-es v2.0.0
// @generated from file messages/user.proto (package messages.user, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file messages/user.proto.
 */
export declare const file_messages_user: GenFile;

/**
 * @generated from message messages.user.User
 */
export declare type User = Message<"messages.user.User"> & {
  /**
   * ユーザーID
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * ユーザー名
   *
   * @generated from field: string name = 2;
   */
  name: string;
};

/**
 * Describes the message messages.user.User.
 * Use `create(UserSchema)` to create a new message.
 */
export declare const UserSchema: GenMessage<User>;

