package repository

import (
	"context"
	"database/sql"

	"github.com/kimnopal/ci-lab-go/internal/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, name, email string) (model.User, error) {
	var u model.User
	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO users(name, email) VALUES ($1, $2)
         RETURNING id, name, email`,
		name, email,
	).Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

func (r *Repository) List(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, email FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}
