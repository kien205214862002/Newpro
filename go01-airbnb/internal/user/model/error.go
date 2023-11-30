package usermodel

import (
	"errors"
	"go01-airbnb/pkg/common"
)

// Định nghĩa các error cho riêng phần user
var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
	)

	ErrCannotCreateAccount = common.NewCustomError(
		errors.New("can not create your account"),
		"can not create your account",
	)
)
