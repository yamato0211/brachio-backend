// The PokePoke API
package api

import "net"

info: {
	title:   *"The PokePoke API" | string
	version: *"0.0.1" | string
}
// 特性
#Ability: {
	// 特性名
	name?: string

	// 特性の説明
	description?: string
	...
}

#Card: {
	// カードID
	id?: string

	// カード名
	name?: string

	// レアリティ
	rarity?: int

	// カードナンバー
	cardNumber?: int

	// 拡張パック名
	expansion?: string

	// カードタイプ
	cardType?: "Monster" | "Supporter" | "Goods"

	// カード画像URL
	imageUrl?: net.AbsURL

	// カードの効果
	rules?: [...string]
	...
}

// 属性
#Element: "Grass" | "Fire" | "Water" | "Lightning" | "Psychic" | "Fighting" | "Darkness" | "Metal" | "Dragon" | "Normal"

#Gacha: {
	// ガチャID
	id?: string

	// ガチャ名
	name?: string

	// ガチャ画像URL
	imageUrl?: net.AbsURL
	...
}

#GachaDrawRequest: {
	// 10連ガチャかどうか
	isTenDraw?: bool
	...
}

#GachaDrawResponse: {
	// ガチャで引いたカード
	packs?: [...#Pack]
	...
}

#MonsterCard: {
	// カードID
	id?: string

	// カード名
	name?: string

	// レアリティ
	rarity?: int

	// カードナンバー
	cardNumber?: int

	// 拡張パック名
	expansion?: string

	// カードタイプ
	cardType?: "Monster" | "Supporter" | "Goods"

	// カード画像URL
	imageUrl?: net.AbsURL

	// カードの効果
	rules?: [...string]

	// 進化段階 (たね, 1進化, 2進化)
	subType?: "Basic" | "Stage1" | "Stage2"
	type?:    #Element

	// HP
	hp?: int

	// ワザ
	skills?: [...#Skill]
	weekness?: #Element

	// にげるコスト
	retreatCost?: int
	...
}

// カード5枚セット
#Pack: {
	// 中身
	cards?: [...#Card]
	...
}

// ワザ
#Skill: {
	// ワザ名
	name?: string

	// ワザの説明
	text?: string

	// ダメージ
	damage?: int

	// コスト
	cost?: [...#Element]
	...
}

#User: {
	// ユーザーID
	id?: string

	// ユーザー名
	name?: string
	...
}
