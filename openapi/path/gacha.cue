package path

import "github.com/yamato0211/brachio-backend/openapi/definition"

"/gachas": definition.#path & {
	get: {
		tags: ["gacha"]
		summary:     "ガチャ一覧取得"
		description: "ガチャ一覧を取得します。"
		operationId: "getGachaList"
		security: [#bearer]
		responses: {
			"200": {
				description: "OK"
				content: {
					"application/json": {
						schema: {
							type: "array"
							items: {
								$ref: "#/components/schemas/Gacha"
							}
						}
					}
				}
			}
		}
	}
}

"/gachas/{gachaId}": definition.#path & {
	post: {
		tags: ["gacha"]
		summary:     "ガチャを引く"
		description: "ガチャを引いてカードを取得します。"
		operationId: "postGachaDraw"
		security: [#bearer]
		parameters: [
			{
				name:        "gachaId"
				in:          "path"
				description: "ガチャID"
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
						$ref: "#/components/schemas/GachaDrawRequest"
					}
				}
			}
			required: true
		}
		responses: {
			"200": {
				description: "OK"
				content: {
					"application/json": {
						schema: {
							$ref: "#/components/schemas/GachaDrawResponse"
						}
					}
				}
			}
		}
	}
}

"/pack-power": definition.#path & {
	get: {
		tags: ["gacha"]
		summary:     "パックパワーの溜まり状況取得"
		description: "パックパワーの溜まり状況を取得します。"
		operationId: "getPackPower"
		security: [#bearer]
		responses: {
			"200": {
				description: "OK"
				content: {
					"application/json": {
						schema: {
							$ref: "#/components/schemas/PackPower"
						}
					}
				}
			}
		}
	}
}