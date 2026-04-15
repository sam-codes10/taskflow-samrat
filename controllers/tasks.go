package controllers

import (
	apihelpers "taskflow-samrat/apiRes"
	"taskflow-samrat/models"
	"taskflow-samrat/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Tags task
// @Description Create Task using project id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param payload body models.CreateAndUpdateTaskReq true "Create Task Payload"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId/tasks [post]
func CreateTaskUsingProjectId(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)

	var task models.CreateAndUpdateTaskReq
	if err := c.ShouldBindJSON(&task); err != nil {
		logrus.Error("failed to decode payload : "+err.Error(), " | reqId: "+reqH.ReqId)
		apihelpers.SendBadRequestFromController(c, "Invalid request payload")
		return
	}
	code, res := services.CreateTaskUsingProjectId(task, reqH.UserId)
	apiName := "/projects/:projectId/tasks [POST]"
	apihelpers.CustomResponse(c, code, res, apiName)
}

// @Tags task
// @Description Get All Tasks By Project Id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param projectId path string true "Project ID"
// @Param status query string false "Status"
// @Param assignee_id query string false "Assignee ID"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /projects/:projectId/tasks [get]
func GetAllTasksByProjectId(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)

	projectId := c.Param("projectId")
	status := c.Query("status")
	assignee_id := c.Query("assignee_id")
	code, res := services.GetAllTasksByProjectId(projectId, status, assignee_id, reqH.UserId)
	apiName := "/projects/:projectId/tasks [GET]"
	apihelpers.CustomResponse(c, code, res, apiName)
}

// @Tags task
// @Description Get Task By Id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param taskId path string true "Task ID"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /tasks/:taskId [get]
func GetTaskById(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)

	taskId := c.Param("taskId")
	code, res := services.GetTaskById(taskId, reqH.UserId)
	apiName := "/tasks/:taskId [GET]"
	apihelpers.CustomResponse(c, code, res, apiName)
}

// @Tags task
// @Description Update Task By Id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param taskId path string true "Task ID"
// @Param payload body models.CreateAndUpdateTaskReq true "Update Task Payload"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /tasks/:taskId [patch]
func UpdateTaskById(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)

	taskId := c.Param("taskId")
	var task models.CreateAndUpdateTaskReq
	if err := c.ShouldBindJSON(&task); err != nil {
		logrus.Error("failed to decode payload : "+err.Error(), " | reqId: "+reqH.ReqId)
		apihelpers.SendBadRequestFromController(c, "Invalid request payload")
		return
	}
	code, res := services.UpdateTaskById(taskId, task, reqH.UserId)
	apiName := "/tasks/:taskId [PATCH]"
	apihelpers.CustomResponse(c, code, res, apiName)
}

// @Tags task
// @Description Delete Task By Id
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization Header"
// @Param taskId path string true "Task ID"
// @Success 200 {object} apihelpers.ApiResponse
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /tasks/:taskId [delete]
func DeleteTaskById(c *gin.Context) {
	cRH, _ := c.Get("reqH")
	reqH := cRH.(models.RequestHeader)

	taskId := c.Param("taskId")
	code, res := services.DeleteTaskById(taskId, reqH.UserId)
	apiName := "/tasks/:taskId [DELETE]"
	apihelpers.CustomResponse(c, code, res, apiName)
}
