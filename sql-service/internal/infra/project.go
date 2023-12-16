package infra

import (
	"context"
	"database/sql"

	"github.com/newrelic/go-agent/v3/newrelic"
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

const (
	DefaultImage = "https://firebasestorage.googleapis.com/v0/b/geek-vol10.appspot.com/o/whiteout.png?alt=media&token=604eedd5-1005-4234-abfd-cb76c594ec28"
)

func (r *projectRepo) Create(ctx context.Context, param domain.CreateProjectParam) (*domain.Project, error) {
	defer newrelic.FromContext(ctx).StartSegment("projectRepo-Create").End()
	const query = `INSERT INTO projects (project_id,title,last_image) VALUES ($1,$2,$3) RETURNING project_id,title,last_image`
	row := r.db.QueryRowContext(ctx, query, param.ProjectID, param.Title, DefaultImage)
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil

}

func (r *projectRepo) ReadOne(ctx context.Context, projectID string) (*domain.Project, error) {
	defer newrelic.FromContext(ctx).StartSegment("projectRepo-ReadOne").End()
	const query = `SELECT project_id,title,last_image FROM projects WHERE project_id = $1 AND is_delete = false`
	row := r.db.QueryRowContext(ctx, query, projectID)
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *projectRepo) ReadAll(ctx context.Context, arg domain.ReadProjectsParam) ([]*domain.Project, error) {
	defer newrelic.FromContext(ctx).StartSegment("projectRepo-ReadAll").End()
	const query = `SELECT projects.project_id, projects.title, projects.last_image FROM projects LEFT OUTER JOIN project_members ON projects.project_id = project_members.project_id WHERE user_id = $1 AND is_delete = false ORDER BY created_at DESC LIMIT $2 OFFSET $3`
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
	defer newrelic.FromContext(ctx).StartSegment("projectRepo-UpdateTitle").End()
	const query = `UPDATE projects SET title = $1 WHERE project_id = $2 RETURNING project_id,title,last_image`
	row := r.db.QueryRowContext(ctx, query, title, projectID)
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *projectRepo) UpdateLastImage(ctx context.Context, projectID, lastImage string) (*domain.Project, error) {
	defer newrelic.FromContext(ctx).StartSegment("projectRepo-UpdateLastImage").End()
	const query = `UPDATE projects SET last_image = $1 WHERE project_id = $2 RETURNING project_id,title,last_image`
	row := r.db.QueryRowContext(ctx, query, lastImage, projectID)
	var info domain.Project
	if err := row.Scan(&info.ProjectID, &info.Title, &info.LastImage); err != nil {
		return nil, err
	}
	return &info, nil
}

func (r *projectRepo) Delete(ctx context.Context, projectID string) error {
	defer newrelic.FromContext(ctx).StartSegment("projectRepo-Delete").End()
	const query = `UPDATE projects SET is_delete = true WHERE project_id = $1`
	_, err := r.db.ExecContext(ctx, query, projectID)
	return err
}