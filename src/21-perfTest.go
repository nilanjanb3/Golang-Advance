package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	for i := 0; i < 10000; i++ {
		fmt.Println(i * i)
	}

	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))
}
