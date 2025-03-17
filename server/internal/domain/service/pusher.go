package service

import (
	"context"

	"github.com/yamato0211/brachio-backend/internal/domain/model"
)

type GameEventSender interface {
	SendXXX(ctx context.Context, userID model.UserID, event any) error
}
