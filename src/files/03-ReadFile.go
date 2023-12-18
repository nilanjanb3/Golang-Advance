package main

import (
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
func main() {
	var err error
	op, err := os.ReadFile("sample.txt")
	handleError(err)
	fmt.Printf("%s", op)

}
