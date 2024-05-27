package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/nilemarezz/go-init-template/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(config *config.Config) error {
	// Define logs directory path
	logsDir := filepath.Join(config.Log.Path)

	// Check if logs directory exists
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		// Create logs directory
		if err := os.Mkdir(logsDir, 0755); err != nil {
			return fmt.Errorf("failed to create logs directory: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to check logs directory: %v", err)
	}

	// Define log file name based on current date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	logFilePath := filepath.Join(logsDir, logFileName)

	// Create a log file
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	// Configure log encoding
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	// Configure log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.InfoLevel)

	// Create a Zap encoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	// Create a Zap core for writing to the file
	fileCore := zapcore.NewCore(encoder, zapcore.AddSync(logFile), atomicLevel)

	// Create a Zap core for writing to the console
	consoleCore := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), atomicLevel)

	// Combine both cores
	multiCore := zapcore.NewTee(fileCore, consoleCore)

	// Create a Zap logger
	Logger = zap.New(multiCore, zap.AddCaller())

	return nil
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Warning(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func InitTestLogger() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		panic("Failed to initialize logger for tests: " + err.Error())
	}
}
