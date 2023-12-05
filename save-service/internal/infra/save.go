package infra

import (
	"context"
	"database/sql"

	"github.com/schema-creator/services/save-service/internal/domain"
)

type saveRepo struct {
	db *sql.DB
}

type SaveRepo interface {
	Create(ctx context.Context, arg domain.CreateSaveParam) error
	Get(ctx context.Context, arg domain.GetEditorParam) (*domain.Token, error)
	Delete(ctx context.Context, arg domain.DeleteEditorParam) error
}

func NewSaveRepo(db *sql.DB) SaveRepo {
	return &saveRepo{db: db}
}

func (s *saveRepo) Create(ctx context.Context, arg domain.CreateSaveParam) error {
	const query = `INSERT INTO saves (save_id, project_id, editor, object) VALUES ($1,$2,$3,$4)`
	row := s.db.QueryRowContext(ctx, query, arg.SaveID, arg.ProjectID, arg.Editor, arg.Object)

	return row.Err()
}

func (s *saveRepo) Get(ctx context.Context, arg domain.GetTokenParam) (*domain.Token, error) {
	const query = `SELECT project_id, FROM tokens WHERE token_id = $1`
	row := s.db.QueryRowContext(ctx, query, arg.TokenID)
	var token domain.Token
	if err := row.Scan(&token.TokenID, &token.ProjectID, &token.Authority); err != nil {
		return nil, err
	}
	return &token, nil
}

func (s *saveRepo) Delete(ctx context.Context, arg domain.DeleteTokenParam) error {
	const query = `DELETE FROM tokens WHERE project_id = $1`
	row, err := s.db.ExecContext(ctx, query, arg.ProjectID)
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
