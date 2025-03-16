package schema

import "github.com/yamato0211/brachio-backend/openapi/definition"

User: definition.#schema & {
	type: "object"
	properties: {
		id: {
			type:        "string"
			description: "ユーザーID"
			readOnly:    true
		}
		name: {
			type:        "string"
			description: "ユーザー名"
		}
	}
}
