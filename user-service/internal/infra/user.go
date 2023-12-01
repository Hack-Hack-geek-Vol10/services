package infra

import (
	"context"
	"database/sql"

	"github.com/schema-creator/services/user-service/internal/domain"
)

type userRepo struct {
	db *sql.DB
}

type UserRepo interface {
	Create(ctx context.Context, arg domain.CreateUserParams) (*domain.User, error)
	ReadOne(ctx context.Context, userID string) (*domain.User, error)
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, arg domain.CreateUserParams) (*domain.User, error) {
	const query = `INSERT INTO users (user_id, name, email)VALUES($1, $2, $3) RETURNING user_id, name, email`

	row := r.db.QueryRowContext(ctx, query, arg.UserID, arg.Name, arg.Email)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var user domain.User
	if err := row.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) ReadOne(ctx context.Context, userID string) (*domain.User, error) {
	const query = `SELECT user_id, name, email FROM users WHERE user_id = $1 AND is_delete = false`

	row := r.db.QueryRow(query, userID)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var user domain.User
	if err := row.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
