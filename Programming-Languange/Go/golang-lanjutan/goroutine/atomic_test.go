package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
- atomic adalah package untuk menggunakan data primitive secara aman pada proses concurrent
- data primitive itu seperti int, float, bolean, string
- mutex locking ---> counter++ ---> mutex unlocking ===> ini sebenarnya bisa pakai atomic package
- selebihnya bisa dilihat di
https://golang.org/pkg/sync/atomic/
*/

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	var counter int64

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter", counter)
}
