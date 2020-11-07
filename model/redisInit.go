package model

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"time"
)

type Redis struct {
	Self	*redis.Pool
}

var RedisDb *Redis

//连接redis
func (rdb *Redis) Init {
	pool := &redis.Pool{
		// 初始化链接数量
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(viper.GetString("redis.network"), viper.GetString("redis.addr"))
		},
	}

	RedisDb = &Redis{Self: pool}
}

//关闭redis
func (rdb *Redis) Close() error {
	if err := RedisDb.Self.Close(); err != nil {
		return err
	}
	return nil
}