package apihelpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnInternalServerErrorFromService(message string) (int, interface{}) {
	return http.StatusInternalServerError, gin.H{"error": "internal server error: " + message}
}

func ReturnConflictRequestFromService(message string) (int, interface{}) {
	return http.StatusConflict, gin.H{"error": "conflict: " + message}
}

func ReturnSuccessResponseFromService(message string, data interface{}) (int, interface{}) {
	return http.StatusOK, ApiResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func SendBadRequestFromController(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}

func SendValidationError(c *gin.Context, fields map[string]string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  "validation failed",
		"fields": fields,
	})
}

func HandleValidationError(c *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		fields := make(map[string]string)
		for _, e := range validationErrors {
			fields[e.Field()] = "is " + e.Tag()
		}
		SendValidationError(c, fields)
		return
	}
	SendBadRequestFromController(c, "Invalid request payload")
}

func SendInternalServerErrorFromController(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error: " + message})
}

func SendSuccessResponseFromController(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ReturnUnauthorizedRequestFromService(message string) (int, interface{}) {
	return http.StatusUnauthorized, gin.H{"error": "unauthorized"}
}

func ReturnForbiddenRequestFromService(message string) (int, interface{}) {
	return http.StatusForbidden, gin.H{"error": "forbidden"}
}

func ReturnUnauthorized(message string) (int, interface{}) {
	return http.StatusUnauthorized, gin.H{"error": "unauthorized"}
}

func ReturnNotFoundRequestFromService(message string) (int, interface{}) {
	return http.StatusNotFound, gin.H{"error": "not found"}
}

func CustomResponse(c *gin.Context, code int, data interface{}, apiName string) {
	// log optionalParams
	logrus.Info("API call completed : ", apiName, " | code: ", code, " | data: ", data)
	// send json res
	c.JSON(code, data)
}
