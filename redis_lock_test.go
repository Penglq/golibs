package model

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestRedisLock(t *testing.T) {
	OneTestScope(func(m *Model) {
		var wg sync.WaitGroup
		type Share struct {
			cnt int
		}
		var share Share
		//without lock
		for i := 0; i < 3000; i++ {
			wg.Add(1)
			go func(no int) {
				defer wg.Add(-1)
				share.cnt += 1
				share.cnt -= 1
			}(i)
		}
		wg.Wait()
		fmt.Println("at last " + strconv.Itoa(share.cnt))
		time.Sleep(time.Second)
		share.cnt = 0

		//with lock
		for i := 0; i < 3000; i++ {
			wg.Add(1)
			go func(no int) {
				defer wg.Add(-1)
				lock := m.NewRedisMutex("test_lock")
				for {
					err := lock.Lock()
					if err != nil {
						return
					} else {
						break
					}
				}
				defer lock.Unlock()
				share.cnt += 1
				if share.cnt > 1 {
					return
				}
				share.cnt -= 1
			}(i)
		}
		wg.Wait()
		if share.cnt != 0 {
			t.Fatal(share.cnt)
		}
	})
}
