package usecase

import (
	"context"

	"github.com/rozy97/tinder-apps/user-service/config"
	"github.com/rozy97/tinder-apps/user-service/model"
	"github.com/rozy97/tinder-apps/user-service/request"
)

func (u *Usecase) Activation(ctx context.Context, userID uint64, profile request.Profile) error {
	return u.repository.UpdateUserByID(ctx, model.User{
		ID:         userID,
		FullName:   profile.FullName,
		Gender:     profile.Gender,
		ProfileURL: profile.ProfileURL,
		Status:     config.StatusActive,
		CityID:     profile.CityID,
	})
}
