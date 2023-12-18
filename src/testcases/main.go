package main

import (
	"fmt"
	"testing"
)

func TestEven(t *testing.T) {
	n := 2
	expected := "YES"
	res := CheckEven(n)

	if expected != res {
		t.Errorf("Expected %v, Got: %v", expected, res)
	}
}

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
