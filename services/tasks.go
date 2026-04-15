package services

import (
	apihelpers "taskflow-samrat/apiRes"
	"taskflow-samrat/db"
	"taskflow-samrat/models"
)

func CreateTaskUsingProjectId(task models.CreateAndUpdateTaskReq, projectId, ownerId string) (int, apihelpers.ApiResponse) {
	valid, err := db.ValidateProjectIdAndOwnerId(projectId, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to validate project id and owner id due to error " + err.Error())
	}
	if !valid {
		return apihelpers.ReturnForbiddenRequestFromService("project id and owner id are not valid")
	}
	res, err := db.CreateTaskUsingProjectId(task, projectId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to create task due to error " + err.Error())
	}
	return apihelpers.ReturnSuccessResponseFromService("Task created successfully", res)
}

func GetAllTasksByProjectId(projectId, status, assignee_id string, ownerId string) (int, apihelpers.ApiResponse) {
	valid, err := db.ValidateProjectIdAndOwnerId(projectId, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to validate project id and owner id due to error " + err.Error())
	}
	if !valid {
		return apihelpers.ReturnForbiddenRequestFromService("project id and owner id are not valid")
	}
	res, err := db.GetAllTasksByProjectId(projectId, status, assignee_id)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to get all tasks due to error " + err.Error())
	}
	return apihelpers.ReturnSuccessResponseFromService("All tasks fetched successfully", res)
}

func GetTaskById(taskId string, ownerId string) (int, apihelpers.ApiResponse) {
	valid, err := db.ValidateTaskIdAndOwnerId(taskId, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to validate task id and owner id due to error " + err.Error())
	}
	if !valid {
		return apihelpers.ReturnForbiddenRequestFromService("task id and owner id are not valid")
	}
	res, err := db.GetTaskById(taskId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to get task due to error " + err.Error())
	}
	return apihelpers.ReturnSuccessResponseFromService("Task fetched successfully", res)
}

func UpdateTaskById(taskId string, task models.CreateAndUpdateTaskReq, ownerId string) (int, apihelpers.ApiResponse) {
	valid, err := db.ValidateTaskIdAndOwnerId(taskId, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to validate task id and owner id due to error " + err.Error())
	}
	if !valid {
		return apihelpers.ReturnForbiddenRequestFromService("task id and owner id are not valid")
	}
	res, err := db.UpdateTaskById(taskId, task)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to update task due to error " + err.Error())
	}
	return apihelpers.ReturnSuccessResponseFromService("Task updated successfully", res)
}

func DeleteTaskById(taskId string, ownerId string) (int, apihelpers.ApiResponse) {
	valid, err := db.ValidateTaskIdAndOwnerId(taskId, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to validate task id and owner id due to error " + err.Error())
	}
	if !valid {
		return apihelpers.ReturnForbiddenRequestFromService("task id and owner id are not valid")
	}
	res, err := db.DeleteTaskById(taskId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to delete task due to error " + err.Error())
	}
	return apihelpers.ReturnSuccessResponseFromService("Task deleted successfully", res)
}
