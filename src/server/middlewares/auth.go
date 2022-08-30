package middlewares

import (
	"errors"
	"net/http"
	"opengin/server/schemas"
	"opengin/server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func validateToken(tokenHeader string) (*utils.CustomClaims, error) {
	if tokenHeader == "" {
		return nil, errors.New("invalid token")
	}
	tokenSlice := strings.SplitN(tokenHeader, " ", 2)
	if len(tokenSlice) != 2 || tokenSlice[0] != "Bearer" {
		return nil, errors.New("invalid token")
	}

	token := tokenSlice[1]
	claims, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		claims, err := validateToken(tokenHeader)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, schemas.UniResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})
			ctx.Abort()
			return
		}

		ctx.Set("UserId", claims.Subject)
		ctx.Set("ClientId", claims.ClientId)
		ctx.Next()
	}
}
