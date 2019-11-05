package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	redisClient *redis.Client
	Addr        string
	Password    string
}

func NewRedis(redis2 *Redis) (*Redis, error) {
	var err error
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redis2.Addr,
		Password: redis2.Password,
	})
	redis2.redisClient = redisClient
	_, err = redisClient.Ping().Result()
	return redis2, err
}

func (r *Redis) Get(key string)*redis.StringCmd {
	return r.redisClient.Get(key)
}

func (r *Redis) Set(key string, val interface{}, t time.Duration) *redis.StatusCmd {
	return r.redisClient.Set(key, val, t)
}

func (r *Redis) Exists(key string) *redis.IntCmd {
	return r.redisClient.Exists(key)
}

func (r *Redis) Delete(key string) *redis.IntCmd {
	return r.redisClient.Del(key)
}
