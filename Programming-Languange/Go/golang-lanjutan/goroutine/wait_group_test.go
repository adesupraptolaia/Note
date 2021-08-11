package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
- wait group adalah feature di golang agar bisa menunggu proses seluruh goroutine selesai
- untuk menandai proses goroutine yang mau kita tunggu, kita cukup menggunakan method Add(int)
- kemudia apabila goroutine tsb sudah selesai, kita gunakan method Done()
- Untuk menunggu semua proses selesai, gunakan method Wait()
*/

type user struct {
	sync.Mutex
	sync.WaitGroup
	number int
}

func (u *user) RunAsyncronous() {
	defer u.WaitGroup.Done()

	for i := 0; i < 1000; i++ {
		u.Mutex.Lock()
		u.number++
		u.Mutex.Unlock()
	}
}

func TestWaitGroup(t *testing.T) {
	var u user

	for i := 0; i < 1000; i++ {
		u.WaitGroup.Add(1)
		go u.RunAsyncronous()
	}

	u.WaitGroup.Wait()
	fmt.Println("u.number", u.number)
}
