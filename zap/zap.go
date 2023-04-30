package main

// Here's an example of using zap for logging in Go, with a custom error message:

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Starting the program")
	logger.Debug("Doing some work")
	logger.Info("Finished the work")

	err := someError()
	logger.Error("An error occurred",
		zap.Error(err),
		zap.String("message", err.Error()),
	)
}

func someError() error {
	return &CustomError{message: "Custom error message"}
}

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}
