package middleware

import (
	"go01-airbnb/pkg/common"

	"github.com/gin-gonic/gin"
)

func (m *middleareManager) Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// Để có thể chạy tới middleware Recover mặc định của Gin ta cần tiếp tục panic err
					panic(err)
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
