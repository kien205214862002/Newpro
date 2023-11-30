package userhttp

import (
	"context"
	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Register(context.Context, *usermodel.UserRegister) error
	Login(context.Context, *usermodel.UserLogin) (*utils.Token, error)
}

type userHandler struct {
	userUC UserUsecase
}

func NewUserHandler(userUC UserUsecase) *userHandler {
	return &userHandler{userUC}
}

func (hdl *userHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserRegister

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrBadRequest(err))
		}

		if err := hdl.userUC.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"data": data.Id})
	}
}

func (hdl *userHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials usermodel.UserLogin

		if err := c.ShouldBind(&credentials); err != nil {
			panic(common.ErrBadRequest(err))
		}

		token, err := hdl.userUC.Login(c.Request.Context(), &credentials)
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{"data": token})
	}
}
