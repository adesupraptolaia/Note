package goroutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
- ticker ini untuk kejadian berulang setiap waktu yang diberikan
- misalnya kejadian berulang setiap 1 detik, dsb
*/

func TestTicker(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ticker := time.NewTicker(1 * time.Second)

	go func(ctx context.Context) {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
			case <-ctx.Done():
				ticker.Stop()
				fmt.Println("done")
				return
			}
		}
	}(ctx)

	time.Sleep(15 * time.Second)
	cancel()
}

func TestTickAja(t *testing.T) {
	tick := time.Tick(1 * time.Second)
	timer := time.After(10 * time.Second)

loop:
	for {
		select {
		case <-tick:
			fmt.Println(time.Now())
		case <-timer:
			break loop
		}
	}

	fmt.Println("selesai")
}
