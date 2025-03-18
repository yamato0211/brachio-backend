// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package schema

import (
	"encoding/json"
	"fmt"

	"github.com/oapi-codegen/runtime"
)

// Defines values for Element.
const (
	Alchohol   Element = "alchohol"
	Knowledge  Element = "knowledge"
	Money      Element = "money"
	Muscle     Element = "muscle"
	Null       Element = "null"
	Popularity Element = "popularity"
)

// Defines values for MasterCardType.
const (
	Goods     MasterCardType = "goods"
	Monster   MasterCardType = "monster"
	Supporter MasterCardType = "supporter"
)

// Defines values for MasterMonsterCardSubType.
const (
	Basic  MasterMonsterCardSubType = "basic"
	Stage1 MasterMonsterCardSubType = "stage1"
	Stage2 MasterMonsterCardSubType = "stage2"
)

// Defines values for SkillDamageOption.
const (
	Plus SkillDamageOption = "+"
	X    SkillDamageOption = "x"
)

// Defines values for GetCardsParamsIsAll.
const (
	N0 GetCardsParamsIsAll = 0
	N1 GetCardsParamsIsAll = 1
)

// Ability 特性
type Ability struct {
	// Name 特性名
	Name string `json:"name"`

	// Text 説明文
	Text string `json:"text"`
}

// Card defines model for card.
type Card struct {
	// Id カードID
	Id    *string `json:"id,omitempty"`
	union json.RawMessage
}

// CreateNewDeck201Response defines model for createNewDeck_201_response.
type CreateNewDeck201Response struct {
	// Id デッキID
	Id string `json:"id"`
}

// Deck defines model for deck.
type Deck struct {
	Cards    []Card    `json:"cards"`
	Color    Element   `json:"color"`
	Energies []Element `json:"energies"`

	// Id デッキID
	Id *string `json:"id,omitempty"`

	// Name デッキ名
	Name          string `json:"name"`
	ThumbnailCard Card   `json:"thumbnailCard"`
}

// DeckBase デッキ
type DeckBase struct {
	Color Element `json:"color"`

	// Id デッキID
	Id *string `json:"id,omitempty"`

	// Name デッキ名
	Name          string `json:"name"`
	ThumbnailCard Card   `json:"thumbnailCard"`
}

// DeckBaseWithId defines model for deckBaseWithId.
type DeckBaseWithId struct {
	Color Element `json:"color"`

	// Id デッキID
	Id *string `json:"id,omitempty"`

	// Name デッキ名
	Name          string `json:"name"`
	ThumbnailCard Card   `json:"thumbnailCard"`
}

// DeckWithId defines model for deckWithId.
type DeckWithId struct {
	Cards    []Card    `json:"cards"`
	Color    Element   `json:"color"`
	Energies []Element `json:"energies"`

	// Id デッキID
	Id *string `json:"id,omitempty"`

	// Name デッキ名
	Name          string `json:"name"`
	ThumbnailCard Card   `json:"thumbnailCard"`
}

// DrawGachaRequest defines model for drawGachaRequest.
type DrawGachaRequest struct {
	// IsTenDraw 10連ガチャかどうか
	IsTenDraw bool `json:"isTenDraw"`
}

// Element defines model for element.
type Element string

// Gacha defines model for gacha.
type Gacha struct {
	// Id ガチャID
	Id *string `json:"id,omitempty"`

	// ImageUrl ガチャ画像URL
	ImageUrl string `json:"imageUrl"`

	// Name ガチャ名
	Name string `json:"name"`
}

// GachaPower defines model for gachaPower.
type GachaPower struct {
	// Charged 現在のガチャパワー
	Charged int `json:"charged"`

	// Next 次のガチャパワーが貯まるまでの秒数
	Next int `json:"next"`
}

// Item defines model for item.
type Item struct {
	// Count 所持数
	Count int `json:"count"`

	// Id アイテムID（アイテムごとに一意）
	Id string `json:"id"`

	// ImageUrl 画像URL
	ImageUrl string `json:"imageUrl"`

	// Name アイテム名
	Name string `json:"name"`
}

// MasterCard defines model for masterCard.
type MasterCard struct {
	union json.RawMessage
}

// MasterCardBase 全カード共通のプロパティ
type MasterCardBase struct {
	CardType MasterCardType `json:"cardType"`

	// Expansion カードセット名
	Expansion *string `json:"expansion,omitempty"`

	// ImageUrl カード画像URL
	ImageUrl string `json:"imageUrl"`

	// MasterCardId カードID
	MasterCardId *string `json:"masterCardId,omitempty"`

	// Name カード名
	Name string `json:"name"`

	// Rarity レアリティ
	Rarity *int `json:"rarity,omitempty"`
}

