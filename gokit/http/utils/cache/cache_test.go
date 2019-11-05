package cache

import (
	"testing"
	"time"
)

func TestGetCache(t *testing.T) {
	NewCache(false, []string{"127.0.0.1:6379"}, "pwdseek")
	err := GetCache().Set("hello", "a", time.Second)
	if err != nil {
		t.Fatal(err)
	}
	a, err := GetCache().Get("hello")
	if a != "a" {
		t.Log(a, "a")
		t.Fail()
	}
	if err != nil {
		t.Fatal(err)
	}

	NewCache(true, []string{"10.141.6.87:6379", "10.141.6.88:6380", "10.141.6.89:6381"}, "wR68543q71")
	err = GetCache().Set("hello", "b", time.Second)
	if err != nil {
		t.Fatal(err)
	}
	b, err := GetCache().Get("hello")
	if b != "b" {
		t.Log(b, "b")
		t.Fail()
	}
	if err != nil {
		t.Fatal(err)
	}
}
