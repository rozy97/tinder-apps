package usecase

import (
	"context"
	"fmt"

	"github.com/rozy97/tinder-apps/user-service/config"
	"github.com/rozy97/tinder-apps/user-service/request"
	"github.com/rozy97/tinder-apps/user-service/response"
)

func (u *Usecase) Login(ctx context.Context, req request.Login) (*response.Login, error) {
	user, err := u.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if config.GetMD5Hash(req.Password+u.env.SecretPassword) != user.Password {
		return nil, fmt.Errorf(response.ErrMessageIncorrectPassword)
	}

	return &response.Login{
		UserID:     user.ID,
		Status:     user.Status,
		IsVerified: user.IsVerified,
	}, nil
}
