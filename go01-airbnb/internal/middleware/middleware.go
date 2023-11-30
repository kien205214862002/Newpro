package middleware

import (
	"context"
	"go01-airbnb/config"
	usermodel "go01-airbnb/internal/user/model"
)

type UserRepository interface {
	FindDataWithCondition(context.Context, map[string]any) (*usermodel.User, error)
}

type middleareManager struct {
	cfg      *config.Config
	userRepo UserRepository
}

func NewMiddlewareManager(cfg *config.Config, userRepo UserRepository) *middleareManager {
	return &middleareManager{cfg, userRepo}
}
