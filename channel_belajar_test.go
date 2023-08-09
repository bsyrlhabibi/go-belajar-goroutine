package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponseSeratus(ch chan<- string, message string) {
	time.Sleep(1 * time.Second)
	ch <- message
}

func GenerateData(ch chan<- string, numData int) {
	for i := 0; i < numData; i++ {
		ch <- fmt.Sprintf("Data %d", i+1)
	}
	close(ch)
}

func TestChannelSeratus(t *testing.T) {
	numChannels := 100
	channels := make([]chan string, numChannels)

	// Inisialisasi dan tutup semua channel saat selesai
	for i := 0; i < numChannels; i++ {
		channels[i] = make(chan string)
		defer close(channels[i])
	}

	numData := 200

	// Fungsi goroutine untuk mengisi setiap channel dengan data
	for i := 0; i < numChannels; i++ {
		go GiveMeResponseSeratus(channels[i], fmt.Sprintf("Data dari channel %d", i+1))
	}

	// Fungsi goroutine untuk menghasilkan data dan mengirimkannya ke 100 channel
	go func() {
		for i := 0; i < numData; i++ {
			for j := 0; j < numChannels; j++ {
				channels[j] <- fmt.Sprintf("Data %d dari channel %d", i+1, j+1)
			}
		}
	}()

	counter := 0
	for counter < numChannels*numData {
		// Menerima data dari setiap channel yang siap
		for i := 0; i < numChannels; i++ {
			data, ok := <-channels[i]
			if ok {
				fmt.Println(data)
				counter++
			}
		}
	}
	fmt.Println()
}

func TestChannelDong(t *testing.T) {
	numChannels := 100
	channels := make([]chan string, numChannels)

	// Inisialisasi dan tutup semua channel saat selesai
	for i := 0; i < numChannels; i++ {
		channels[i] = make(chan string)
		defer close(channels[i])
	}

	// Fungsi goroutine untuk mengisi setiap channel dengan data
	for i := 0; i < numChannels; i++ {
		go GiveMeResponseSeratus(channels[i], fmt.Sprintf("Data dari channel %d", i+1))
	}

	counter := 0
	for counter < numChannels {
		// Menerima data dari setiap channel yang siap
		for i := 0; i < numChannels; i++ {
			data, ok := <-channels[i]
			if ok {
				fmt.Println(data)
				counter++
			}
		}
	}
}
