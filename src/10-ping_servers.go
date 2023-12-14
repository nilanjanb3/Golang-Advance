// /usr/bin/true;  exec  /usr/local/go/bin/go run "$0" "$@" ; exit
package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func pingServer(server string, wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command("ping", "-c", "4", server)
	output, err := cmd.CombinedOutput()

	fmt.Printf("Ping results for %s:\n", server)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(output))
	}
}

func main() {
	servers := []string{"google.com", "example.com", "localhost"}

	var wg sync.WaitGroup

	for _, server := range servers {
		wg.Add(1)
		go pingServer(server, &wg)
	}

	wg.Wait()
}
