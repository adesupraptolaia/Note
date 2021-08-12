package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
- Timer ini seperti jam weker, sebagai pengingat
- Timer ini intinya bakal balikin struc Timer yang isinya ada chanel
- yang mana channel ini akan ke trigger ketika duration dari timer habis
*/

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

/*
- kadang kita butuh channel nya saja tanpa butuh data timer nya
- oleh karena itu kita bisa pakai time.After
*/

func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	tick := <-channel
	fmt.Println(tick)
}

/*
- time.AfterFunc akan mengeksekusi fungsi setelah durasi yang diberikan
*/

func TestTimerAfterFunc(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println(time.Now().Format("15:04:05"))

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("setelah 5 detik", time.Now().Format("15:04:05"))
		wg.Done()
	})

	fmt.Println("atas")
	wg.Wait()
	fmt.Println("bawah")
}
