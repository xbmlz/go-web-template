package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xbmlz/go-web-template/api/constant"
	"github.com/xbmlz/go-web-template/internal/token"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		claims, err := token.Provider.Validate(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": constant.ErrInvalidToken.Error()})
		}

		c.Set(constant.CtxUserClaimsKey, claims)
		c.Next()
	}
}
