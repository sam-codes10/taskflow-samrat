package controllers

import (
	"taskflow-samrat/apiRes"
	"taskflow-samrat/models"
	"taskflow-samrat/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Tags projects
// @Description Create Project
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param payload body models.CreateProjectReq true "Create Project Payload"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects [post]
func CreateProject(c *gin.Context) {
	var payload models.CreateProjectReq
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		logrus.Error("failed to decode payload : "+err.Error(), " | reqId: "+reqH.ReqId)
		apihelpers.SendBadRequestFromController(c, "Invalid request payload")
		return
	}

	statusCode, response := services.CreateProject(payload, reqH.UserId)
	apiName := "/projects [POST]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Get All Projects
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects [get]
func GetAllProjects(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)
	statusCode, response := services.GetAllProjects(reqH.UserId)
	apiName := "/projects [GET]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Get Project By Id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId [get]
func GetProjectById(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)
	projectId := c.Param("projectId")
	statusCode, response := services.GetProjectById(projectId, reqH.UserId)
	apiName := "/projects/:projectId [GET]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Update Project By Id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId [patch]
func UpdateProjectById(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)
	projectId := c.Param("projectId")
	var payload models.UpdateProjectReq
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		apihelpers.SendBadRequestFromController(c, "Invalid request payload")
		return
	}

	statusCode, response := services.UpdateProjectById(projectId, payload, reqH.UserId)
	apiName := "/projects/:projectId [PATCH]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}

// @Tags projects
// @Description Delete Project By Id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId [delete]
func DeleteProjectById(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)
	projectId := c.Param("projectId")
	statusCode, response := services.DeleteProjectById(projectId, reqH.UserId)
	apiName := "/projects/:projectId [DELETE]"
	apihelpers.CustomResponse(c, statusCode, response, apiName)
}
