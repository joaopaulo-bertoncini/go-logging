package main

// Here’s an example of using logrus for logging in Go, with a custom error message:

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	log.Info("Starting the program")
	log.Debug("Doing some work")
	log.Info("Finished the work")

	err := someError()
	log.WithError(err).WithFields(log.Fields{
		"message": err.Error(),
	}).Error("An error occurred")
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
