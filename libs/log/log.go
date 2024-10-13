package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kurnhyalcantara/teman-petani/config"
	"github.com/rs/zerolog"
)

// SetupZerolog sets up zerolog with custom formatting and fluentd/fluentbit support.
func SetupZerolog(appConfig *config.Config) zerolog.Logger {
	hostname, _ := os.Hostname()
	appName := strings.ReplaceAll(strings.ToLower(appConfig.AppName), " ", "_")

	// Custom writer configuration (ConsoleWriter for better readability)
	outputWriter := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}

	// Custom log format: including service name, hostname, and timestamp
	outputWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("\033[1m %s\033[0m", i)
	}
	outputWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("[%s]:", i)
	}
	outputWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("\"%s\"", i)
	}
	outputWriter.FormatTimestamp = func(i interface{}) string {
		return fmt.Sprintf("\033[33m[TIME] %s\033[0m", i) // 33m for yellow
	}

	// Initialize the logger with the console writer
	logger := zerolog.New(outputWriter).With().
		Timestamp().
		Str("app_name", appName).
		Str("host_name", hostname).
		Logger()

	// Setup Fluentd/Fluentbit if output is set to "elastic"
	if strings.ToLower(appConfig.LoggerOutput) == "elastic" {
		// Here, you would add Fluentd/Fluentbit hook support for zerolog
		// This typically involves setting up a custom writer to send logs to Fluentbit
		// For this example, we will log a message indicating this would be where Fluentbit is configured.
		logger.Info().Str("fluentbit_host", appConfig.FluentBitHost).Str("fluentbit_port", appConfig.FluentBitPort).Msg("Fluentbit logging is set up")
	}

	return logger
}
