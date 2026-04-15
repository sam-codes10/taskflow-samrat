package services

import (
	"net/http"
	apihelpers "taskflow-samrat/apiRes"
	"taskflow-samrat/db"
	"taskflow-samrat/models"
)

func CreateProject(payload models.CreateProjectReq, ownerId string) (int, apihelpers.ApiResponse) {
	dbRes, err := db.CreateProject(payload, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to create project due to error " + err.Error())
	}
	return http.StatusOK, apihelpers.ApiResponse{
		Status:  true,
		Message: "Project created successfully",
		Data:    dbRes,
	}
}

func GetAllProjects(ownerId string) (int, apihelpers.ApiResponse) {
	dbRes, err := db.GetAllProjects(ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to get all projects due to error " + err.Error())
	}
	return http.StatusOK, apihelpers.ApiResponse{
		Status:  true,
		Message: "All projects fetched successfully",
		Data:    dbRes,
	}
}

func GetProjectById(projectId string, ownerId string) (int, apihelpers.ApiResponse) {
	dbRes, err := db.GetProjectById(projectId, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to get project by id due to error " + err.Error())
	}
	return http.StatusOK, apihelpers.ApiResponse{
		Status:  true,
		Message: "Project fetched successfully",
		Data:    dbRes,
	}
}

func UpdateProjectById(projectId string, payload models.UpdateProjectReq, ownerId string) (int, apihelpers.ApiResponse) {
	dbRes, err := db.UpdateProjectById(projectId, payload, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to update project by id due to error " + err.Error())
	}
	return http.StatusOK, apihelpers.ApiResponse{
		Status:  true,
		Message: "Project updated successfully",
		Data:    dbRes,
	}
}

func DeleteProjectById(projectId , ownerId string) (int, apihelpers.ApiResponse) {
	dbRes, err := db.DeleteProjectById(projectId, ownerId)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("unable to delete project by id due to error " + err.Error())
	}
	return http.StatusOK, apihelpers.ApiResponse{
		Status:  true,
		Message: "Project deleted successfully",
		Data:    dbRes,
	}
}
