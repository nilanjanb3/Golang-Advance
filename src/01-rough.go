package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-al")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Printf("%s", output)
}
