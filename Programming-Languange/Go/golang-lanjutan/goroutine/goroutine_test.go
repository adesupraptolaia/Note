package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
- func main itu adalah goroutine tersendiri (main goroutine mungkin)
- Cara buat go routine cukup tambahin go didepan functionnya
- gorountine akan dijalankan secara asyncronous
- jika proses go routine belum selesai, tapi proramnya udah selesai, maka otomatis goroutine tsb akan dikill
*/

func RunHelloWorld() {
	fmt.Println("func Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("main goroutine")

	time.Sleep(1 * time.Second)
}

/*
- goroutine sangat ringan
- kita bahkan bisa membuat 1juta goroutine tanpa takut kehabisan memory
*/

func DisplayNumber(number int) {
	fmt.Println("jumlah goroutine", runtime.NumGoroutine(), "number", number)
}

func TestCreateManyGoroutine(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("jumlah goroutine", runtime.NumGoroutine())
}
