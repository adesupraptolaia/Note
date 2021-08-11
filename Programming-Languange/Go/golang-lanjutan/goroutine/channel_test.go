package goroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
### Channel ###
- channel itu tempat komunikasi secara syncronous antar goroutine
- channel ini dapat menjadi pengirim dan penerima, biasanya pengirim dan penerima-nya adalah goroutine yang berbeda
- saat ngirim data ke channel maka goroutine akan terblok sampai ada goroutine yang menerima data dari channel tsb
- oleh karena itulah channel disebut sebagai alat komunikasi syncronous (blocking)
- ini bisa menjadi alternatif seperti mekanisme "async await" di bahasa pemograman lain
- channel defaultnya cuma bisa nampung 1 data, jika mau masukin data lain maka harus nunggu data existing diambil dulu oleh goroutine
- channel cuma bisa diisi dengan 1 jenis data
- channel bisa dikonsumsi oleh lebih dari 1 goroutine
- channel harus diclose agar tidak ada memory leak
- closed channel akan menyebakan channel tidak bisa menerima data lagi
- tapi closed channel masih bisa untuk mengirim data yang masih ada didalamnya (lihat TestBufferedChannel)
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		fmt.Println("atas")
		time.Sleep(2 * time.Second)
		fmt.Println("bawah")

		channel <- "Ade"
		fmt.Println("data dikirim")
	}()

	data := <-channel
	fmt.Println("data diterima", data)
}

/*
### Channel As Parameter ###
- jika dijadikan sebagai parameter sebuah fungsi, channel secara default adalah pass by reference
*/

func GiveResponse(channel chan string) {
	fmt.Println("atas")
	time.Sleep(2 * time.Second)
	fmt.Println("bawah")

	channel <- "Ade"
	fmt.Println("data dikirim")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveResponse(channel)

	data := <-channel
	fmt.Println("data diterima", data)
}

/*
### Channel In/Out As Parameter ###

- ini hanya bisa dilakukan sebagai parameter function
- kita bisa mendefenisikan channel hanya sebagai penerima atau pengirim saja
- ch<- : artinya channel nenerima data
- <-ch : artinya cheannel mengirim data
*/

func OnlyIn(channel chan<- string) {
	fmt.Println("atas OnlyIn")
	time.Sleep(2 * time.Second)
	fmt.Println("bawah OnlyIn")

	channel <- "Ade"
	fmt.Println("data dikirim")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println("data diterima", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

/*
### Buffered Channel ###

- ada kasus dimana kita ingin data antrian di channel bisa lebih dari 1
- untuk itu kita bisa gunakan buffered channel
- make(chan string, 2) ==> artinya kita bisa nambahin data sebanyak 1+2=3 secara bersamaan didalam channel

### Range Channel ###
- gunakan for range channel
- perulangan akan berhenti jika tidak ada data lagi di dalamn channel dan channel telah di close
- tetapi, jika channel telah diclose tetapi masih ada data didalamnya, maka perulangan akan terus berlanjut sampai data di dalam channel benar-benar habis
*/

func write(channel chan<- int) {
	for i := 1; i <= 6; i++ {
		channel <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(channel)
	fmt.Println("channel closed")
}

func TestBufferedChannelAndRange(t *testing.T) {
	channel := make(chan int, 3)

	go write(channel)

	for data := range channel {
		fmt.Println("recieved data", data, "from channel")
		time.Sleep(2 * time.Second)
	}

	fmt.Println("END")
}

/*
### SELECT Channel ###
- select digunakan ketika ada beberapa channel
- kita mau menerapkan perlakuan berbeda untuk tiap-tiap channel tsb
- select statement akan finish jika telah mendapatkan salah satu case
- bisa pakai for-infine-loop jika ada kebutuhan khusus, dengan catatan untuk memberi case break-loop tsb

### Default Statement in Select ###
- default akan true ketika channel-channel diatasnya belum mempunyai nilai
*/

func SendViaPos(channel chan<- string) {
	fmt.Println("atas SendViaPos")
	time.Sleep(3 * time.Second)
	fmt.Println("bawah SendViaPos")

	channel <- "pos"
	fmt.Println("data dikirim dari Pos")
}

func SendViaCargo(channel chan<- string) {
	fmt.Println("atas SendViaCargo")
	time.Sleep(6 * time.Second)
	fmt.Println("bawah SendViaCargo")

	channel <- "cargo"
	fmt.Println("data dikirim dari Cargo")
}

func TestSelectChannel(t *testing.T) {
	channelPos := make(chan string)
	channelCargo := make(chan string)
	defer close(channelPos)
	defer close(channelCargo)

	go SendViaPos(channelPos)
	go SendViaCargo(channelCargo)

	var counter int
	for {
		select {
		case pos := <-channelPos:
			fmt.Println("data diterima dari Pos - ", pos)
			counter++
		case cargo := <-channelCargo:
			fmt.Println("data diterima dari Cargo - ", cargo)
			counter++
		}

		if counter == 2 {
			fmt.Println("END")
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go func() {
		channel1 <- "Channel 1"
	}()

	go func() {
		channel2 <- "Channel 2"
	}()

	var counter int
	for {
		select {
		case data := <-channel1:
			fmt.Println("diterima data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("diterima data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
