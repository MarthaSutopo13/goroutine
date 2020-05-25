package main

import (
	"fmt"
)

func main() {
	go CetakNama("Hi")
	go CetakNama("Hello")

	var i string
	fmt.Scanln(&i)

	fmt.Println(i)
}

func CetakNama(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		// time.Sleep(time.Second)
	}
}
