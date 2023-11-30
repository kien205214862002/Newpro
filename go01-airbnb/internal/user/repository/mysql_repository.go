package userrepository

import (
	"context"

	usermodel "go01-airbnb/internal/user/model"
	"go01-airbnb/pkg/common"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Constructor
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, data *usermodel.UserRegister) error {
	db := r.db.Begin()

	if err := db.Table(usermodel.User{}.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

func (r *userRepository) FindDataWithCondition(ctx context.Context, conditions map[string]any) (*usermodel.User, error) {
	var user usermodel.User

	if err := r.db.Table(usermodel.User{}.TableName()).Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(usermodel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
