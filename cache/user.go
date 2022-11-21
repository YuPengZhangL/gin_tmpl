package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"gg_web_tmpl/common/log"
	"gg_web_tmpl/model"
	"github.com/go-redis/redis/v8"
)

func getUserKey(id int64) string {
	return fmt.Sprintf("gg_web_tmpl.user.id.%v", id)
}

func SetUser(ctx context.Context, client *redis.Client, user *model.User) error {
	logger := log.GetLoggerWithCtx(ctx)
	str, err := json.Marshal(user)
	if err != nil {
		logger.Errorf("set user json marshal err: %v", err)
		return err
	}
	err = client.Set(ctx, getUserKey(user.ID), str, 0).Err()
	if err != nil {
		logger.Errorf("set user redis err: %v", err)
		return err
	}
	return nil
}

func GetUser(ctx context.Context, client *redis.Client, id int64) (*model.User, error) {
	logger := log.GetLoggerWithCtx(ctx)
	result, err := client.Get(ctx, getUserKey(id)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		logger.Errorf("get user redis err: %v", err)
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(result), user)
	if err != nil {
		logger.Errorf("set user json unmarshal err: %v", err)
		return nil, err
	}
	return user, nil
}
