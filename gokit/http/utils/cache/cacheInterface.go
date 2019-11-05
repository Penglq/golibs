package cache

import (
	"github.com/go-redis/redis"
	"github.com/penglq/QLog"
	"time"
)

var (
	RedisType        CacheType = "redis"
	RedisClusterType CacheType = "redis_cluster"
)

type CacheType string

var cache Cache

type Cache interface {
	Get(string) *redis.StringCmd
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	Exists(string) *redis.IntCmd
	Delete(string) *redis.IntCmd
}

func NewCache(cluster bool, addrs []string, pwd string) {
	var err error
	if cluster {
		cache, err = NewRedisCluster(&RedisCluster{
			Addrs:    addrs,
			Password: pwd,
		})
	} else {
		cache, err = NewRedis(&Redis{
			Addr:     addrs[0],
			Password: pwd,
		})
	}
	if err != nil {
		QLog.GetLogger().Alert("action", "连接redis出错")
		panic(-1)
	}
}

func GetCache() Cache {
	return cache
}
