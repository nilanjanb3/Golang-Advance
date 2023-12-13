package main

import (
	"fmt"
	"time"
)

func buy(ch chan string) {
	fmt.Println("Receiving...")
	product := <-ch
	fmt.Printf("%v Received", product)

}

func sell(ch chan string) {
	product := "Ice Cream"
	ch <- product
	fmt.Println("Sent Product")
}
func main() {
	ch := make(chan string)
	go sell(ch)
	go buy(ch)

	time.Sleep(2 * time.Second)
}
