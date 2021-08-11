package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
- cond adalah locking berbasis kondisi
- cond butuh Loker (bisa diisi Mutex atau RWMutex)
- func Wait() ==> agar goroutine yang sedang nge-lock tsb untuk nunggu (iddle) di posisi func Wait()
- func Signal() ===> untuk memberitahu hanya ke goroutine yg sdg lock untuk lanjut dari posisi Wait()
- func Broadcast() ===> sama seperti signal, tapi kalo broadcast berlaku untuk semua goroutine
*/

var cond = sync.NewCond(&sync.Mutex{})
var wg = sync.WaitGroup{}

func WaitCondition(value int) {
	cond.L.Lock()

	fmt.Println("atas wait")
	cond.Wait()

	fmt.Println("Done", value)
	cond.L.Unlock()
	wg.Done()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WaitCondition(i)
	}

	// using signal
	/*
	 */
	for i := 0; i < 10; i++ {
		fmt.Println("tunggu sleep signal", i)
		time.Sleep(1 * time.Second)
		fmt.Println("kirim signal", i)
		cond.Signal()
	}

	fmt.Println("tunggu sleep broadcast")
	time.Sleep(5 * time.Second)
	fmt.Println("kirim broadcast")
	cond.Broadcast()

	// using broadcast

	wg.Wait()
}
