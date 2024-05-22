package response

import "github.com/rozy97/tinder-apps/user-service/config"

const (
	ErrCodeEmailAlreadyRegistered = 123

	ErrMessageEmailAlreadyRegistered = "email already registered"
	ErrMessageIncorrectPassword      = "incorrect password, please try again"
)

type Register struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Login struct {
	UserID     uint64        `json:"uint64"`
	Status     config.Status `json:"status"`
	IsVerified uint8         `json:"is_verified"`
}

type Profile struct {
	Email      string        `json:"email"`
	FullName   string        `json:"full_name" `
	Gender     config.Gender `json:"gender"`
	ProfileURL string        `json:"profile_url"`
	Status     config.Status `json:"status"`
	Isverified uint8         `json:"is_verified"`
}
