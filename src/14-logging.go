package main

import "log"

func main() {

	logger := log.Default()

	logger.SetPrefix("INFO: ")

	logger.Printf("Logging...")
}
