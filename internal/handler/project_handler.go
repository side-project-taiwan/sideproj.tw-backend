package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spt/internal/model"
	"spt/internal/usecase"
)

type ProjectHandler interface {
	CreateProject(c *gin.Context)
	GetProjectList(c *gin.Context)
	GetProjectByID(c *gin.Context)
}

type projectHandler struct {
	usecase usecase.ProjectUsecase
}

func NewProjectHandler(usecase usecase.ProjectUsecase) ProjectHandler {
	return &projectHandler{
		usecase: usecase,
	}
}

// CreateProject  godoc
// @Summary     crete project
// @Description Create a new project
// @Tags        Project Management
// @Accept      application/json
// @Produce     application/json
// @Param       project body model.ProjectCreate true "Project"
// @Success     200 {object} interface{}
// @Router      /project [post]
func (h *projectHandler) CreateProject(c *gin.Context) {
	g := Gin{c}

	var req model.ProjectCreate

	err := c.Bind(&req)
	if err != nil {
		g.Response(http.StatusBadRequest, http.StatusBadRequest, "", err.Error())
		return
	}

	err = h.usecase.CreateNewProject(c, &req)
	if err != nil {
		g.Response(http.StatusInternalServerError, http.StatusInternalServerError, "", err.Error())
		return
	}
	g.Response(http.StatusOK, http.StatusOK, "Get project list successfully", "")
}

// GetProjectList  godoc
// @Summary     Get project list
// @Description Retrieves a list of projects
// @Tags        Project Management
// @Accept      application/json
// @Produce     application/json
// @Success     200 {object} []model.Project
// @Router      /projects [get]
func (h *projectHandler) GetProjectList(c *gin.Context) {
	g := Gin{c}
	data, err := h.usecase.GetProjectList(c)
	if err != nil {
		g.Response(http.StatusInternalServerError, http.StatusInternalServerError, "", err.Error())
		return
	}
	g.Response(http.StatusOK, http.StatusOK, "Get project list successfully", data)
}

// GetProjectByID  godoc
// @Summary     Get project
// @Description Retrieves a project by id
// @Tags        Project Management
// @Accept      application/json
// @Produce     application/json
// @Param       id path string true "Project ID"
// @Success     200 {object} model.Project
// @Router      /project/{id} [get]
func (h *projectHandler) GetProjectByID(c *gin.Context) {
	g := Gin{c}

	projectID := c.Param("id")

	if projectID == "" {
		g.Response(http.StatusBadRequest, http.StatusBadRequest, "projectID is empty", nil)
		return
	}

	data, err := h.usecase.GetProjectByID(c, projectID)
	if err != nil {
		g.Response(http.StatusInternalServerError, http.StatusInternalServerError, "", err.Error())
		return
	}
	g.Response(http.StatusOK, http.StatusOK, "Get project successfully", data)
}
