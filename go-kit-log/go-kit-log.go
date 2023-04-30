package main

// Here’s an example of using go-kit/log for logging in Go:

import (
	"os"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestamp)

	logger.Log("msg", "Starting the program")
	logger.Log("msg", "Doing some work")
	logger.Log("msg", "Finished the work")

	logger.Log("error", "some error")
}
