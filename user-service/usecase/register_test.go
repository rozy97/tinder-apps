package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/rozy97/tinder-apps/user-service/config"
	mocks "github.com/rozy97/tinder-apps/user-service/mocks/usecase"
	"github.com/rozy97/tinder-apps/user-service/model"
	"github.com/rozy97/tinder-apps/user-service/request"
	"github.com/rozy97/tinder-apps/user-service/response"
	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	mockRepository := mocks.NewRepository(t)
	ctx := context.Background()
	email := "rozy.firmansyah@gmail.com"
	password := "12345"
	request := request.Register{
		Email:    email,
		Password: password,
	}

	env := config.Environment{
		SecretPassword: "secret",
	}

	t.Run("failed test case, error while count email in database", func(t *testing.T) {
		usecase := NewUsecase(mockRepository, env)
		expectedErrorMessage := "error while execute query"
		mockRepository.EXPECT().CountUserByEmail(ctx, email).Return(0, fmt.Errorf(expectedErrorMessage)).Once()
		result, err := usecase.Register(ctx, request)

		assert.Nil(t, result)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})

	t.Run("failed test case, error email already registered", func(t *testing.T) {
		usecase := NewUsecase(mockRepository, env)
		mockRepository.EXPECT().CountUserByEmail(ctx, email).Return(1, nil).Once()
		result, err := usecase.Register(ctx, request)

		assert.Equal(t, response.ErrCodeEmailAlreadyRegistered, result.Code)
		assert.Equal(t, response.ErrMessageEmailAlreadyRegistered, result.Message)
		assert.Equal(t, nil, err)
	})

	t.Run("failed test case, error while insert user", func(t *testing.T) {
		usecase := NewUsecase(mockRepository, env)
		mockRepository.EXPECT().CountUserByEmail(ctx, email).Return(0, nil).Once()
		expectedErrorMessage := "error while execute query"
		mockRepository.EXPECT().InsertUser(ctx, model.User{
			Password: config.GetMD5Hash(password + env.SecretPassword),
			Email:    email,
		}).Return(fmt.Errorf(expectedErrorMessage)).Once()
		result, err := usecase.Register(ctx, request)

		assert.Nil(t, result)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})

	t.Run("success test case", func(t *testing.T) {
		usecase := NewUsecase(mockRepository, env)
		mockRepository.EXPECT().CountUserByEmail(ctx, email).Return(0, nil).Once()
		mockRepository.EXPECT().InsertUser(ctx, model.User{
			Password: config.GetMD5Hash(password + env.SecretPassword),
			Email:    email,
		}).Return(nil).Once()
		result, err := usecase.Register(ctx, request)

		assert.Equal(t, 0, result.Code)
		assert.Equal(t, "success", result.Message)
		assert.Nil(t, err)
	})

}
