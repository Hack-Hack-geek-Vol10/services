package http

import (
	"github.com/gin-gonic/gin"
	"github.com/schema-creator/services/sql-service/internal/project/application"
	"github.com/schema-creator/services/sql-service/internal/project/domain/repository"
	"github.com/schema-creator/services/sql-service/pkg/logger"
)

type UserRouter struct {
	Gin    *gin.Engine
	logger logger.Logger
}

func NewProjectRouter(gin *gin.Engine, l logger.Logger, pr repository.ProjectRepository) *gin.Engine {
	router := &UserRouter{
		Gin:    gin,
		logger: l,
	}
	router.setupCors()

	projectApplication := application.NewProjectApplication(l, pr)

	router.Gin.GET("/users/:user_id/projects", projectApplication.GetProjectsByUserID)
	router.Gin.POST("/users/:user_id/projects", projectApplication.CreateProject)
	router.Gin.DELETE("/users/:user_id/projects", projectApplication.DeleteAllByUserID)

	router.Gin.GET("/users/:user_id/projects/:project_id", projectApplication.GetProject)
	router.Gin.PUT("/users/:user_id/projects/:project_id", projectApplication.UpdateProject)
	router.Gin.PATCH("/users/:user_id/projects/:project_id", projectApplication.JoinProject)
	router.Gin.DELETE("/users/:user_id/projects/:project_id", projectApplication.DeleteProject)
	router.Gin.GET("/users/:user_id/projects/:project_id/ddl", projectApplication.ConvertDDL)

	return router.Gin
}
