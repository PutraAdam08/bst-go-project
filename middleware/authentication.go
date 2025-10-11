package middleware

import (
	"net/http"
	"strings"

	apix "BSTproject.com/utils/api"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
	GetUserByTokenID(token string) (uint, error)
}

func Authenticate(jwtService JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apix.HTTPResponse{
				Message: "Missing Authorization header",
				Data:    nil,
			})
			return
		}

		// ✅ Properly trim "Bearer " prefix and any extra spaces
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apix.HTTPResponse{
				Message: "Invalid or expired token",
				Data:    nil,
			})
			return
		}

		// ✅ Extract user ID from the token
		userID, err := jwtService.GetUserByTokenID(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, apix.HTTPResponse{
				Message: "Invalid token payload",
				Data:    nil,
			})
			return
		}

		ctx.Set("user_id", int(userID))
		ctx.Next()
	}
}
