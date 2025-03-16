package schema

import "github.com/yamato0211/brachio-backend/openapi/definition"

Deck: definition.#schema & {
  type: "object"
  properties: {
    id: {
      type:        "string"
      description: "デッキID"
      readOnly:    true
    }
    name: {
      type:        "string"
      description: "デッキ名"
    }
    elements: {
      type:        "array"
      description: "エネルギーの属性"
      items: {
        $ref: "#/components/schemas/Element"
      }
    }
    cards: {
      type:        "array"
      description: "カードリスト"
      items: {
        $ref: "#/components/schemas/Card"
      }
    }
  }
}

CreateDeckRequest: definition.#schema & {
  type: "object"
  properties: {
    name: {
      type:        "string"
      description: "デッキ名"
    }
    elements: {
      type:        "array"
      description: "エネルギーの属性"
      items: {
        $ref: "#/components/schemas/Element"
      }
    }
    cardIds: {
      type:        "array"
      description: "カードIDリスト"
      items: {
        type: "string"
      }
    }
  }
}