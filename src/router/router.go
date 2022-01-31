package router

import (
	"fmt"
	"go-api/src/helpers/response"
	"go-api/src/middleware"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	app := gin.Default()

	// middleware

	return app
}

var middlewareJWT = middleware.NewJWT()

func AuthorizedJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", 1)

		parse, err := middlewareJWT.ValidateToken(token)
		if parse == nil {
			response.Error(ctx, response.Response{
				Status: http.StatusUnauthorized,
				Error:  http.StatusText(http.StatusUnauthorized),
			})
			return
		}

		if err != nil || !parse.Valid {
			response.Error(ctx, response.Response{
				Status: http.StatusUnauthorized,
				Error:  err.Error(),
			})
			return
		}

		claims := parse.Claims.(jwt.MapClaims)
		fmt.Println("parse", claims)
		ctx.Next()
	}
}
