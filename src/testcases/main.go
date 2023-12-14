package main

import "fmt"

func CheckEven(n int) string {
	if n%2 == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	n := 4
	fmt.Println(CheckEven(n))
}
