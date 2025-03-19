package service

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	"golang.org/x/xerrors"
)

const (
	WinPoint = 3
)

type GameMasterService interface {
	InitializeGame(state *model.GameState) error
	FlipCoin() bool
	ShuffleDeck(deck []*model.Card) []*model.Card
	DrawCards(player *model.Player, count int) []*model.Card
	RunEffect(state *model.GameState, effects []*model.Effect, trigger string, args any) ([]*model.Effect, error)
	Matched(ctx context.Context, roomID model.RoomID) error
	ReadyForStart(ctx context.Context, roomID model.RoomID, userID model.UserID) error
	ChangeTurn(ctx context.Context, roomID model.RoomID) error
}

type gameMasterService struct {
	GameStateRepository repository.GameStateRepository
	GameEventSender     GameEventSender
}

func NewGameMasterService(
	gsr repository.GameStateRepository,
	ges GameEventSender,
) GameMasterService {
	return &gameMasterService{
		GameStateRepository: gsr,
		GameEventSender:     ges,
	}
}

func (s *gameMasterService) InitializeGame(state *model.GameState) error {
	// コイントスを行い、先攻を決める
	if s.FlipCoin() {
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
		fmt.Printf("masterCard: %+v\n", masterCard)
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

	// デッキをシャッフルする
	s.ShuffleDeck(player.Deck)

	var hand []*model.Card
	// 初手札に SubType が model.SubTypeBasic のカードが含まれていない場合、再抽選する
	for {
		if len(player.Deck) < 5 {
			return fmt.Errorf("デッキ内のカード数が足りません")
		}
		hand = player.Deck[:5]
		// 手札の中に SubType が model.SubTypeBasic のカードがあるかチェック
		hasBasic := lo.SomeBy(hand, func(card *model.Card) bool {
			return card.MasterCard.SubType == model.MonsterSubTypeBasic
		})
		if hasBasic {
			break
		}
		// 含まれていなければ、デッキ全体を再シャッフルして再抽選
		s.ShuffleDeck(player.Deck)
	}

	// デッキから手札にカードを 5 枚引く
	player.Hands, player.Deck = hand, player.Deck[5:]

	// 初ターンのエネルギーを抽選する
	s.lotteryNextEnergy(player)

	return nil
}

func (s *gameMasterService) ShuffleDeck(deck []*model.Card) []*model.Card {
	mutable.Shuffle(deck)
	return deck
}

func (s *gameMasterService) lotteryNextEnergy(player *model.Player) {
	if !player.NextEnergy.IsZero() {
		player.CurrentEnergies = append(player.CurrentEnergies, player.NextEnergy)
	}
	player.NextEnergy = player.BaseDeck.Energies[rand.IntN(len(player.BaseDeck.Energies))]
}

func (s *gameMasterService) ChangeTurn(ctx context.Context, roomID model.RoomID) error {
	err := s.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := s.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		state.Turn++
		state.TurnPlayer, state.NonTurnPlayer = state.NonTurnPlayer, state.TurnPlayer
		s.lotteryNextEnergy(state.TurnPlayer)
		drawedCard := s.DrawCards(state.TurnPlayer, 1)

		if err := s.GameEventSender.SendTurnStartEvent(ctx, state.TurnPlayer.UserID, state.TurnPlayer.UserID); err != nil {
			return err
		}
		if err := s.GameEventSender.SendTurnStartEvent(context.Background(), state.NonTurnPlayer.UserID, state.TurnPlayer.UserID); err != nil {
			return err
		}

		if err := s.GameEventSender.SendDrawCardsEventToActor(ctx, state.TurnPlayer.UserID, len(state.TurnPlayer.Deck), drawedCard...); err != nil {
			return err
		}
		if err := s.GameEventSender.SendDrawCardsEventToRecipient(ctx, state.NonTurnPlayer.UserID, len(state.TurnPlayer.Deck), drawedCard...); err != nil {
			return err
		}

		if err := s.GameEventSender.SendNextEnergyEventToActor(ctx, state.TurnPlayer.UserID, state.TurnPlayer.NextEnergy); err != nil {
			return err
		}
		if err := s.GameEventSender.SendNextEnergyEventToRecipient(ctx, state.NonTurnPlayer.UserID, state.TurnPlayer.NextEnergy); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return err
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

func (s *gameMasterService) DrawCards(player *model.Player, count int) []*model.Card {
	if len(player.Deck) == 0 {
		return nil
	}

	count = min(count, len(player.Deck))

	cards := player.Deck[:count]
	player.Deck = player.Deck[count:]
	player.Hands = append(player.Hands, cards...)

	return cards
}

func (s *gameMasterService) RunEffect(state *model.GameState, effects []*model.Effect, trigger string, args any) ([]*model.Effect, error) {
	for _, effect := range effects {
		if effect.Trigger == trigger {
			isDelete, err := effect.Fn(state, args)
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

func (s *gameMasterService) ReadyForStart(ctx context.Context, roomID model.RoomID, userID model.UserID) error {
	err := s.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := s.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		me, err := state.FindMeByUserID(userID)
		if err != nil {
			return err
		}
		me.IsReady = true

		enemy, err := state.FindEnemyByUserID(userID)
		if err != nil {
			return err
		}
		if enemy.IsReady {
			state.Phase = model.GamePhaseBattle
		}

		drawedCard := s.DrawCards(state.TurnPlayer, 1)

		if err := s.GameStateRepository.Store(ctx, state); err != nil {
			return err
		}

		if !enemy.IsReady {
			return nil
		}

		var events []*messages.Effect
		events = append(events, &messages.Effect{
			Effect: &messages.Effect_Summon{
				Summon: &messages.SummonEffect{
					Card:     messages.NewCard(state.TurnPlayer.BattleMonster.BaseCard),
					Position: int32(0),
				},
			},
		})
		for i, monster := range state.TurnPlayer.BenchMonsters {
			if monster == nil {
				continue
			}
			events = append(events, &messages.Effect{
				Effect: &messages.Effect_Summon{
					Summon: &messages.SummonEffect{
						Card:     messages.NewCard(monster.BaseCard),
						Position: int32(i + 1),
					},
				},
			})
		}

		if err := s.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.NonTurnPlayer.UserID, events...); err != nil {
			return err
		}

		events = nil
		events = append(events, &messages.Effect{
			Effect: &messages.Effect_Summon{
				Summon: &messages.SummonEffect{
					Card:     messages.NewCard(state.NonTurnPlayer.BattleMonster.BaseCard),
					Position: int32(0),
				},
			},
		})
		for i, monster := range state.NonTurnPlayer.BenchMonsters {
			if monster == nil {
				continue
			}
			events = append(events, &messages.Effect{
				Effect: &messages.Effect_Summon{
					Summon: &messages.SummonEffect{
						Card:     messages.NewCard(monster.BaseCard),
						Position: int32(i + 1),
					},
				},
			})
		}

		if err := s.GameEventSender.SendDrawEffectEventToRecipient(ctx, state.TurnPlayer.UserID, events...); err != nil {
			return err
		}

		if err := s.GameEventSender.SendGameStartEvent(ctx, state.TurnPlayer.UserID); err != nil {
			return err
		}

		if err := s.GameEventSender.SendTurnStartEvent(ctx, state.TurnPlayer.UserID, state.TurnPlayer.UserID); err != nil {
			return err
		}
		if err := s.GameEventSender.SendTurnStartEvent(context.Background(), state.NonTurnPlayer.UserID, state.TurnPlayer.UserID); err != nil {
			return err
		}

		if err := s.GameEventSender.SendDrawCardsEventToActor(ctx, state.TurnPlayer.UserID, len(state.TurnPlayer.Deck), drawedCard...); err != nil {
			return err
		}
		if err := s.GameEventSender.SendDrawCardsEventToRecipient(ctx, state.NonTurnPlayer.UserID, len(state.TurnPlayer.Deck), drawedCard...); err != nil {
			return err
		}

		if err := s.GameEventSender.SendNextEnergyEventToActor(ctx, state.TurnPlayer.UserID, state.TurnPlayer.NextEnergy); err != nil {
			return err
		}
		if err := s.GameEventSender.SendNextEnergyEventToRecipient(ctx, state.NonTurnPlayer.UserID, state.TurnPlayer.NextEnergy); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *gameMasterService) Matched(ctx context.Context, roomID model.RoomID) error {
	var users []*model.User
	err := s.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := s.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}
		if state.TurnPlayer == nil || state.NonTurnPlayer == nil {
			return xerrors.Errorf("player not found")
		}

		fmt.Printf("state: %+v\n", state)
		fmt.Printf("deck: %+v\n", state.TurnPlayer.BaseDeck)
		fmt.Printf("deck: %+v\n", state.NonTurnPlayer.BaseDeck)

		users = []*model.User{
			{ID: state.TurnPlayer.UserID, Name: "Player1"},
			{ID: state.NonTurnPlayer.UserID, Name: "Player2"},
		}

		// プレイヤーの初期化
		if err := s.InitializeGame(state); err != nil {
			return err
		}

		if err := s.GameEventSender.SendMatchingComplete(ctx, users[0].ID, users[1].ID, roomID); err != nil {
			return err
		}
		if err := s.GameEventSender.SendMatchingComplete(ctx, users[1].ID, users[0].ID, roomID); err != nil {
			return err
		}

		if err := s.GameEventSender.SendDecideOrderEvent(ctx, users[0].ID, users[0].ID, users[1].ID); err != nil {
			return err
		}
		if err := s.GameEventSender.SendDecideOrderEvent(ctx, users[1].ID, users[0].ID, users[1].ID); err != nil {
			return err
		}

		// カード配布
		if err := s.GameEventSender.SendDrawCardsEventToActor(ctx, state.TurnPlayer.UserID, len(state.TurnPlayer.Deck), state.TurnPlayer.Hands...); err != nil {
			return err
		}
		if err := s.GameEventSender.SendDrawCardsEventToRecipient(ctx, state.NonTurnPlayer.UserID, len(state.TurnPlayer.Deck), state.TurnPlayer.Hands...); err != nil {
			return err
		}

		if err := s.GameEventSender.SendDrawCardsEventToActor(ctx, state.NonTurnPlayer.UserID, len(state.NonTurnPlayer.Deck), state.NonTurnPlayer.Hands...); err != nil {
			return err
		}
		if err := s.GameEventSender.SendDrawCardsEventToRecipient(ctx, state.TurnPlayer.UserID, len(state.NonTurnPlayer.Deck), state.NonTurnPlayer.Hands...); err != nil {
			return err
		}

		return s.GameStateRepository.Store(ctx, state)
	})
	if err != nil {
		return err
	}

	return nil
}
