package middleware

import (
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"

	"github.com/gin-gonic/gin"
)

// Cần phải gọi middlware RequiredAuth trước
func (m *middleareManager) RequiredRoles(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(*usermodel.User)

		for i := range roles {
			if user.GetUserRole() == roles[i] {
				c.Next()
				return
			}
		}

		panic(common.ErrForbidden(nil))
	}
}
