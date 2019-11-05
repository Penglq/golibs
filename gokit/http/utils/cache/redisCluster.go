package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisCluster struct {
	redisClusterClient *redis.ClusterClient
	Addrs              []string
	Password           string
}

func NewRedisCluster(redis2 *RedisCluster) (*RedisCluster, error) {
	var err error
	clusterRedisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    redis2.Addrs,
		Password: redis2.Password,
	})
	redis2.redisClusterClient = clusterRedisClient
	_, err = clusterRedisClient.Ping().Result()
	return redis2, err
}

func (r *RedisCluster) Get(key string) *redis.StringCmd {
	return r.redisClusterClient.Get(key)
}

func (r *RedisCluster) Set(key string, val interface{}, t time.Duration) *redis.StatusCmd {
	return r.redisClusterClient.Set(key, val, t)
}

func (r *RedisCluster) Exists(key string) *redis.IntCmd {
	return r.redisClusterClient.Exists(key)
}

func (r *RedisCluster) Delete(key string) *redis.IntCmd {
	return r.redisClusterClient.Del(key)
}
