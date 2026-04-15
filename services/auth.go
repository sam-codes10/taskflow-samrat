package services

import (
	apihelpers "taskflow-samrat/apiRes"
	"taskflow-samrat/db"
	"taskflow-samrat/middleware"
	"taskflow-samrat/models"

	"github.com/lib/pq"
)

func RegisterUser(payload models.UserRegister) (int, interface{}) {
	var res models.UserAuthRes
	dbRes, err := db.RegisterUser(payload)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return apihelpers.ReturnConflictRequestFromService("email already exists" + err.Error())
			}
		}
		return apihelpers.ReturnInternalServerErrorFromService("internal server error" + err.Error())
	}
	res.ID = dbRes.ID
	res.Email = dbRes.Email
	res.Token, err = middleware.GenerateJWT(dbRes)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("internal server error " + err.Error())
	}
	return apihelpers.ReturnSuccessResponseFromService("user registered successfully", res)
}

func LoginUser(payload models.UserLogin) (int, interface{}) {
	var res models.UserAuthRes
	allowed, dbRes, err := db.LoginUser(payload)
	if err != nil && err.Error() != "invalid credentials" {
		return apihelpers.ReturnInternalServerErrorFromService("internal server error " + err.Error())
	}
	if !allowed {
		return apihelpers.ReturnUnauthorizedRequestFromService("invalid credentials")
	}
	res.ID = dbRes.ID
	res.Email = dbRes.Email
	res.Token, err = middleware.GenerateJWT(dbRes)
	if err != nil {
		return apihelpers.ReturnInternalServerErrorFromService("internal server error " + err.Error())
	}
	return apihelpers.ReturnSuccessResponseFromService("user logged in successfully", res)
}
