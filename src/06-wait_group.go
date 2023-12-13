package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(n int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println(n * n)
	wg.Done()

}
func main() {

	startTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go calculateSquare(i, &wg)

	}
	endTime := time.Now()
	time.Sleep(2 * time.Second)
	fmt.Println(endTime.Sub(startTime))
}
