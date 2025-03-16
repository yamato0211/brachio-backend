package path

import "github.com/yamato0211/brachio-backend/openapi/definition"

"/cards:": definition.#path & {
  get: {
    tags: ["card"]
    summary:     "所持カード一覧取得"
    description: "所持カード一覧を取得します。"
    operationId: "getMyCardList"
    security: [#bearer]
    responses: {
      "200": {
        description: "カード一覧"
        content: {
          "application/json": {
            schema: {
              type: "array"
              items: {
                $ref: "#/components/schemas/Card"
              }
            }
          }
        }
      }
    }
  }
}

"/cards/{cardNumber}:": definition.#path & {
  get: {
    tags: ["card"]
    summary:     "カード取得"
    description: "カードを取得します。"
    operationId: "getMyCard"
    security: [#bearer]
    parameters: [
      {
        name:        "cardNumber"
        in:          "path"
        description: "カード番号"
        required:    true
        schema: {
          type: "string"
        }
      },
    ]
    responses: {
      "200": {
        description: "カード"
        content: {
          "application/json": {
            schema: {
              $ref: "#/components/schemas/Card"
            }
          }
        }
      }
    }
  }
}