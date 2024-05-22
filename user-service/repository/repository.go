package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rozy97/tinder-apps/user-service/model"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CountUserByEmail(ctx context.Context, email string) (total uint, err error) {
	query := `SELECT COUNT(id) AS total FROM users WHERE email = ?`
	err = r.db.GetContext(ctx, &total, query, email)
	return
}

func (r *Repository) InsertUser(ctx context.Context, user model.User) error {
	query := `
		INSERT INTO users(password, email, created_at, updated_at)
		VALUES (?, ?, UTC_TIMESTAMP, UTC_TIMESTAMP)
	`
	_, err := r.db.Exec(query, user.Password, user.Email)
	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	query := `
		SELECT id, password, email,
			COALESCE(full_name, '') AS full_name,
			gender,
			COALESCE(profile_url, '') AS profile_url,
			status,
			COALESCE(city_id, 0) AS city_id,
			COALESCE(is_verified, 0) AS is_verified,
			verified_until, created_at, updated_at, deleted_at
		FROM users
		WHERE email = ?
	`
	err = r.db.GetContext(ctx, &user, query, email)
	return
}

func (r *Repository) UpdateUserByID(ctx context.Context, user model.User) error {
	query := `
		UPDATE users SET full_name = ?, gender = ?, profile_url = ?, status = ?, city_id = ?, updated_at = UTC_TIMESTAMP
		WHERE id = ?	
	`
	_, err := r.db.Exec(query, user.FullName, user.Gender, user.ProfileURL, user.Status, user.CityID, user.ID)
	return err
}
