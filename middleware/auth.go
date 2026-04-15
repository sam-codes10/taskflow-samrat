package middleware

import (
	"errors"
	"os"
	"strings"
	apihelpers "taskflow-samrat/apiRes"
	"taskflow-samrat/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func NoAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqId := uuid.New().String()
		reqHeader := models.RequestHeader{
			ReqId: reqId,
		}
		c.Set("reqH", reqHeader)
		c.Next()
	}
}

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqH models.RequestHeader

		token := c.GetHeader("Authorization")
		if token == "" {
			code, res := apihelpers.ReturnUnauthorized("Authorization token is required")
			c.JSON(code, res)
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		claims, err := parseJWT(token)
		if err != nil {
			code, res := apihelpers.ReturnUnauthorized("Invalid authorization token")
			c.JSON(code, res)
			c.Abort()
			return
		}

		reqH.Email = claims.Email
		reqH.UserId = claims.UserID
		reqH.ReqId = uuid.NewString()

		c.Set("reqH", reqH)
		c.Next()

	}
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "taskflow-samrat",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func parseJWT(tokenString string) (*models.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*models.JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
