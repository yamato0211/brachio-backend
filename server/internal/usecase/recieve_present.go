package usecase

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type ReceivePresentInputPort interface {
	Execute(ctx context.Context, userID string, presentID string) error
}

type ReceivePresentInteractor struct {
	presentRepository repository.PresentRepository
	userRepository    repository.UserRepository
}

func NewReceivePresentUsecase(pr repository.PresentRepository, ur repository.UserRepository) ReceivePresentInputPort {
	return &ReceivePresentInteractor{
		presentRepository: pr,
		userRepository:    ur,
	}
}

func (r *ReceivePresentInteractor) Execute(ctx context.Context, userID string, presentID string) error {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return err
	}

	pid, err := model.ParsePresentID(presentID)
	if err != nil {
		return err
	}

	present, err := r.presentRepository.Find(ctx, pid)
	if err != nil {
		return err
	}

	user, err := r.userRepository.Find(ctx, uid)
	if err != nil {
		return err
	}

	// プレゼントの受け取り処理
	present.ReceivedUserIDs = append(present.ReceivedUserIDs, uid)
	if err := r.presentRepository.Store(ctx, present); err != nil {
		return err
	}

	// ユーザーのアイテムを追加
	user.ItemIDsWithCount[present.Item.ItemID.String()] += present.ItemCount
	if err := r.userRepository.Store(ctx, user); err != nil {
		return err
	}
	return nil
}
