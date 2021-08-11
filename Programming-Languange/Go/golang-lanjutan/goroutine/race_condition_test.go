package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
- kalo saya ngartiiin race condition itu seperti kejadian ketika sebuah data diupdate oleh lebih dari 1 pengubah
- misalnya ketika kita update nilai x = 1
- katakan ada si A dan si B mau ngupdate si X, masing-masing x+1
- tapi ternyata diupdate diwaktu bersamaan
- ketika A update nilai x masih 1 ==> x+1=2
- ketika B update nilai x masih juga 1 ==> x+1=2
- akibatnya nilai akhir x bukanlah 3, melainkan hanya 2

---- inilah contoh kejadian yang dimaksud race condition
*/

func TestRaceCondition(t *testing.T) {
	var wg sync.WaitGroup
	var x int

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("x", x)
}
