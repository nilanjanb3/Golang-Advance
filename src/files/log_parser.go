package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.OpenFile("access.log", os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(file.Read())
	// buff := make([]byte, 4)
	// fmt.Println("Debugging")
	// n, err := file.Read(buff)
	// fmt.Println(n, err)

	// buf, _ := os.ReadFile("access.log")
	// logs := string(buf[:])
	// lines := strings.Fields(logs)
	// for _, line := range lines {
	// 	fmt.Println(line)
	// 	fmt.Print("✅")
	// }

	// fmt.Print(string(buf[:]))
	// for {
	// 	n, err := file.Read(buff)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		fmt.Print(string(buff[:n]))
	// 	}
	// }
	// Let's try
	file, err := os.OpenFile("access.log", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
