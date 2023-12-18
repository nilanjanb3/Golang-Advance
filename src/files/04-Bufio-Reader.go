package main

import (
	"bufio"
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
	file, err := os.OpenFile("sample.txt", os.O_RDONLY, 0600)
	handleError(err)
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		fmt.Printf("%s\n", line)
	}

}
