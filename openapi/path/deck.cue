package path

import "github.com/yamato0211/brachio-backend/openapi/definition"

"/decks": definition.#path & {
  get: {
    tags: ["deck"]
    summary:     "デッキ一覧取得"
    description: "デッキ一覧を取得します。"
    operationId: "getMyDeckList"
    security: [#bearer]
    responses: {
      "200": {
        description: "デッキ一覧"
        content: {
          "application/json": {
            schema: {
              type: "array"
              items: {
                $ref: "#/components/schemas/Deck"
              }
            }
          }
        }
      }
    }
  }
  post: {
    tags: ["deck"]
    summary:     "デッキ作成"
    description: "デッキを作成します。"
    operationId: "postMyDeck"
    security: [#bearer]
    requestBody: {
      content: {
        "application/json": {
          schema: {
            $ref: "#/components/schemas/Deck"
          }
        }
      }
      required: true
    }
    responses: {
      "201": {
        description: "デッキ作成成功"
        content: {}
      }
    }
  }
}

"/decks/{deckId}": definition.#path & {
  get: {
    tags: ["deck"]
    summary:     "デッキ取得"
    description: "デッキを取得します。"
    operationId: "getMyDeck"
    security: [#bearer]
    parameters: [
      {
        name:        "deckId"
        in:          "path"
        description: "デッキID"
        required:    true
        schema: {
          type: "string"
        }
      },
    ]
    responses: {
      "200": {
        description: "デッキ"
        content: {
          "application/json": {
            schema: {
              $ref: "#/components/schemas/Deck"
            }
          }
        }
      }
    }
  }
  put: {
    tags: ["deck"]
    summary:     "デッキ編集"
    description: "デッキを編集します。"
    operationId: "putMyDeck"
    security: [#bearer]
    parameters: [
      {
        name:        "deckId"
        in:          "path"
        description: "デッキID"
        required:    true
        schema: {
          type: "string"
        }
      },
    ]
    requestBody: {
      content: {
        "application/json": {
          schema: {
            $ref: "#/components/schemas/Deck"
          }
        }
      }
      required: true
    }
    responses: {
      "200": {
        description: "デッキ作成成功"
        content: {}
      }
    }
  }
}