package path

import "github.com/yamato0211/brachio-backend/openapi/definition"

"/items:": definition.#path & {
  get: {
    tags: ["item"]
    summary:     "所持アイテム一覧取得"
    description: "所持アイテム一覧を取得します。"
    operationId: "getMyItemList"
    security: [#bearer]
    responses: {
      "200": {
        description: "アイテム一覧"
        content: {
          "application/json": {
            schema: {
              type: "array"
              items: {
                $ref: "#/components/schemas/Item"
              }
            }
          }
        }
      }
    }
  }
}