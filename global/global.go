package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/lliuhuan/arco-design-pro-gin/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	AdpDb                 *gorm.DB
	AdpVp                 *viper.Viper
	AdpConfig             config.Server
	AdpLog                *zap.Logger
	AdpConcurrencyControl = &singleflight.Group{}
	AdpValidator          ut.Translator
	AdpRedis              *redis.Client

	BlackCache local_cache.Cache
)
