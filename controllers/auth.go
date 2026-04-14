package controllers

import (
	apihelpers "taskflow-samrat/apiRes"
	"taskflow-samrat/models"
	"taskflow-samrat/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func Register(c *gin.Context) {
	var payload models.UserRegister
	reqHeader := c.MustGet("reqHeader").(models.RequestHeader)
	if err := c.ShouldBindJSON(&payload); err != nil {
		logrus.Error("failed to decode payload : "+err.Error(), " | reqId: "+reqHeader.ReqId)
		apihelpers.SendBadRequestFromController(c, "failed to decode payload : "+err.Error())
		return
	}

	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		logrus.Error("failed to validate payload : "+err.Error(), " | reqId: "+reqHeader.ReqId)
		apihelpers.SendBadRequestFromController(c, "failed to validate payload : "+err.Error())
		return
	}

	code, apiRes := services.RegisterUser(payload)
	apiName := "/auth/register"
	apihelpers.CustomResponse(c, code, apiRes, apiName)
}

func Login(c *gin.Context) {
	var payload models.UserLogin
	reqHeader := c.MustGet("reqHeader").(models.RequestHeader)
	if err := c.ShouldBindJSON(&payload); err != nil {
		logrus.Error("failed to decode payload : "+err.Error(), " | reqId: "+reqHeader.ReqId)
		apihelpers.SendBadRequestFromController(c, "failed to decode payload : "+err.Error())
		return
	}

	validate := validator.New()
	err := validate.Struct(payload)
	if err != nil {
		logrus.Error("failed to validate payload : "+err.Error(), " | reqId: "+reqHeader.ReqId)
		apihelpers.SendBadRequestFromController(c, "failed to validate payload : "+err.Error())
		return
	}

	code, apiRes := services.LoginUser(payload)
	apiName := "/auth/login"
	apihelpers.CustomResponse(c, code, apiRes, apiName)
}
