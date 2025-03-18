package usecase

import (
	"context"
	"slices"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

type GetMyPresentsInputPort interface {
	Execute(ctx context.Context, userID string) ([]*model.Present, error)
}

type GetMyPresentsInteractor struct {
	presentRepository repository.PresentRepository
}

func NewGetMyPresentsUsecase(pr repository.PresentRepository) GetMyPresentsInputPort {
	return &GetMyPresentsInteractor{
		presentRepository: pr,
	}
}

func (g *GetMyPresentsInteractor) Execute(ctx context.Context, userID string) ([]*model.Present, error) {
	uid, err := model.ParseUserID(userID)
	if err != nil {
		return nil, err
	}

	presents, err := g.presentRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(presents, func(i *model.Present, j *model.Present) int {
		return i.Time - j.Time
	})

	resp := make([]*model.Present, 0, len(presents))
	for _, p := range presents {
		// 未受け取りのプレゼントのみを取得
		if !slices.Contains(p.ReceivedUserIDs, uid) {
			resp = append(resp, p)
		}
	}
	return resp, nil
}