// MasterCardType defines model for masterCardType.
type MasterCardType string

// MasterCardWithCount defines model for masterCardWithCount.
type MasterCardWithCount struct {
	// Count カード枚数
	Count      int        `json:"count"`
	MasterCard MasterCard `json:"masterCard"`
}

// MasterGoodsCard defines model for masterGoodsCard.
type MasterGoodsCard struct {
	CardType MasterCardType `json:"cardType"`

	// Expansion カードセット名
	Expansion *string `json:"expansion,omitempty"`

	// ImageUrl カード画像URL
	ImageUrl string `json:"imageUrl"`

	// MasterCardId カードID
	MasterCardId *string `json:"masterCardId,omitempty"`

	// Name カード名
	Name string `json:"name"`

	// Rarity レアリティ
	Rarity *int `json:"rarity,omitempty"`

	// Text 説明文
	Text string `json:"text"`
}

// MasterMonsterCard defines model for masterMonsterCard.
type MasterMonsterCard struct {
	// Ability 特性
	Ability  *Ability       `json:"ability,omitempty"`
	CardType MasterCardType `json:"cardType"`
	Element  Element        `json:"element"`

	// EvolvesFrom 進化元
	EvolvesFrom *[]string `json:"evolvesFrom,omitempty"`

	// EvolvesTo 進化先
	EvolvesTo *[]string `json:"evolvesTo,omitempty"`

	// Expansion カードセット名
	Expansion *string `json:"expansion,omitempty"`

	// Hp HP
	Hp int `json:"hp"`

	// ImageUrl カード画像URL
	ImageUrl string `json:"imageUrl"`

	// IsEx EXカードかどうか
	IsEx *bool `json:"isEx,omitempty"`

	// MasterCardId カードID
	MasterCardId *string `json:"masterCardId,omitempty"`

	// Name カード名
	Name string `json:"name"`

	// Rarity レアリティ
	Rarity int `json:"rarity"`

	// RetreatCost 逃げるコスト
	RetreatCost *int `json:"retreatCost,omitempty"`

	// Skills ワザ
	Skills   []Skill                   `json:"skills"`
	SubType  *MasterMonsterCardSubType `json:"subType,omitempty"`
	Weakness Element                   `json:"weakness"`
}

// MasterMonsterCardSubType defines model for MasterMonsterCard.SubType.
type MasterMonsterCardSubType string

// MasterSupporterCard defines model for masterSupporterCard.
type MasterSupporterCard struct {
	CardType MasterCardType `json:"cardType"`

	// Expansion カードセット名
	Expansion *string `json:"expansion,omitempty"`

	// ImageUrl カード画像URL
	ImageUrl string `json:"imageUrl"`

	// MasterCardId カードID
	MasterCardId *string `json:"masterCardId,omitempty"`

	// Name カード名
	Name string `json:"name"`

	// Rarity レアリティ
	Rarity *int `json:"rarity,omitempty"`

	// Text 説明文
	Text string `json:"text"`
}

// MyCardList defines model for myCardList.
type MyCardList = []MasterCardWithCount

// Pack defines model for pack.
type Pack struct {
	Cards *[]MasterCard `json:"cards,omitempty"`
}

// Skill ワザ
type Skill struct {
	// Cost コスト
	Cost []Element `json:"cost"`

	// Damage ダメージ
	Damage int `json:"damage"`

	// DamageOption x or +
	DamageOption *SkillDamageOption `json:"damageOption,omitempty"`

	// Name ワザ名
	Name string `json:"name"`

	// Text 説明文
	Text string `json:"text"`
}

// SkillDamageOption x or +
type SkillDamageOption string

// UpdateDeck デッキ更新リクエスト
type UpdateDeck struct {
	Color    Element   `json:"color"`
	Energies []Element `json:"energies"`

	// MasterCardIds マスターカードID
	MasterCardIds []string `json:"masterCardIds"`

	// Name デッキ名
	Name string `json:"name"`

	// ThumbnailCardId サムネイルカードID
	ThumbnailCardId string `json:"thumbnailCardId"`
}

// GetCardsParams defines parameters for GetCards.
type GetCardsParams struct {
	// IsAll 全件取得フラグ（0: 非全件, 1: 全件)
	IsAll *GetCardsParamsIsAll `form:"is_all,omitempty" json:"is_all,omitempty"`
}

