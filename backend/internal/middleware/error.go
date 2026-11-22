package middleware

import (
	"stockflow/internal/handler"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}

		handler.RespondUsecaseError(c, c.Errors.Last().Err)
	}
}
