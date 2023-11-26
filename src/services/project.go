package services

import (
	"context"

	project "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/project-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
	"github.com/Hack-Hack-geek-Vol10/services/src/storages"
	"github.com/google/uuid"
)

type projectService struct {
	project.UnimplementedProjectServiceServer
	projectRepo storages.ProjectRepo
}

func NewProjectService(projectRepo storages.ProjectRepo) project.ProjectServiceServer {
	return &projectService{
		projectRepo: projectRepo,
	}
}

func (s *projectService) CreateProject(ctx context.Context, arg *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	if len(arg.Title) == 0 {
		arg.Title = "untitled"
	}

	param := domain.CreateProjectParam{
		ProjectID: uuid.New().String(),
		Title:     arg.Title,
	}

	if err := s.projectRepo.Create(ctx, param); err != nil {
		return nil, err
	}

	return &project.CreateProjectResponse{
		Id:         param.ProjectID,
		Title:      param.Title,
		LastImage:  "",
		IsPersonal: true,
	}, nil
}
