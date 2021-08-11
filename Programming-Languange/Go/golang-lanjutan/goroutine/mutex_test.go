package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
- untuk mengatasi masalah race condition di golang karena tiap goroutine mengakses resource yang sama
- maka di golang menyediakan struct mutex
- gunanya agar bisa di locking dan unlocking
- jadi ketika di lock, maka hanya 1 goroutine yang bisa akses line tsb sampai di unlock oleh goroutine tsb
*/

func TestMutex(t *testing.T) {
	var x int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("x", x)
}

/*
- RWMutex berguna untuk memisahkan lock antara Write dan Read
- jadi kalo yang dilock adalah write, maka read masih bisa akses dan ngelock juga
*/

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	fmt.Println("lock write")
	account.Balance = account.Balance + amount
	fmt.Println("ulock write", account.Balance)
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	fmt.Println("lock read")
	balance := account.Balance
	fmt.Println("ulock read", account.Balance)
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				account.AddBalance(1)
				account.GetBalance()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Total Balance", account.GetBalance())
}

/*
- deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun goroutine yang bisa jalan
- pada kejadian dibawah
- user1 dan user2 sama-sama dilock
- sehingga mereka saling menunggu user1 dan user2 unlock dulu, tapi karena 22nya sedang di lock, maka jadinya prosesnya idle
- inilah disebut deadlock
*/

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int, event string) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name, event)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)
	fmt.Println("after 1 minutes")

	user2.Lock()
	fmt.Println("Lock user2", user2.Name, event)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
	fmt.Println("unlock", event)
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Eko",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Budi",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000, "atas")
	go Transfer(&user2, &user1, 200000, "bawah")

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}
