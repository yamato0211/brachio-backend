package service

import (
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

const (
	WinPoint = 3
)

type GameMasterService interface {
	InitializeGame(state *model.GameState) error
	FlipCoin() bool
	ShuffleDeck(deck []*model.Card) []*model.Card
	DrawCards(player *model.Player, count int)
	RunEffect(effects []*model.Effect, trigger string, args any) ([]*model.Effect, error)
}

type gameMasterService struct{}

func NewGameMasterService() GameMasterService {
	return &gameMasterService{}
}

func (s *gameMasterService) InitializeGame(state *model.GameState) error {
	// コイントスを行い、先攻を決める
	// 表なら Player1 が先攻、裏なら Player2 が先攻
	if !s.FlipCoin() {
		state.TurnPlayer, state.NonTurnPlayer = state.NonTurnPlayer, state.TurnPlayer
	}

	// プレイヤーの初期化
	if err := s.initializePlayer(state.TurnPlayer); err != nil {
		return err
	}
	if err := s.initializePlayer(state.NonTurnPlayer); err != nil {
		return err
	}

	return nil
}

// FlipCoin はコイントスを行い、表なら true 裏なら false を返す
func (s *gameMasterService) FlipCoin() bool {
	return rand.IntN(2) == 0
}

func (s *gameMasterService) initializePlayer(player *model.Player) error {
	cards := make([]*model.Card, 0, len(player.BaseDeck.MasterCards))
	for i, masterCard := range player.BaseDeck.MasterCards {
		var monsterID string
		if masterCard.CardType == model.CardTypeMonster {
			monsterID = uuid.New().String()
		}

		cards = append(cards, &model.Card{
			CardID:            model.NewCardID(i),
			MasterCard:        masterCard,
			ReservedMonsterID: monsterID,
		})
	}
	player.Deck = cards

	s.ShuffleDeck(player.Deck)

	// デッキから手札にカードを 5 枚引く
	player.Hands, player.Deck = player.Deck[:5], player.Deck[5:]

	// 初ターンのエネルギーを抽選する
	s.lotteryNextEnergy(player)

	return nil
}

func (s *gameMasterService) ShuffleDeck(deck []*model.Card) []*model.Card {
	mutable.Shuffle(deck)
	return deck
}

func (s *gameMasterService) lotteryNextEnergy(player *model.Player) {
	nextEnergy := &player.BaseDeck.Energies[rand.IntN(len(player.BaseDeck.Energies))]
	player.CurrentEnergy, player.NextEnergy = player.NextEnergy, nextEnergy
}

func (s *gameMasterService) ChangeTurn(state *model.GameState) {
	state.Turn++

	state.TurnPlayer, state.NonTurnPlayer = state.NonTurnPlayer, state.TurnPlayer

	s.DrawCards(state.TurnPlayer, 1)
}

func (s *gameMasterService) IsFinish(state *model.GameState) bool {
	if state.TurnPlayer.Point >= WinPoint {
		return true
	}
	if state.NonTurnPlayer.Point >= WinPoint {
		return true
	}
	return false
}

func (s *gameMasterService) DrawCards(player *model.Player, count int) {
	if len(player.Deck) == 0 {
		return
	}

	count = min(count, len(player.Deck))

	cards := player.Deck[:count]
	player.Deck = player.Deck[count:]
	player.Hands = append(player.Hands, cards...)
}

func (s *gameMasterService) RunEffect(effects []*model.Effect, trigger string, args any) ([]*model.Effect, error) {
	for _, effect := range effects {
		if effect.Trigger == trigger {
			isDelete, err := effect.Fn(args)
			if err != nil {
				return nil, err
			}

			if isDelete {
				effect = nil
			}
		}
	}

	return lo.Filter(effects, func(effect *model.Effect, _ int) bool { return effect != nil }), nil
}
