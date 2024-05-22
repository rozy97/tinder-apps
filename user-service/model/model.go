package model

import (
	"time"

	"github.com/rozy97/tinder-apps/user-service/config"
)

type User struct {
	ID            uint64        `db:"id"`
	Password      string        `db:"password"`
	Email         string        `db:"email"`
	FullName      string        `db:"full_name"`
	Gender        config.Gender `db:"gender"`
	ProfileURL    string        `db:"profile_url"`
	Status        config.Status `db:"status"`
	CityID        uint64        `db:"city_id"`
	IsVerified    uint8         `db:"is_verified"`
	VerifiedUntil *time.Time    `db:"verified_until"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
	DeletedAt     *time.Time    `db:"deleted_at"`
}
