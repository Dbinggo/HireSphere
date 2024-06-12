package myRedis

import (
	"context"
	"fmt"
	"github.com/Dbinggo/HireSphere/server/configs"
	"github.com/Dbinggo/HireSphere/server/global"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"log"
)

const (
	redisAddr = "%s:%d"
)

func InitMyRedis() {
	if configs.Conf.RedisConfig.Enable {
		logrus.Info("do not need redis")
		return
	}
	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               fmt.Sprintf(redisAddr, configs.Conf.RedisConfig.Host, configs.Conf.RedisConfig.Port),
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           configs.Conf.RedisConfig.Password,
		DB:                 configs.Conf.RedisConfig.DB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolFIFO:           false,
		PoolSize:           1000,
		MinIdleConns:       1,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		logrus.Fatal("redis cant connect")
	} else {
		log.Println("init redis")
		global.Redis = client
	}
}
