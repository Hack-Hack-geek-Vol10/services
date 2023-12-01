package infra

import (
	"context"
	"database/sql"

	"github.com/Hack-Hack-geek-Vol10/services/project-service/internal/domain"
	"github.com/google/uuid"
)

type tokenRepo struct {
	db *sql.DB
}

type TokenRepo interface {
	Create(ctx context.Context, arg domain.CreateTokenParam) error
	Get(ctx context.Context, arg domain.GetTokenParam) (*domain.Token, error)
	Delete(ctx context.Context, arg domain.DeleteTokenParam) error
}

func NewTokenRepo(db *sql.DB) TokenRepo {
	return &tokenRepo{db: db}
}

func (t *tokenRepo) Create(ctx context.Context, arg domain.CreateTokenParam) error {
	const query = `INSERT INTO tokens (token_id, project_id, authority) VALUES ($1,$2,$3)` // <- ここに返すカラムを書く
	uuid := uuid.New().String()                                                            // <- string型のuuidを生成したほうがいいかも uuid.New().String()

	row := t.db.QueryRowContext(ctx, query, uuid, arg.ProjectID, arg.Authority)

	return row.Err()
}

func (t *tokenRepo) Get(ctx context.Context, arg domain.GetTokenParam) (*domain.Token, error) {
	const query = `SELECT token_id,project_id,authority FROM tokens WHERE token = $1`
	row := t.db.QueryRowContext(ctx, query, arg.TokenID)
	var token domain.Token
	if err := row.Scan(&token.TokenID, &token.ProjectID, &token.Authority); err != nil {
		return nil, err
	}
	return &token, nil
}

func (t *tokenRepo) Delete(ctx context.Context, arg domain.DeleteTokenParam) error {
	const query = `DELETE FROM tokens WHERE project_id = $1`
	row := t.db.QueryRowContext(ctx, query, arg.ProjectID)

	return row.Err()
}
