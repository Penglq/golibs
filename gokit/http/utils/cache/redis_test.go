package cache

import (
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
	r, _ := NewRedis(&Redis{
		Addr:     "127.0.0.1:6379",
		Password: "pwdseek",
	})
	err := r.Set("hello", "a", time.Second)
	s := r.Get("hello")
	t.Log(s, err)
}
