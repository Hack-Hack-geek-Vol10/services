package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schema-creator/services/sql-service/internal/project/domain/repository"
	"github.com/schema-creator/services/sql-service/internal/project/domain/request"
	"github.com/schema-creator/services/sql-service/internal/project/domain/service"
	"github.com/schema-creator/services/sql-service/internal/project/services"
	"github.com/schema-creator/services/sql-service/pkg/logger"
)

type ProjectApplication interface {
	ConvertDDL(gin *gin.Context)
}

type projectApplication struct {
	logger         logger.Logger
	projectService service.ProjectService
}

func NewProjectApplication(l logger.Logger, pr repository.ProjectRepository) ProjectApplication {
	return &projectApplication{
		logger:         l,
		projectService: services.NewProjectService(pr),
	}
}

func (pa *projectApplication) ConvertDDL(gin *gin.Context) {
	var reqURI request.ProjectUserRequestWildcard
	var reqBody request.ConvertDDL
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := gin.ShouldBindQuery(&reqBody); err != nil {
		pa.logger.Warnf("Failed to bind json: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}

	ddl, err := pa.projectService.ConvertDDL(gin, reqURI.ProjectID, reqBody.ConvertType)
	if err != nil {
		pa.logger.Warnf("Failed to convert ddl: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}
	gin.JSON(http.StatusOK, ddl)
}
