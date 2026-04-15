package controllers

import (
	apihelpers "taskflow-samrat/apiRes"
	"taskflow-samrat/models"
	"taskflow-samrat/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// @Tags auth 
// @Description User Registration
// @Accept json
// @Produce json
// @Param payload body models.UserRegister true "User Registration Payload"
// @Success 200 {object} apihelpers.ApiResponse{data=models.UserAuthRes}
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var payload models.UserRegister
	reqH := c.MustGet("reqH").(models.RequestHeader)
	if err := c.ShouldBindJSON(&payload); err != nil {
		logrus.Error("failed to decode payload : "+err.Error(), " | reqId: "+reqH.ReqId)
		apihelpers.HandleValidationError(c, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		logrus.Error("failed to validate payload : "+err.Error(), " | reqId: "+reqH.ReqId)
		apihelpers.HandleValidationError(c, err)
		return
	}

	code, apiRes := services.RegisterUser(payload)
	apiName := "/auth/register"
	apihelpers.CustomResponse(c, code, apiRes, apiName)
}

// @Tags auth 
// @Description User Login
// @Accept json
// @Produce json
// @Param payload body models.UserLogin true "User Login Payload"
// @Success 200 {object} apihelpers.ApiResponse{data=models.UserAuthRes}
// @Failure 400 {object} apihelpers.ApiResponse
// @Failure 409 {object} apihelpers.ApiResponse
// @Failure 500 {object} apihelpers.ApiResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var payload models.UserLogin
	reqH := c.MustGet("reqH").(models.RequestHeader)
	if err := c.ShouldBindJSON(&payload); err != nil {
		logrus.Error("failed to decode payload : "+err.Error(), " | reqId: "+reqH.ReqId)
		apihelpers.HandleValidationError(c, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		logrus.Error("failed to validate payload : "+err.Error(), " | reqId: "+reqH.ReqId)
		apihelpers.HandleValidationError(c, err)
		return
	}

	code, apiRes := services.LoginUser(payload)
	apiName := "/auth/login"
	apihelpers.CustomResponse(c, code, apiRes, apiName)
}
