package usecase

import (
	"context"
	"slices"

	"github.com/samber/lo"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
	"github.com/yamato0211/brachio-backend/internal/domain/service"
	"github.com/yamato0211/brachio-backend/internal/handler/schema/messages"
	"golang.org/x/xerrors"
)

type SupplyEnergyInputPort interface {
	Execute(ctx context.Context, input *SupplyEnergyInput) error
}

type SupplyEnergyInput struct {
	RoomID    string
	UserID    string
	Positions [][]model.MonsterType
}

type SupplyEnergyInteractor struct {
	GameStateRepository repository.GameStateRepository
	GameEventSender     service.GameEventSender
}

func NewApplyEnergyUsecase(gameStateRepository repository.GameStateRepository) SupplyEnergyInputPort {
	return &SupplyEnergyInteractor{
		GameStateRepository: gameStateRepository,
	}
}

func (i *SupplyEnergyInteractor) Execute(ctx context.Context, input *SupplyEnergyInput) error {
	roomID, err := model.ParseRoomID(input.RoomID)
	if err != nil {
		return err
	}

	userID, err := model.ParseUserID(input.UserID)
	if err != nil {
		return err
	}

	var OppoID model.UserID

	var eventsForMe []*messages.EffectWithSecret
	var eventsForOppo []*messages.Effect
	err = i.GameStateRepository.Transaction(ctx, roomID, func(ctx context.Context) error {
		state, err := i.GameStateRepository.Find(ctx, roomID)
		if err != nil {
			return err
		}

		if !state.IsMyTurn(userID) {
			return xerrors.Errorf("not your turn")
		}

		OppoID = state.NonTurnPlayer.UserID

		for position, energies := range input.Positions {
			monster, err := state.TurnPlayer.GetMonsterByPosition(position)
			if err != nil {
				return err
			}

			for _, energy := range energies {
				_, idx, isFound := lo.FindIndexOf(state.TurnPlayer.CurrentEnergies, func(e model.MonsterType) bool {
					return e == energy
				})
				if !isFound {
					return xerrors.Errorf("energy not found: %s", energy)
				}

				state.TurnPlayer.CurrentEnergies = slices.Delete(state.TurnPlayer.CurrentEnergies, idx, idx+1)
			}

			eventsForMe = append(eventsForMe, i.makeEventForMe(position, energies))
			eventsForOppo = append(eventsForOppo, i.makeEventForOppo(position, energies))

			monster.Energies = append(monster.Energies, energies...)
		}

		return i.GameStateRepository.Store(ctx, state)
	})
	if err != nil {
		return err
	}

	if err := i.GameEventSender.SendDrawEffectEventToActor(ctx, userID, eventsForMe...); err != nil {
		return err
	}
	if err := i.GameEventSender.SendDrawEffectEventToRecipient(ctx, OppoID, eventsForOppo...); err != nil {
		return err
	}

	return nil
}

func (i *SupplyEnergyInteractor) makeEventForMe(position int, energies []model.MonsterType) *messages.EffectWithSecret {
	return &messages.EffectWithSecret{
		Effect: &messages.EffectWithSecret_AttachEnergy{
			AttachEnergy: &messages.AttachEnergyEffect{
				Position: int32(position),
				Energies: lo.Map(energies, func(e model.MonsterType, _ int) messages.Element {
					switch e {
					case model.MonsterTypeAlchohol:
						return messages.Element_ALCHOHOL
					case model.MonsterTypeKnowledge:
						return messages.Element_KNOWLEDGE
					case model.MonsterTypeMoney:
						return messages.Element_MONEY
					case model.MonsterTypeMuscle:
						return messages.Element_MUSCLE
					case model.MonsterTypePopularity:
						return messages.Element_POPULARITY
					default:
						return messages.Element_NULL
					}
				}),
			},
		},
	}
}

func (i *SupplyEnergyInteractor) makeEventForOppo(position int, energies []model.MonsterType) *messages.Effect {
	return &messages.Effect{
		Effect: &messages.Effect_AttachEnergy{
			AttachEnergy: &messages.AttachEnergyEffect{
				Position: int32(position),
				Energies: lo.Map(energies, func(e model.MonsterType, _ int) messages.Element {
					switch e {
					case model.MonsterTypeAlchohol:
						return messages.Element_ALCHOHOL
					case model.MonsterTypeKnowledge:
						return messages.Element_KNOWLEDGE
					case model.MonsterTypeMoney:
						return messages.Element_MONEY
					case model.MonsterTypeMuscle:
						return messages.Element_MUSCLE
					case model.MonsterTypePopularity:
						return messages.Element_POPULARITY
					default:
						return messages.Element_NULL
					}
				}),
			},
		},
	}
}
