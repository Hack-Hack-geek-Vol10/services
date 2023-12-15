package services

import (
	"context"

	"github.com/schema-creator/services/sql-service/internal/project/domain/repository"
	"github.com/schema-creator/services/sql-service/internal/project/domain/response"
	"github.com/schema-creator/services/sql-service/internal/project/domain/service"
	"github.com/schema-creator/services/sql-service/internal/project/domain/values"
	"github.com/schema-creator/services/sql-service/pkg/ddl"
)

type projectService struct {
	projectRepository repository.ProjectRepository
}

func NewProjectService(pr repository.ProjectRepository) service.ProjectService {
	return &projectService{pr}
}

func (ps *projectService) CreateProject(ctx context.Context, args values.CreateProject) (*values.Project, error) {
	return ps.projectRepository.Create(ctx, args)
}

func (ps *projectService) GetProjectsByUserID(ctx context.Context, userID string) ([]*values.Project, error) {
	return ps.projectRepository.GetProjectsByUserID(ctx, userID)
}

func (ps *projectService) DeleteAllByUserID(ctx context.Context, userID string) error {
	return ps.projectRepository.DeleteAllByUserID(ctx, userID)
}

func (ps *projectService) GetOneByID(ctx context.Context, id string) (*values.Project, error) {
	return ps.projectRepository.GetOneByID(ctx, id)
}

func (ps *projectService) UpdateProject(ctx context.Context, args values.Project) (*values.Project, error) {
	return ps.projectRepository.UpdateProject(ctx, args)
}

func (ps *projectService) Delete(ctx context.Context, id string) error {
	return ps.projectRepository.Delete(ctx, id)
}

func (ps *projectService) JoinProject(ctx context.Context, userID string, projectID string) (*values.Project, error) {
	return ps.projectRepository.JoinProject(ctx, userID, projectID)
}

func (ps *projectService) ConvertDDL(ctx context.Context, id string, convertType ddl.ConvertType) (*response.ConvertDDL, error) {
	project, err := ps.projectRepository.GetOneByID(ctx, id)
	if err != nil {
		return nil, err
	}

	convert, err := ddl.Convert(project.Editor, convertType)
	if err != nil {
		return nil, err
	}

	return &response.ConvertDDL{
		ProjectID:    id,
		ProjectTitle: project.Title,
		Data:         convert,
	}, nil

}
