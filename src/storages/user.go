package storage

import (
	"database/sql"

	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
)

type userRepo struct {
	db *sql.DB
}

type UserRepo interface {
	Create(arg domain.CreateUserParams) error
	ReadOne(userID string) (*domain.User, error)
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(arg domain.CreateUserParams) error {
	const query = `INSERT INTO users (user_id, name, email)VALUES($1, $2, $3)`
	row := r.db.QueryRow(query, arg.UserID, arg.Name, arg.Email)
	return row.Err()
}

func (r *userRepo) ReadOne(userID string) (*domain.User, error) {
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
