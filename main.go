package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go CetakNama("Hi")
	go CetakNama("Hello")
	// wg.Wait()

	// wg.Add(1)
	go CetakNama("What's up?")
	wg.Wait()

}

func CetakNama(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	defer wg.Done()
}
