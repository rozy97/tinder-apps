package usecase

import (
	"context"

	"github.com/rozy97/tinder-apps/user-service/config"
	"github.com/rozy97/tinder-apps/user-service/model"
)

type Usecase struct {
	repository Repository
	env        config.Environment
}

func NewUsecase(repository Repository, env config.Environment) *Usecase {
	return &Usecase{
		repository: repository,
		env:        env,
	}
}

type Repository interface {
	CountUserByEmail(ctx context.Context, email string) (total uint, err error)
	InsertUser(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	UpdateUserByID(ctx context.Context, user model.User) error
}
