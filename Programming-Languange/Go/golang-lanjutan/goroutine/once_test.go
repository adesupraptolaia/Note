package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

/*
- sync.once adalah feature di golang sehingga suatu fungsi hanya bisa dieksekusi sekali
- meskipun banyak goroutine mengakses fungsi tersebut, tetapi hanya yang pertama saja yang boleh mengeksekusinya
- selainnya tidak boleh
*/

func showNumber() {
	fmt.Println(1000000)
}

func TestOnce(t *testing.T) {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			once.Do(showNumber)
			wg.Done()
		}()
	}

	wg.Wait()
}
