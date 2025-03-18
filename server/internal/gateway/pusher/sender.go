package websocket

import (
	"context"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	wsmsg "github.com/yamato0211/brachio-backend/internal/handler/schema/websocket"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/websocket/event"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/websocket/payload"
	"github.com/yamato0211/brachio-backend/pkg/websocket"
	"google.golang.org/protobuf/proto"
)

type gameEventSender struct {
	pusher websocket.Pusher
}

func NewGameEventSender(pusher websocket.Pusher) service.GameEventSender {
	return &gameEventSender{
		pusher: pusher,
	}
}

func (g *gameEventSender) SendMatchingComplete(ctx context.Context, userID model.UserID, users []*model.User, roomID string) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_MatchingCompleteEventToActor{
			MatchingCompleteEventToActor: &event.MatchingCompleteEventToActor{
				Payload: &payload.MatchingCompleteEventPayload{
					BattleId: roomID,
					Users:    make([]*messages.User, 0, len(users)),
				},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, userID, b)
}

func (g *gameEventSender) SendAbillityToActor(ctx context.Context, userID model.UserID, card *model.Card) error {
	event := wsmsg.EventEnvelope{
		Event: &wsmsg.EventEnvelope_AbilityEventToActor{
			AbilityEventToActor: &event.AbilityEventToActor{
				Payload: &payload.AbilityPayload{
					Card: &messages.Card{
						Id: card.CardID.String(),
						MasterCard: &messages.MasterCard{
							CardVariant: &messages.MasterCard_MasterMonsterCard{
								MasterMonsterCard: &messages.MasterMonsterCard{
									Base:     &messages.MasterCardBase{},
									Hp:       0,
									Element:  g.element(card.MasterCard.Type),
									Weakness: g.element(card.MasterCard.Weakness),
								},
							},
						},
					},
					Ability: &messages.Ability{
						Name: "",
						Text: "",
					},
				},
			},
		},
	}

	b, err := proto.Marshal(&event)
	if err != nil {
		return err
	}

	return g.pusher.Send(ctx, userID, b)
}

// SendAbillityToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendAbillityToRecipient(ctx context.Context, userID model.UserID, abilityID int) error {
	panic("unimplemented")
}

// SendAttackMonsterToActor implements service.GameEventSender.
func (g *gameEventSender) SendAttackMonsterToActor(ctx context.Context, userID model.UserID, attacker *model.Monster, skillID int) error {
	panic("unimplemented")
}

// SendAttackMonsterToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendAttackMonsterToRecipient(ctx context.Context, userID model.UserID, recipient *model.Monster, damage int) error {
	panic("unimplemented")
}

// SendCoinTossToActor implements service.GameEventSender.
func (g *gameEventSender) SendCoinTossToActor(ctx context.Context, userID model.UserID, result bool) error {
	panic("unimplemented")
}

// SendCoinTossToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendCoinTossToRecipient(ctx context.Context, userID model.UserID, result bool) error {
	panic("unimplemented")
}

// SendConfirmActionToActor implements service.GameEventSender.
func (g *gameEventSender) SendConfirmActionToActor(ctx context.Context, userID model.UserID) error {
	panic("unimplemented")
}

// SendConfirmEnergyToActor implements service.GameEventSender.
func (g *gameEventSender) SendConfirmEnergyToActor(ctx context.Context, userID model.UserID, energy int) error {
	panic("unimplemented")
}

// SendConfirmEnergyToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendConfirmEnergyToRecipient(ctx context.Context, userID model.UserID, energy int) error {
	panic("unimplemented")
}

// SendConfirmTargetToActor implements service.GameEventSender.
func (g *gameEventSender) SendConfirmTargetToActor(ctx context.Context, userID model.UserID) error {
	panic("unimplemented")
}

// SendConfirmTargetToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendConfirmTargetToRecipient(ctx context.Context, userID model.UserID) error {
	panic("unimplemented")
}

// SendDecideOrderToActor implements service.GameEventSender.
func (g *gameEventSender) SendDecideOrderToActor(ctx context.Context, userID model.UserID, order int) error {
	panic("unimplemented")
}

// SendDecideOrderToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendDecideOrderToRecipient(ctx context.Context, userID model.UserID, order int) error {
	panic("unimplemented")
}

// SendDrawCardToActor implements service.GameEventSender.
func (g *gameEventSender) SendDrawCardToActor(ctx context.Context, userID model.UserID, card *model.Card) error {
	panic("unimplemented")
}

// SendDrawCardToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendDrawCardToRecipient(ctx context.Context, userID model.UserID, card *model.Card) error {
	panic("unimplemented")
}

// SendEvolutionMonsterToActor implements service.GameEventSender.
func (g *gameEventSender) SendEvolutionMonsterToActor(ctx context.Context, userID model.UserID, monster *model.Monster) error {
	panic("unimplemented")
}

// SendEvolutionMonsterToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendEvolutionMonsterToRecipient(ctx context.Context, userID model.UserID, monster *model.Monster) error {
	panic("unimplemented")
}

// SendExchangeDeckToActor implements service.GameEventSender.
func (g *gameEventSender) SendExchangeDeckToActor(ctx context.Context, userID model.UserID, deck []*model.Card) error {
	panic("unimplemented")
}

// SendExchangeDeckToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendExchangeDeckToRecipient(ctx context.Context, userID model.UserID, deck []*model.Card) error {
	panic("unimplemented")
}

// SendGiveUpToActor implements service.GameEventSender.
func (g *gameEventSender) SendGiveUpToActor(ctx context.Context, userID model.UserID) error {
	panic("unimplemented")
}

// SendGiveUpToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendGiveUpToRecipient(ctx context.Context, userID model.UserID) error {
	panic("unimplemented")
}

// SendNextEnergyToActor implements service.GameEventSender.
func (g *gameEventSender) SendNextEnergyToActor(ctx context.Context, userID model.UserID, energy int) error {
	panic("unimplemented")
}

// SendStartGame implements service.GameEventSender.
func (g *gameEventSender) SendStartGame(ctx context.Context, userID model.UserID, users []*model.User, roomID string) error {
	panic("unimplemented")
}

// SendSummonMonsterToActor implements service.GameEventSender.
func (g *gameEventSender) SendSummonMonsterToActor(ctx context.Context, userID model.UserID, monster *model.Monster, position int) error {
	panic("unimplemented")
}

// SendSummonMonsterToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendSummonMonsterToRecipient(ctx context.Context, userID model.UserID, monster *model.Monster, position int) error {
	panic("unimplemented")
}

// SendSupplyEnergyToActor implements service.GameEventSender.
func (g *gameEventSender) SendSupplyEnergyToActor(ctx context.Context, userID model.UserID, energy int) error {
	panic("unimplemented")
}

// SendSupplyEnergyToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendSupplyEnergyToRecipient(ctx context.Context, userID model.UserID, energy int) error {
	panic("unimplemented")
}

// SendTurnEnd implements service.GameEventSender.
func (g *gameEventSender) SendTurnEnd(ctx context.Context, userID model.UserID) error {
	panic("unimplemented")
}

// SendTurnStart implements service.GameEventSender.
func (g *gameEventSender) SendTurnStart(ctx context.Context, userID model.UserID) error {
	panic("unimplemented")
}

// SendUseGoodsToActor implements service.GameEventSender.
func (g *gameEventSender) SendUseGoodsToActor(ctx context.Context, userID model.UserID, goodsID int) error {
	panic("unimplemented")
}

// SendUseGoodsToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendUseGoodsToRecipient(ctx context.Context, userID model.UserID, goodsID int) error {
	panic("unimplemented")
}

// SendUseSupportToActor implements service.GameEventSender.
func (g *gameEventSender) SendUseSupportToActor(ctx context.Context, userID model.UserID, skillID int) error {
	panic("unimplemented")
}

// SendUseSupportToRecipient implements service.GameEventSender.
func (g *gameEventSender) SendUseSupportToRecipient(ctx context.Context, userID model.UserID, skillID int) error {
	panic("unimplemented")
}

func (g *gameEventSender) element(e model.MonsterType) messages.Element {
	switch e {
	case model.MonsterTypeAlchohol:
		return messages.Element_ALCHOHOL
	case model.MonsterTypeMoney:
		return messages.Element_MONEY
	case model.MonsterTypeKnowledge:
		return messages.Element_KNOWLEDGE
	case model.MonsterTypePopularity:
		return messages.Element_POPULARITY
	case model.MonsterTypeMuscle:
		return messages.Element_MUSCLE
	case model.MonsterTypeNull:
		return messages.Element_NULL
	default:
		return messages.Element_ELEMENT_UNKNOWN
	}
}

func (g *gameEventSender) monsterCard(card *model.Card) *messages.MasterCard {
	ret := &messages.MasterCard{
		CardVariant: &messages.MasterCard_MasterMonsterCard{
			MasterMonsterCard: &messages.MasterMonsterCard{
				Base:     &messages.MasterCardBase{},
				Hp:       0,
				Element:  g.element(card.MasterCard.Type),
				Weakness: g.element(card.MasterCard.Weakness),
				Skills: lo.Map(card.MasterCard.Skills, func(s *model.Skill, _ int) *messages.Skill {
					return &messages.Skill{
						Name:   s.Name,
						Text:   s.Text,
						Damage: int32(s.Damage),
						DamageOption: func() *messages.DamageOption {
							switch s.DamageOption {
							case "x":
								return lo.ToPtr(messages.DamageOption_X)
							case "+":
								return lo.ToPtr(messages.DamageOption_PLUS)
							default:
								return nil
							}
						}(),
						Cost: lo.Map(s.Cost, func(e model.MonsterType, _ int) messages.Element { return g.element(e) }),
					}
				}),
				Ability: lo.Ternary(card.MasterCard.Ability != nil, &messages.Ability{
					Name: card.MasterCard.Ability.Name,
					Text: card.MasterCard.Ability.Text,
				}, nil),
				RetreatCost: int32(card.MasterCard.RetreatCost),
				EvolvesFrom: lo.Map(card.MasterCard.EvolvesFrom, func(id model.MasterCardID, _ int) string { return id.String() }),
				EvolvesTo:   lo.Map(card.MasterCard.EvolvesTo, func(id model.MasterCardID, _ int) string { return id.String() }),
			},
		},
	}
	switch card.MasterCard.CardType {
	case model.CardTypeMonster:
		ret.CardVariant = &messages.MasterCard_MasterMonsterCard{
			MasterMonsterCard: &messages.MasterMonsterCard{
				Base:     &messages.MasterCardBase{},
				Hp:       0,
				Element:  g.element(card.MasterCard.Type),
				Weakness: g.element(card.MasterCard.Weakness),
			},
		}
	case model.CardTypeSupporter:
		ret.CardVariant = &messages.MasterCard_MasterSupporterCard{
			MasterSupporterCard: &messages.MasterSupporterCard{
				Base: &messages.MasterCardBase{},
			},
		}
	case model.CardTypeGoods:
		ret.CardVariant = &messages.MasterCard_MasterGoodsCard{
			MasterGoodsCard: &messages.MasterGoodsCard{
				Base: &messages.MasterCardBase{},
			},
		}
	}

	return ret
}
