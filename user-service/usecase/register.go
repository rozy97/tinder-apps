package usecase

import (
	"context"

	"github.com/rozy97/tinder-apps/user-service/config"
	"github.com/rozy97/tinder-apps/user-service/model"
	"github.com/rozy97/tinder-apps/user-service/request"
	"github.com/rozy97/tinder-apps/user-service/response"
)

func (u *Usecase) Register(ctx context.Context, req request.Register) (*response.Register, error) {
	count, err := u.repository.CountUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return &response.Register{
			Code:    response.ErrCodeEmailAlreadyRegistered,
			Message: response.ErrMessageEmailAlreadyRegistered,
		}, nil
	}

	hashPassword := config.GetMD5Hash(req.Password + u.env.SecretPassword)

	user := model.User{
		Password: hashPassword,
		Email:    req.Email,
	}

	err = u.repository.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &response.Register{Code: 0, Message: "success"}, nil
}
