package models

import "github.com/golang-jwt/jwt/v5"

type RequestHeader struct {
	ReqId         string `json:"reqId"`
	Authorization string `json:"Authorization"`
}

type JwtClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
