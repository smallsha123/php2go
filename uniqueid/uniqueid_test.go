package uniqueid

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sync"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	wgg := sync.WaitGroup{}
	lock := sync.Mutex{}
	machineIDNum := 1 //机器数量
	wgg.Add(machineIDNum)
	cache := make(map[string]struct{})

	count := 500 // 每台机器生成的id数量
	re := redis.New("127.0.0.1:6379", func(r *redis.Redis) {
		r.Type = "node"
		r.Pass = "GwP7iMrpwHYX"
	})

	for i := 0; i < machineIDNum; i++ {

		g, err := NewTimeUnique(re)
		if err != nil {
			t.Errorf("错误：%s", err)
		}
		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				lock.Lock()
				cache[g.Generate()] = struct{}{}
				lock.Unlock()
				wg.Done()
			}()

		}

		wg.Wait()

		wgg.Done()
	}

	wgg.Wait()

	if repeatNum := count*machineIDNum - len(cache); repeatNum > 0 {
		t.Errorf("生成的唯一id有重复，总数：%d，重复数:%d", count*machineIDNum, repeatNum)
	}

	time.Sleep(time.Second * 10)

}
