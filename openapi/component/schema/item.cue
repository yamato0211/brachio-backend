package schema

import "github.com/yamato0211/brachio-backend/openapi/definition"

Item: definition.#schema & {
  type: "object"
  properties: {
    id: {
      type:       "string"
      description: "アイテムID（アイテムごとに一意）"
    }
    name: {
      type:        "string"
      description: "アイテム名"
    }
    count: {
      type:        "integer"
      description: "所持数"
    }
  }
}