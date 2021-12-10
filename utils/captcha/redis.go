// Package captcha
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 13:53
package captcha

import (
	"context"
	"time"

	"github.com/mojocn/base64Captcha"

	"github.com/lliuhuan/arco-design-pro-gin/global"

	"go.uber.org/zap"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.AdpRedis.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.AdpLog.Error("RedisStoreSetError!", zap.Error(err))
	}
	return err
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.AdpRedis.Get(rs.Context, key).Result()
	if err != nil {
		global.AdpLog.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := global.AdpRedis.Del(rs.Context, key).Err()
		if err != nil {
			global.AdpLog.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
