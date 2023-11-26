package storages

import (
	"context"
	"database/sql"

	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
)

type projectRepo struct {
	db *sql.DB
}

type ProjectRepo interface {
	Create(context.Context, domain.CreateProjectParam) error
}

func NewProjectRepo(db *sql.DB) ProjectRepo {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(ctx context.Context, param domain.CreateProjectParam) error {
	const query = `INSERT INTO projects (project_id,title,last_image) VALUES ($1,$2,$3) `
	row := r.db.QueryRowContext(ctx, query, param.ProjectID, param.Title, "")
	return row.Err()
}
