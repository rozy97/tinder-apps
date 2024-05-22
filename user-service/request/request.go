package request

import "github.com/rozy97/tinder-apps/user-service/config"

type Register struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"  binding:"required,min=3"`
}

type Login struct {
	Email    string `json:"email"  binding:"required,email"`
	Password string `json:"password"  binding:"required,min=3"`
}

type Profile struct {
	FullName   string        `json:"full_name" binding:"required"`
	Gender     config.Gender `json:"gender" binding:"required"`
	ProfileURL string        `json:"profile_url" binding:"required,url"`
	CityID     uint64        `json:"city_id" binding:"required"`
}
