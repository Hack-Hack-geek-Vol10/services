package usecase

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
	sql "github.com/schema-creator/services/sql-service/api/v1"
	"github.com/schema-creator/services/sql-service/internal/domain"
	"github.com/schema-creator/services/sql-service/internal/infra"
)

type sqlService struct {
	sql.UnimplementedSqlServiceServer
	sqlRepo infra.SqlRepo
}

func NewSqlService(sqlRepo infra.ProjectRepo) sql.ProjectServer {
	return &sqlService{
		sqlRepo: sqlRepo,
	}
}

func (s *sqlService) CreateProject(ctx context.Context, arg *sql.CreateProjectRequest) (*sql.ProjectDetails, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-CreateProject").End()
	if len(arg.Title) == 0 {
		arg.Title = "untitled"
	}

	param := domain.CreateProjectParam{
		ProjectID: uuid.New().String(),
		Title:     arg.Title,
	}

	result, err := s.sqlRepo.Create(ctx, param)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &sql.ProjectDetails{
		ProjectId:  result.ProjectID,
		Title:      result.Title,
		LastImage:  result.LastImage,
		IsPersonal: true,
	}, nil
}

func (s *sqlService) GetProject(ctx context.Context, arg *sql.GetProjectRequest) (*sql.ProjectDetails, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-GetProject").End()
	projectInfo, err := s.sqlRepo.ReadOne(ctx, arg.ProjectId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &sql.ProjectDetails{
		ProjectId:  projectInfo.ProjectID,
		Title:      projectInfo.Title,
		LastImage:  projectInfo.LastImage,
		IsPersonal: projectInfo.IsPersonal,
	}, nil
}

func (s *sqlService) ListProjects(ctx context.Context, arg *sql.ListProjectsRequest) (*sql.ListProjectsResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-ListProjects").End()
	if arg.Limit == 0 {
		arg.Limit = 1000
	}

	if arg.Offset == 0 {
		arg.Offset = 1
	}

	result, err := s.sqlRepo.ReadAll(ctx, domain.ReadProjectsParam{
		UserID: arg.UserId,
		Limit:  arg.Limit,
		Offset: (arg.Offset - 1) * arg.Limit,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var sqls []*sql.ProjectDetails
	for _, info := range result {
		sqls = append(sqls, &sql.ProjectDetails{
			ProjectId:  info.ProjectID,
			Title:      info.Title,
			LastImage:  info.LastImage,
			IsPersonal: info.IsPersonal,
		})
	}

	return &sql.ListProjectsResponse{
		Projects: sqls,
	}, nil
}

func (s *sqlService) UpdateTitle(ctx context.Context, arg *sql.UpdateTitleRequest) (*sql.ProjectDetails, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-UpdateTitle").End()
	projectInfo, err := s.sqlRepo.UpdateTitle(ctx, arg.ProjectId, arg.Title)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &sql.ProjectDetails{
		ProjectId:  projectInfo.ProjectID,
		Title:      projectInfo.Title,
		LastImage:  projectInfo.LastImage,
		IsPersonal: projectInfo.IsPersonal,
	}, nil
}

func (s *sqlService) UpdateImage(ctx context.Context, arg *sql.UpdateImageRequest) (*sql.ProjectDetails, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-UpdateImage").End()
	projectInfo, err := s.sqlRepo.UpdateLastImage(ctx, arg.ProjectId, arg.LastImage)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &sql.ProjectDetails{
		ProjectId:  projectInfo.ProjectID,
		Title:      projectInfo.Title,
		LastImage:  projectInfo.LastImage,
		IsPersonal: projectInfo.IsPersonal,
	}, nil
}

func (s *sqlService) DeleteProject(ctx context.Context, arg *sql.DeleteProjectRequest) (*sql.DeleteProjectResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("grpc-DeleteProject").End()
	err := s.sqlRepo.Delete(ctx, arg.ProjectId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &sql.DeleteProjectResponse{
		ProjectId: arg.ProjectId,
	}, nil
}

func (s *sqlService) mustEmbedUnimplementedProjectServiceServer() {}
