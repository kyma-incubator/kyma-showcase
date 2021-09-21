package logging

import (
	"fmt"
	"github.com/kyma-project/kyma/common/logging/logger"
	"os"
)

func InitLogger() *logger.Logger {
	level, err := logger.MapLevel("debug")
	if err != nil {
		if logErr := logger.LogFatalError("Failed to map log level from options: %s", err.Error()); logErr != nil {
			fmt.Printf("Failed to initializie default fatal error logging: %s, Failed to map log level from options: %s", logErr, err)
		}
		os.Exit(2)
	}

	format, err := logger.MapFormat("text")
	if err != nil {
		if logErr := logger.LogFatalError("Failed to map log format from options: %s", err.Error()); logErr != nil {
			fmt.Printf("Failed to initializie default fatal error logging: %s, Failed to map log format from options: %s", logErr, err)
		}
		os.Exit(3)
	}

	log, err := logger.New(format, level)
	if err != nil {
		if logErr := logger.LogFatalError("Failed to initialize logging: %s", err.Error()); logErr != nil {
			fmt.Printf("Failed to initializie default fatal error logging: %s, Failed to initialize logging: %s", logErr, err)
		}
		os.Exit(4)
	}
	if err := logger.InitKlog(log, level); err != nil {
		log.WithContext().Errorf("While initializing klog logging: %s", err.Error())
		os.Exit(5)
	}
	return log
}
