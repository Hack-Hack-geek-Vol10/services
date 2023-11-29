package usecase

import (
	"context"

	project "github.com/Hack-Hack-geek-Vol10/services/project-service/api/v1"
	"github.com/Hack-Hack-geek-Vol10/services/project-service/internal/domain"
	"github.com/Hack-Hack-geek-Vol10/services/project-service/internal/infra"
	"github.com/google/uuid"
)

type projectService struct {
	project.UnimplementedProjectServiceServer
	projectRepo infra.ProjectRepo
}

func NewProjectService(projectRepo infra.ProjectRepo) project.ProjectServiceServer {
	return &projectService{
		projectRepo: projectRepo,
	}
}

func (s *projectService) CreateProject(ctx context.Context, arg *project.CreateProjectRequest) (*project.ProjectDetails, error) {
	if len(arg.Title) == 0 {
		arg.Title = "untitled"
	}

	param := domain.CreateProjectParam{
		ProjectID: uuid.New().String(),
		Title:     arg.Title,
	}

	result, err := s.projectRepo.Create(ctx, param)
	if err != nil {
		return nil, err
	}

	return &project.ProjectDetails{
		ProjectId:  result.ProjectID,
		Title:      result.Title,
		LastImage:  result.LastImage,
		IsPersonal: true,
	}, nil
}

func (s *projectService) GetProject(ctx context.Context, arg *project.GetProjectRequest) (*project.ProjectDetails, error) {
	projectInfo, err := s.projectRepo.ReadOne(ctx, arg.ProjectId)
	if err != nil {
		return nil, err
	}

	return &project.ProjectDetails{
		ProjectId:  projectInfo.ProjectID,
		Title:      projectInfo.Title,
		LastImage:  projectInfo.LastImage,
		IsPersonal: projectInfo.IsPersonal,
	}, nil
}

func (s *projectService) ListProjects(ctx context.Context, arg *project.ListProjectsRequest) (*project.ListProjectsResponse, error) {
	if arg.Limit == 0 {
		arg.Limit = 10
	}

	if arg.Offset == 0 {
		arg.Offset = 1
	}

	result, err := s.projectRepo.ReadAll(ctx, domain.ReadProjectsParam{
		UserID: arg.UserId,
		Limit:  arg.Limit,
		Offset: (arg.Offset - 1) * arg.Limit,
	})

	if err != nil {
		return nil, err
	}

	var projects []*project.ProjectDetails
	for _, info := range result {
		projects = append(projects, &project.ProjectDetails{
			ProjectId:  info.ProjectID,
			Title:      info.Title,
			LastImage:  info.LastImage,
			IsPersonal: info.IsPersonal,
		})
	}

	return &project.ListProjectsResponse{
		Projects: projects,
	}, nil
}

func (s *projectService) UpdateTitle(ctx context.Context, arg *project.UpdateTitleRequest) (*project.ProjectDetails, error) {
	projectInfo, err := s.projectRepo.UpdateTitle(ctx, arg.ProjectId, arg.Title)
	if err != nil {
		return nil, err
	}

	return &project.ProjectDetails{
		ProjectId:  projectInfo.ProjectID,
		Title:      projectInfo.Title,
		LastImage:  projectInfo.LastImage,
		IsPersonal: projectInfo.IsPersonal,
	}, nil
}

func (s *projectService) UpdateImage(ctx context.Context, arg *project.UpdateImageRequest) (*project.ProjectDetails, error) {
	projectInfo, err := s.projectRepo.UpdateLastImage(ctx, arg.ProjectId, arg.LastImage)
	if err != nil {
		return nil, err
	}

	return &project.ProjectDetails{
		ProjectId:  projectInfo.ProjectID,
		Title:      projectInfo.Title,
		LastImage:  projectInfo.LastImage,
		IsPersonal: projectInfo.IsPersonal,
	}, nil
}

func (s *projectService) DeleteProject(ctx context.Context, arg *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {
	err := s.projectRepo.Delete(ctx, arg.ProjectId)
	if err != nil {
		return nil, err
	}

	return &project.DeleteProjectResponse{
		ProjectId: arg.ProjectId,
	}, nil
}

func (s *projectService) mustEmbedUnimplementedProjectServiceServer() {}
