package db

import (
	"context"

	"github.com/guregu/dynamo/v2"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/domain/repository"
)

const (
	userTable   = "Users"
	userHashKey = "UserId"
)

type userRepository struct {
	db *dynamo.DB
}

func (u *userRepository) Find(ctx context.Context, userID model.UserID) (*model.User, error) {
	var data model.User
	if err := u.db.Table(userTable).Get(userHashKey, userID).One(ctx, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *userRepository) Store(ctx context.Context, user *model.User) error {
	return u.db.Table(userTable).Put(user).Run(ctx)
}

func NewUserRepository(db *dynamo.DB) repository.UserRepository {
	return &userRepository{db: db}
}
