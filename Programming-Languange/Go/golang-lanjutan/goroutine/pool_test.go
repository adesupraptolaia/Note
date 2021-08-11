package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
- ini adalah implementasi design pattern "object pool pattern"
- kita bisa menggunakan data di pool, jika sudah selesai, kita kembalikan lagi ke pool nya
- jika suatu waktu kita mendapati pool kosong, maka kita harus menunggu sampai orang lain mengembalikan datanya ke pool
- setelah data dipool ada, maka kita bisa menggunakannya lagi
- implementasi pool di golang sudah aman dari race condition
*/

func TestPool(t *testing.T) {
	var pool sync.Pool
	var wg sync.WaitGroup

	/*
		- jika ingin datanya != nil ketika poolnya kosong
		- ntar data="New" jika poolnya kosong
	*/
	// pool = sync.Pool{
	// 	New: func() interface{} {
	// 		return "New"
	// 	},
	// }

	pool.Put("Ade")
	pool.Put("Suprapto")
	pool.Put("Laia")
	pool.Put("Putra")

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println("data dari pool atas", data)
			pool.Put(data)
			wg.Done()
		}()
	}

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			data := pool.Get()
			for data == nil {
				data = pool.Get()
			}

			fmt.Println("data dari pool bawah", data)

			time.Sleep(1 * time.Second)
			pool.Put(data)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("selesai")
}
