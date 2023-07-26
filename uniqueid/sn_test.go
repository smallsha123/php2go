package uniqueid

import (
	"fmt"
	"sync"
	"testing"
)

func TestGenSn(t *testing.T) {
	//fmt.Printf("%v", -1^(-1<<10))
	var (
		wg  sync.WaitGroup
		ch  = make(chan uint64, 3000)
		mux sync.Mutex
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for j := 0; j < 30; j++ {

				id, _ := flake.NextID()
				mux.Lock()
				ch <- id
				mux.Unlock()
				fmt.Printf("订单ID：%v ,编号：%v \n", id, i*10+j)
			}
		}(i)
	}

	wg.Wait()
	close(ch)

	ids := make([]uint64, 0, 10000)
	for id := range ch {
		ids = append(ids, id)
	}
	res, newId := hasDuplicate(ids)
	fmt.Printf("是否重复：%v,重复的数字:%v", res, newId)
}

func hasDuplicate(s []uint64) (bool, uint64) {
	m := make(map[uint64]bool)
	for _, v := range s {
		if m[v] {
			return true, v
		}
		m[v] = true
	}
	return false, 0
}
