package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(2 * time.Second)
			wg.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("jumlah CPU", totalCPU)

	// kalo diatas 0, maka akan merubah jumlah thread
	// kalo dibawah 0, maka akan menampilkan yang existing atau tidak merubah jumlah thread
	totalTread := runtime.GOMAXPROCS(-1)
	fmt.Println("jumlah thread", totalTread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("jumlah goroutine", totalGoroutine)

	wg.Wait()

	runtime.GOMAXPROCS(3)
	totalTread = runtime.GOMAXPROCS(-1)
	fmt.Println("jumlah thread setelah diubah", totalTread)
}
