package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
- sync.Map mirip seperti map
- hanya saja map ini aman dari concurrent menggunakan goroutine
- tidak ada race condition

Ada beberapa function yang bisa kita gunakan di sync.Map :
- Store(key, value) => store data key-value
- Load(key) => get val by key
- Delete(key) => delete data key-value
- Range(function(key, value)) => untuk range/loop
*/

type userUsingSyncMap struct {
	wg   sync.WaitGroup
	data sync.Map
}

func (u *userUsingSyncMap) addToMap(value int) {
	defer u.wg.Done()
	u.data.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	var u userUsingSyncMap

	for i := 0; i < 100; i++ {
		u.wg.Add(1)
		go u.addToMap(i)
	}

	u.wg.Wait()

	var counter int
	u.data.Range(func(key, val interface{}) bool {
		fmt.Println(key, val)
		counter++
		return true
	})

	fmt.Println("counter = ", counter)
}
