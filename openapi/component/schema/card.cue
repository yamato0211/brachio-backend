package schema

import "github.com/yamato0211/brachio-backend/openapi/definition"

CardBase: definition.#schema & {
	type: "object"
	properties: {
		id: {
			type:        "string"
			description: "全ユーザーで一意に定まるID"
			readOnly:    true
		}
		name: {
			type:        "string"
			description: "カード名"
		}
		rarity: {
			type:        "integer"
			description: "レアリティ"
		}
		cardNumber: {
			type:        "integer"
			description: "カードナンバー"
		}
		expansion: {
			type:        "string"
			description: "拡張パック名"
		}
		cardType: {
			type:        "string"
			description: "カードタイプ"
			enum: ["Monster", "Supporter", "Goods"]
		}
		imageUrl: {
			type:        "string"
			description: "カード画像URL"
			format:      "uri"
		}
	}
}

Element: definition.#schema & {
	type:        "string"
	description: "属性"
	enum: [
		"Grass",
		"Fire",
		"Water",
		"Lightning",
		"Psychic",
		"Fighting",
		"Darkness",
		"Metal",
		"Dragon",
		"Normal",
	]
}

Skill: definition.#schema & {
	type:        "object"
	description: "ワザ"
	properties: {
		name: {
			type:        "string"
			description: "ワザ名"
		}
		text: {
			type:        "string"
			description: "ワザの説明"
		}
		damage: {
			type:        "integer"
			description: "ダメージ"
		}
		damageOption?: {
			type:        "string"
			description: "ダメージオプション"
			examples: ["x", "+"]
		}
		cost: {
			type:        "array"
			description: "コスト"
			items: {
				$ref: "#/components/schemas/Element"
			}
		}
	}
}

Ability: definition.#schema & {
	type:        "object"
	description: "特性"
	properties: {
		name: {
			type:        "string"
			description: "特性名"
		}
		description: {
			type:        "string"
			description: "特性の説明"
		}
	}
}

MonsterCard: definition.#schema & CardBase & {
	properties: {
		subType: {
			type:        "string"
			description: "進化段階 (たね, 1進化, 2進化)"
			enum: ["Basic", "Stage1", "Stage2"]
		}
    type: {
      $ref: "#/components/schemas/Element"
    }
		hp: {
			type:        "integer"
			description: "HP"
		}
		skills: {
			type:        "array"
			description: "ワザ"
			items: {
				$ref: "#/components/schemas/Skill"
			}
		}
		weekness: {
			$ref: "#/components/schemas/Element"
		}
		ability?: {
			$ref: "#/components/schemas/Ability"
		}
		retreatCost: {
			type:        "integer"
			description: "にげるコスト"
		}
		evolvesFrom?: {
			type:        "string"
			description: "進化元のカード名（存在する場合のみ）"
		}
		evolvesTo?: {
			type:        "string"
			description: "進化先のカード名（存在する場合のみ）"
		}
	}
}

SupporterCard: definition.#schema & CardBase & {
  properties: {
    effect: {
      type:        "string"
      description: "効果"
    }
  }
}

GoodsCard: definition.#schema & CardBase & {
  properties: {
    effect: {
      type:        "string"
      description: "効果"
    }
  }
}

Card: definition.#oneOf & {
  oneOf: [
    { $ref: "#/components/schemas/MonsterCard" },
    { $ref: "#/components/schemas/SupporterCard" },
    { $ref: "#/components/schemas/GoodsCard" },
  ]
}