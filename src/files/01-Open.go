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
	file, err := os.Open("sample.txt")
	handleError(err)
	buff := make([]byte, 5)
	for {
		n, err := file.Read(buff)
		fmt.Printf("%s", buff[:n])
		if err != nil {
			break
		}
	}

}
