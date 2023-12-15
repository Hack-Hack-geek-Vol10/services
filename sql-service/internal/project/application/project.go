package application

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/schema-creator/services/sql-service/internal/project/domain/repository"
	"github.com/schema-creator/services/sql-service/internal/project/domain/request"
	"github.com/schema-creator/services/sql-service/internal/project/domain/service"
	"github.com/schema-creator/services/sql-service/internal/project/domain/values"
	"github.com/schema-creator/services/sql-service/internal/project/services"
	"github.com/schema-creator/services/sql-service/pkg/logger"
)

type ProjectApplication interface {
	CreateProject(gin *gin.Context)
	GetProject(gin *gin.Context)
	GetProjectsByUserID(gin *gin.Context)
	UpdateProject(gin *gin.Context)
	DeleteProject(gin *gin.Context)
	DeleteAllByUserID(gin *gin.Context)
	JoinProject(gin *gin.Context)
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

// @Summary Create Project
// @Description Create Project
// @Tags Project
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Param title body string true "title"
// @Success 200 {object} string
// @Router /users/{user_id}/projects [post]
func (pa *projectApplication) CreateProject(gin *gin.Context) {
	var reqURI request.UserRequestWildcard
	var reqBody request.CreateProject
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := gin.BindJSON(&reqBody); err != nil {
		pa.logger.Warnf("Failed to bind json: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}
	project, err := pa.projectService.CreateProject(gin, values.CreateProject{
		ProjectID: uuid.New().String(),
		Title:     reqBody.Title,
		Users:     []string{reqURI.UserID},
		OwnerID:   reqURI.UserID,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	})
	if err != nil {
		pa.logger.Warnf("Failed to create project: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}

	gin.JSON(http.StatusOK, project)
}

// @Summary Get Project
// @Description Get Project
// @Tags Project
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Param project_id path string true "project_id"
// @Success 200 {object} string
// @Router /users/{user_id}/projects/{project_id} [get]
func (pa *projectApplication) GetProject(gin *gin.Context) {
	var reqURI request.ProjectUserRequestWildcard
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}

	project, err := pa.projectService.GetOneByID(gin, reqURI.ProjectID)
	if err != nil {
		pa.logger.Warnf("Failed to get project: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}

	gin.JSON(http.StatusOK, project)
}

// @Summary Get All Projects
// @Description Get All User Projects
// @Tags Project
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Success 200 {object} string
// @Router /users/{user_id}/projects [get]
func (pa *projectApplication) GetProjectsByUserID(gin *gin.Context) {
	var reqURI request.UserRequestWildcard
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}

	projects, err := pa.projectService.GetProjectsByUserID(gin, reqURI.UserID)
	if err != nil {
		pa.logger.Warnf("Failed to get projects: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}

	gin.JSON(http.StatusOK, projects)
}

// @Summary Update Project
// @Description Update Project
// @Tags Project
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Param project_id path string true "project_id"
// @Param title body string true "title"
// @Success 200 {object} string
// @Router /users/{user_id}/projects/{project_id} [put]
func (pa *projectApplication) UpdateProject(gin *gin.Context) {
	var reqURI request.ProjectUserRequestWildcard
	var reqBody request.CreateProject
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := gin.BindJSON(&reqBody); err != nil {
		pa.logger.Warnf("Failed to bind json: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}

	project, err := pa.projectService.UpdateProject(gin, values.Project{
		ProjectID: reqURI.ProjectID,
		Title:     reqBody.Title,
		OwnerID:   reqURI.UserID,
		UpdateAt:  time.Now(),
	})
	if err != nil {
		pa.logger.Warnf("Failed to update project: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}

	gin.JSON(http.StatusOK, project)
}

// @Summary Delete Project
// @Description Delete Project
// @Tags Project
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Param project_id path string true "project_id"
// @Success 200 {object} string
// @Router /users/{user_id}/projects/{project_id} [delete]
func (pa *projectApplication) DeleteProject(gin *gin.Context) {
	var reqURI request.ProjectUserRequestWildcard
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := pa.projectService.Delete(gin, reqURI.ProjectID); err != nil {
		pa.logger.Warnf("Failed to delete project: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}

	gin.JSON(http.StatusOK, nil)
}

// @Summary Delete All Projects
// @Description Delete All User Projects
// @Tags Project
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Success 200 {object} string
// @Router /users/{user_id}/projects [delete]
func (pa *projectApplication) DeleteAllByUserID(gin *gin.Context) {
	var reqURI request.UserRequestWildcard
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := pa.projectService.DeleteAllByUserID(gin, reqURI.UserID); err != nil {
		pa.logger.Warnf("Failed to delete project: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}

	gin.JSON(http.StatusOK, nil)
}

// @Summary Join Project
// @Description Join Project
// @Tags Project
// @Accept  json
// @Produce  json
// @Param user_id path string true "user_id"
// @Param project_id path string true "project_id"
// @Success 200 {object} string
// @Router /users/{user_id}/projects/{project_id} [patch]
func (pa *projectApplication) JoinProject(gin *gin.Context) {
	var reqURI request.ProjectUserRequestWildcard
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}

	project, err := pa.projectService.JoinProject(gin, reqURI.ProjectID, reqURI.UserID)
	if err != nil {
		pa.logger.Warnf("Failed to join project: %s", err.Error())
		gin.JSON(http.StatusInternalServerError, nil)
		return
	}

	gin.JSON(http.StatusOK, project)
}

func (pa *projectApplication) WsUpdateProject(gin *gin.Context) {
	var reqURI request.ProjectUserRequestWildcard
	var reqBody request.WsUpdateProject
	if err := gin.ShouldBindUri(&reqURI); err != nil {
		pa.logger.Warnf("Failed to bind uri: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
	}
	if err := gin.BindJSON(&reqBody); err != nil {
		pa.logger.Warnf("Failed to bind json: %s", err.Error())
		gin.JSON(http.StatusBadRequest, nil)
		return
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
