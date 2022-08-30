package golibs

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/tumiao/redis"
	"sync"
	"time"
)

const (
	DefaultExpire = 8 * time.Second
	DefaultTries  = 16
	DefaultDelay  = 512 * time.Millisecond
)

var ErrAcquireFailed = errors.New("failed to acquire lock")

type RedisMutex struct {
	Name        string //Resource name
	Expire      time.Duration
	Delay       time.Duration
	TryCount    int    //获取锁的重试次数
	value       string //往redis里存的随机值
	mutex       sync.Mutex
	redisClient *redis.Client
}

func (m *Model) NewRedisMutex(name string) *RedisMutex {
	if name == "" {
		panic("invalid resource name")
	}
	return &RedisMutex{
		Name:        name,
		Expire:      DefaultExpire,
		Delay:       DefaultDelay,
		TryCount:    DefaultTries,
		redisClient: m.redisClient,
	}
}

//带过期时间的redis锁
func (m *Model) NewRedisMutexWithExpire(name string, expire time.Duration) *RedisMutex {
	if name == "" {
		panic("invalid resource name")
	}
	//如果时间小于8秒的，一律按照8秒处理
	if expire < DefaultExpire {
		return m.NewRedisMutex(name)
	}
	delay := (expire + 2*time.Second) / DefaultTries
	return &RedisMutex{
		Name:        name,
		Expire:      expire,
		Delay:       delay,
		redisClient: m.redisClient,
	}
}

func (m *RedisMutex) Lock() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return err
	}
	value := base64.StdEncoding.EncodeToString(b)
	expire := m.Expire
	if expire == 0 {
		expire = DefaultExpire
	}
	delay := m.Delay
	if delay == 0 {
		delay = DefaultDelay
	}
	tryCount := m.TryCount
	if tryCount == 0 {
		tryCount = DefaultTries
	}
	for i := 0; i < tryCount; i++ {
		cmd := m.redisClient.SetNX(m.Name, value, expire)
		if cmd.Err() == nil && cmd.Val() {
			m.value = value
			return nil
		}
		time.Sleep(delay)
	}
	return ErrAcquireFailed
}

func (m *RedisMutex) Unlock() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	value := m.value
	if value == "" {
		panic("unlock of unlocked mutex")
	}
	m.value = ""
	m.redisClient.Eval(`
	        if redis.call("get", KEYS[1]) == ARGV[1] then
			    return redis.call("del", KEYS[1])
			else
			    return 0
			end
	`, []string{m.Name}, []string{value})
}
