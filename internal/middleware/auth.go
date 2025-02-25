package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
	"gitlab.com/gma-vietnam/tanca-event/pkg/jwt"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
)

var (
	errInvalidToken = pkgErrors.NewHTTPError(401, "invalid token")
)

func (m Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if tokenString == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		payload, err := m.jwtManager.Verify(tokenString)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		ctx = jwt.SetPayloadToContext(ctx, payload)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
