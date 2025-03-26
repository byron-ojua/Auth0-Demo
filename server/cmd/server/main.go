package main

import (
	"auth0_demo/internal/api"
	"auth0_demo/internal/config"
	"auth0_demo/pkg/log"
	"flag"
	"os"

	"go.uber.org/zap"
)

func main() {
	// parse the command line flags
	debug := flag.Bool("debug", false, "enable debug mode")
	logFile := flag.String("log-file", "", "log file")
	flag.Parse()

	logOptions := log.LoggerOptions{
		Level:      zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputFile: *logFile,
	}
	if *debug {
		logOptions.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	// Set up the logger
	logger, err := log.New(logOptions)
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Get app configuration
	cfg, err := config.New(logger)
	if err != nil {
		panic(err)
	}

	// Setup API
	apiInstance, err := api.New(logger, cfg)
	if err != nil {
		panic(err)
	}

	// Run the API
	// err = apiInstance.RunLocal()

	// Check for Azure Functions environment
	if port, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); exists {
		// Running in Azure Functions environment
		err = apiInstance.RunAzureFunction(port)
	} else {
		// Running locally
		err = apiInstance.RunLocal()
	}

	if err != nil {
		panic(err)
	}
}
