package repository

import (
	"context"

	"github.com/schema-creator/services/sql-service/internal/project/domain/values"
)

type ProjectRepository interface {
	Create(context.Context, values.CreateProject) (*values.Project, error)
	GetProjectsByUserID(context.Context, string) ([]*values.Project, error)
	DeleteAllByUserID(context.Context, string) error

	GetOneByID(context.Context, string) (*values.Project, error)
	UpdateProject(context.Context, values.Project) (*values.Project, error)
	Delete(context.Context, string) error

	JoinProject(context.Context, string, string) (*values.Project, error)
	WsUpdateEditor(ctx context.Context, projectID, editor string) (*values.Project, error)
	WsUpdateObject(ctx context.Context, projectID, object string) (*values.Project, error)
}
