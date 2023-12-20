package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 3
	var str string = strconv.FormatInt(int64(num), 10)
	fmt.Println(str)
	copy, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(copy)
	var f float64 = float64(num)
	fmt.Printf("%0.2f", f)

}
