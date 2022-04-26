package middlewares

import (
	"edx_go_2/src/app"
	"edx_go_2/src/config"
	jwtpkg "edx_go_2/src/packages/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("Authorization not valid")
	}

	return parts[1], nil
}

func Auth(ctx *gin.Context) {
	token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

	if err != nil {
		panic(app.ErrNoPermission(err))
		return
	}

	userId, err := jwtpkg.ValidateToken(config.Conf.SecretKey, token)

	if err != nil {
		panic(app.ErrNoPermission(err))
		return
	}

	ctx.Set("user_id", userId)
	ctx.Next()
}
