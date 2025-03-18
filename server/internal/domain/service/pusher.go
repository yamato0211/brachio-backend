package service

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type GameEventSender interface {
	SendMatchingComplete(ctx context.Context, userID model.UserID, users []*model.User, roomID string) error
	SendAttackMonsterToActor(ctx context.Context, userID model.UserID, attacker *model.Monster, skillID int) error
	SendAttackMonsterToRecipient(ctx context.Context, userID model.UserID, recipient *model.Monster, damage int) error
	SendSummonMonsterToActor(ctx context.Context, userID model.UserID, monster *model.Monster, position int) error
	SendSummonMonsterToRecipient(ctx context.Context, userID model.UserID, monster *model.Monster, position int) error
	SendEvolutionMonsterToActor(ctx context.Context, userID model.UserID, monster *model.Monster) error
	SendEvolutionMonsterToRecipient(ctx context.Context, userID model.UserID, monster *model.Monster) error
	SendUseSupportToActor(ctx context.Context, userID model.UserID, skillID int) error
	SendUseSupportToRecipient(ctx context.Context, userID model.UserID, skillID int) error
	SendUseGoodsToActor(ctx context.Context, userID model.UserID, goodsID int) error
	SendUseGoodsToRecipient(ctx context.Context, userID model.UserID, goodsID int) error
	SendSupplyEnergyToActor(ctx context.Context, userID model.UserID, energy int) error
	SendSupplyEnergyToRecipient(ctx context.Context, userID model.UserID, energy int) error
	SendGiveUpToActor(ctx context.Context, userID model.UserID) error
	SendGiveUpToRecipient(ctx context.Context, userID model.UserID) error
	SendAbillityToActor(ctx context.Context, userID model.UserID, abilityID int) error
	SendAbillityToRecipient(ctx context.Context, userID model.UserID, abilityID int) error
	SendDrawCardToActor(ctx context.Context, userID model.UserID, card *model.Card) error
	SendDrawCardToRecipient(ctx context.Context, userID model.UserID, card *model.Card) error
	SendConfirmActionToActor(ctx context.Context, userID model.UserID) error
	SendStartGame(ctx context.Context, userID model.UserID, users []*model.User, roomID string) error
	SendTurnStart(ctx context.Context, userID model.UserID) error
	SendTurnEnd(ctx context.Context, userID model.UserID) error
	SendCoinTossToActor(ctx context.Context, userID model.UserID, result bool) error
	SendCoinTossToRecipient(ctx context.Context, userID model.UserID, result bool) error
	SendConfirmEnergyToActor(ctx context.Context, userID model.UserID, energy int) error
	SendConfirmEnergyToRecipient(ctx context.Context, userID model.UserID, energy int) error
	SendConfirmTargetToActor(ctx context.Context, userID model.UserID) error
	SendConfirmTargetToRecipient(ctx context.Context, userID model.UserID) error
	SendNextEnergyToActor(ctx context.Context, userID model.UserID, energy int) error
	SendExchangeDeckToActor(ctx context.Context, userID model.UserID, deck []*model.Card) error
	SendExchangeDeckToRecipient(ctx context.Context, userID model.UserID, deck []*model.Card) error
	SendDecideOrderToActor(ctx context.Context, userID model.UserID, order int) error
	SendDecideOrderToRecipient(ctx context.Context, userID model.UserID, order int) error
}
