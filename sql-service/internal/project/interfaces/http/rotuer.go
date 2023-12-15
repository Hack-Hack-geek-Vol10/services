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

	router.Gin.GET("/users/:user_id/projects/:project_id/ddl", projectApplication.ConvertDDL)

	return router.Gin
}
