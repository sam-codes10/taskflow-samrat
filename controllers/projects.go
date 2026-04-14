package controllers

import (
	"taskflow-samrat/apiRes"
	"taskflow-samrat/models"
	"taskflow-samrat/services"

	"github.com/gin-gonic/gin"
)

// @Tags projects
// @Description Create Project
// @Accept json
// @Produce json
// @Param payload body models.CreateProjectReq true "Create Project Payload"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects [post]
func CreateProject(c *gin.Context) {
	var payload models.CreateProjectReq
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		apihelpers.SendBadRequestFromController(c, "Invalid request payload")
		return
	}

	statusCode, response := services.CreateProject(payload, c.Value("userId").(string))
	apiName := "/projects [POST]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Get All Projects
// @Accept json
// @Produce json
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects [get]
func GetAllProjects(c *gin.Context) {
	statusCode, response := services.GetAllProjects(c.Value("userId").(string))
	apiName := "/projects [GET]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Get Project By Id
// @Accept json
// @Produce json
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId [get]
func GetProjectById(c *gin.Context) {
	projectId := c.Param("projectId")
	statusCode, response := services.GetProjectById(projectId)
	apiName := "/projects/:projectId [GET]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Update Project By Id
// @Accept json
// @Produce json
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId [patch]
func UpdateProjectById(c *gin.Context) {
	projectId := c.Param("projectId")
	var payload models.UpdateProjectReq
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		apihelpers.SendBadRequestFromController(c, "Invalid request payload")
		return
	}

	statusCode, response := services.UpdateProjectById(projectId, payload)
	apiName := "/projects/:projectId [PATCH]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Delete Project By Id
// @Accept json
// @Produce json
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId [delete]
func DeleteProjectById(c *gin.Context) {
	projectId := c.Param("projectId")
	statusCode, response := services.DeleteProjectById(projectId)
	apiName := "/projects/:projectId [DELETE]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}
