package schema

import "github.com/yamato0211/brachio-backend/openapi/definition"

Gacha: definition.#schema & {
	type: "object"
	properties: {
		id: {
			type:        "string"
			description: "ガチャID"
			readOnly:    true
		}
		name: {
			type:        "string"
			description: "ガチャ名"
		}
		imageUrl: {
			type:        "string"
			description: "ガチャ画像URL"
			format:      "uri"
		}
	}
}

Pack: definition.#schema & {
	type:        "object"
	description: "カード5枚セット"
	properties: {
		cards: {
			type:        "array"
			description: "中身"
			items: {
				$ref: "#/components/schemas/Card"
			}
		}
	}
}

GachaDrawRequest: definition.#schema & {
	type: "object"
	properties: {
		isTenDraw: {
			type:        "boolean"
			description: "10連ガチャかどうか"
		}
	}
}

GachaDrawResponse: definition.#schema & {
	type: "object"
	properties: {
		packs: {
			type:        "array"
			description: "ガチャで引いたカード"
			items: {
				$ref: "#/components/schemas/Pack"
			}
		}
	}
}

PackPower: definition.#schema & {
	type: "object"
	properties: {
		next: {
			type: "integer"
			description: "次のパックが貯まるまでの秒数"
		}
		charged: {
			type: "integer"
			description: "現在溜まっているパックの数"
		}
	}
}