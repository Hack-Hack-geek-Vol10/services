package infra

import (
	"context"
	"database/sql"

	"github.com/schema-creator/services/save-service/internal/domain"
)

type collabRepo struct {
	db *sql.DB
}

type CollabRepo interface {
	Create(ctx context.Context, arg domain.CreateEditorParam) error
	Get(ctx context.Context, arg domain.GetEditorParam) (*domain.Token, error)
	Delete(ctx context.Context, arg domain.DeleteEditorParam) error
}

func NewCollabRepo(db *sql.DB) CollabRepo {
	return &collabRepo{db: db}
}

func (c *collabRepo) Create(ctx context.Context, arg domain.CreateEditorParam) error {
	const query = `INSERT INTO editor (project_id, Query) VALUES ($1,$2)`

	row := c.db.QueryRowContext(ctx, query, arg.ProjectID, arg.Query)

	return row.Err()
}

func (c *collabRepo) Get(ctx context.Context, arg domain.GetTokenParam) (*domain.Token, error) {
	const query = `SELECT project_id, FROM tokens WHERE token_id = $1`
	row := c.db.QueryRowContext(ctx, query, arg.TokenID)
	var token domain.Token
	if err := row.Scan(&token.TokenID, &token.ProjectID, &token.Authority); err != nil {
		return nil, err
	}
	return &token, nil
}

func (c *collabRepo) Delete(ctx context.Context, arg domain.DeleteTokenParam) error {
	const query = `DELETE FROM tokens WHERE project_id = $1`
	row, err := c.db.ExecContext(ctx, query, arg.ProjectID)
	if err != nil {
		return err
	}
	count, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}