package apihelpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnInternalServerErrorFromService(message string) (int, ApiResponse) {
	return http.StatusInternalServerError, ApiResponse{
		Status:  false,
		Message: message,
		Data:    nil,
	}
}

func ReturnSuccessResponse(message string, data interface{}) (int, ApiResponse) {
	return http.StatusOK, ApiResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func SendBadRequestFromController(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ApiResponse{
		Status:  false,
		Message: message,
		Data:    nil,
	})
}

func SendInternalServerErrorFromController(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, ApiResponse{
		Status:  false,
		Message: message,
		Data:    nil,
	})
}

func SendSuccessResponseFromController(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}
