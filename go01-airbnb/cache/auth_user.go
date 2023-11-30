package cache

import (
	"context"
	"fmt"
	usermodel "go01-airbnb/internal/user/model"
	"time"
)

const cacheKey = "user:%d"

type UserStore interface {
	FindDataWithCondition(context.Context, map[string]any) (*usermodel.User, error)
}

type authUserCache struct {
	store      UserStore // mysql
	cacheStore Cache     // redis
}

func NewAuthUserCache(store UserStore, cacheStore Cache) *authUserCache {
	return &authUserCache{
		store:      store,
		cacheStore: cacheStore,
	}
}

func (c *authUserCache) FindDataWithCondition(ctx context.Context, conditions map[string]any) (*usermodel.User, error) {
	var user usermodel.User

	userId := conditions["id"].(int)
	key := fmt.Sprintf(cacheKey, userId) // key lưu xuống redis

	// Cố gắng tìm data của user từ cache trước
	c.cacheStore.Get(ctx, key, &user)

	// Tìm được data trong cache, return về data đó
	if user.Id > 0 {
		return &user, nil
	}

	// Tìm data user ở dưới DB
	u, err := c.store.FindDataWithCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	// Lưu data user xuống cache
	c.cacheStore.Set(ctx, key, u, time.Hour*2)

	return u, nil
}
