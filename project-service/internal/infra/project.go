package infra

import (
	"context"
	"database/sql"

	"github.com/schema-creator/services/project-service/internal/domain"
)

type projectRepo struct {
	db *sql.DB
}

type ProjectRepo interface {
	Create(context.Context, domain.CreateProjectParam) (*domain.Project, error)
	ReadOne(ctx context.Context, projectID string) (*domain.Project, error)
	ReadAll(ctx context.Context, arg domain.ReadProjectsParam) ([]*domain.Project, error)
	UpdateTitle(ctx context.Context, projectID, title string) (*domain.Project, error)
	UpdateLastImage(ctx context.Context, projectID, lastImage string) (*domain.Project, error)
	Delete(ctx context.Context, projectID string) error
}

func NewProjectRepo(db *sql.DB) ProjectRepo {
	return &projectRepo{db: db}
}

func (r *projectRepo) Create(ctx context.Context, param domain.CreateProjectParam) (*domain.Project, error) {
	const query = `INSERT INTO projects (project_id,title,last_image) VALUES ($1,$2,$3) RETURNING project_id,title,last_image`
	row := r.db.QueryRowContext(ctx, query, param.ProjectID, param.Title, "")
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil

}

func (r *projectRepo) ReadOne(ctx context.Context, projectID string) (*domain.Project, error) {
	const query = `SELECT project_id,title,last_image FROM projects WHERE project_id = $1`
	row := r.db.QueryRowContext(ctx, query, projectID)
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *projectRepo) ReadAll(ctx context.Context, arg domain.ReadProjectsParam) ([]*domain.Project, error) {
	const query = `SELECT project_id,title,last_image FROM projects WHERE project_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var infos []*domain.Project
	for rows.Next() {
		var info domain.Project
		if err := rows.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
			return nil, err
		}
		infos = append(infos, &info)
	}
	return infos, nil
}

func (r *projectRepo) UpdateTitle(ctx context.Context, projectID, title string) (*domain.Project, error) {
	const query = `UPDATE projects SET title = $1 WHERE project_id = $2 RETURNING project_id,title,last_image`
	row := r.db.QueryRowContext(ctx, query, title, projectID)
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *projectRepo) UpdateLastImage(ctx context.Context, projectID, lastImage string) (*domain.Project, error) {
	const query = `UPDATE projects SET last_image = $1 WHERE project_id = $2 RETURNING project_id,title,last_image`
	row := r.db.QueryRowContext(ctx, query, lastImage, projectID)
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *projectRepo) Delete(ctx context.Context, projectID string) error {
	const query = `DELETE FROM projects WHERE project_id = $1`
	_, err := r.db.ExecContext(ctx, query, projectID)
	return err
}
