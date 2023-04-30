package main

// Here’s an example of using zerolog for logging in Go with a custom error message:

import (
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().Msg("Starting the program")
	logger.Debug().Msg("Doing some work")
	logger.Info().Msg("Finished the work")

	errorMessage := "some error"
	logger.Error().Str("error", errorMessage).Msg("An error occurred")
}