// GetCardsParamsIsAll defines parameters for GetCards.
type GetCardsParamsIsAll int

// UpdateDeckJSONRequestBody defines body for UpdateDeck for application/json ContentType.
type UpdateDeckJSONRequestBody = UpdateDeck

// DrawGachaJSONRequestBody defines body for DrawGacha for application/json ContentType.
type DrawGachaJSONRequestBody = DrawGachaRequest

// AsMasterMonsterCard returns the union data inside the Card as a MasterMonsterCard
func (t Card) AsMasterMonsterCard() (MasterMonsterCard, error) {
	var body MasterMonsterCard
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMasterMonsterCard overwrites any union data inside the Card as the provided MasterMonsterCard
func (t *Card) FromMasterMonsterCard(v MasterMonsterCard) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMasterMonsterCard performs a merge with any union data inside the Card, using the provided MasterMonsterCard
func (t *Card) MergeMasterMonsterCard(v MasterMonsterCard) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMasterSupporterCard returns the union data inside the Card as a MasterSupporterCard
func (t Card) AsMasterSupporterCard() (MasterSupporterCard, error) {
	var body MasterSupporterCard
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMasterSupporterCard overwrites any union data inside the Card as the provided MasterSupporterCard
func (t *Card) FromMasterSupporterCard(v MasterSupporterCard) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMasterSupporterCard performs a merge with any union data inside the Card, using the provided MasterSupporterCard
func (t *Card) MergeMasterSupporterCard(v MasterSupporterCard) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMasterGoodsCard returns the union data inside the Card as a MasterGoodsCard
func (t Card) AsMasterGoodsCard() (MasterGoodsCard, error) {
	var body MasterGoodsCard
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMasterGoodsCard overwrites any union data inside the Card as the provided MasterGoodsCard
func (t *Card) FromMasterGoodsCard(v MasterGoodsCard) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMasterGoodsCard performs a merge with any union data inside the Card, using the provided MasterGoodsCard
func (t *Card) MergeMasterGoodsCard(v MasterGoodsCard) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Card) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	if err != nil {
		return nil, err
	}
	object := make(map[string]json.RawMessage)
	if t.union != nil {
		err = json.Unmarshal(b, &object)
		if err != nil {
			return nil, err
		}
	}

	object["id"], err = json.Marshal(t.Id)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'id': %w", err)
	}

	b, err = json.Marshal(object)
	return b, err
}

func (t *Card) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	object := make(map[string]json.RawMessage)
	err = json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["id"]; found {
		err = json.Unmarshal(raw, &t.Id)
		if err != nil {
			return fmt.Errorf("error reading 'id': %w", err)
		}
	}

	return err
}

// AsMasterMonsterCard returns the union data inside the MasterCard as a MasterMonsterCard
func (t MasterCard) AsMasterMonsterCard() (MasterMonsterCard, error) {
	var body MasterMonsterCard
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMasterMonsterCard overwrites any union data inside the MasterCard as the provided MasterMonsterCard
func (t *MasterCard) FromMasterMonsterCard(v MasterMonsterCard) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMasterMonsterCard performs a merge with any union data inside the MasterCard, using the provided MasterMonsterCard
func (t *MasterCard) MergeMasterMonsterCard(v MasterMonsterCard) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMasterSupporterCard returns the union data inside the MasterCard as a MasterSupporterCard
func (t MasterCard) AsMasterSupporterCard() (MasterSupporterCard, error) {
	var body MasterSupporterCard
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMasterSupporterCard overwrites any union data inside the MasterCard as the provided MasterSupporterCard
func (t *MasterCard) FromMasterSupporterCard(v MasterSupporterCard) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMasterSupporterCard performs a merge with any union data inside the MasterCard, using the provided MasterSupporterCard
func (t *MasterCard) MergeMasterSupporterCard(v MasterSupporterCard) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMasterGoodsCard returns the union data inside the MasterCard as a MasterGoodsCard
func (t MasterCard) AsMasterGoodsCard() (MasterGoodsCard, error) {
	var body MasterGoodsCard
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMasterGoodsCard overwrites any union data inside the MasterCard as the provided MasterGoodsCard
func (t *MasterCard) FromMasterGoodsCard(v MasterGoodsCard) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMasterGoodsCard performs a merge with any union data inside the MasterCard, using the provided MasterGoodsCard
func (t *MasterCard) MergeMasterGoodsCard(v MasterGoodsCard) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MasterCard) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MasterCard) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
