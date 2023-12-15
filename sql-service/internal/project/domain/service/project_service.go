package service

import (
	"context"

	"github.com/schema-creator/services/sql-service/internal/project/domain/response"
	"github.com/schema-creator/services/sql-service/internal/project/domain/values"
	"github.com/schema-creator/services/sql-service/pkg/ddl"
)

type ProjectService interface {
	CreateProject(ctx context.Context, args values.CreateProject) (*values.Project, error)
	GetProjectsByUserID(ctx context.Context, userID string) ([]*values.Project, error)
	DeleteAllByUserID(ctx context.Context, userID string) error
	GetOneByID(ctx context.Context, id string) (*values.Project, error)
	UpdateProject(ctx context.Context, args values.Project) (*values.Project, error)
	Delete(ctx context.Context, id string) error
	JoinProject(ctx context.Context, userID string, projectID string) (*values.Project, error)
	ConvertDDL(ctx context.Context, id string, convertType ddl.ConvertType) (*response.ConvertDDL, error)
}
