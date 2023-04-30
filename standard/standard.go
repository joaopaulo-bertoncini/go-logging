package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	logger.Println("Starting the program")
	logger.Println("Doing some work")
	logger.Println("Finished the work")
	logger.Fatalln("An error occurred")
}
