package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger(appName string) {
	// Get current date for file name
	currentDate := time.Now().Format("2006-01-02")

	// Configure lumberjack for log rotation
	ljLogger := &lumberjack.Logger{
		Filename:   "logs/" + appName + "/log-" + currentDate + ".log",
		MaxSize:    1,    // Max size in MB before rotation
		MaxBackups: 3,    // Max number of old log files to keep
		MaxAge:     2,    // Max age in days to keep log files
		Compress:   true, // Compress old log files
	}

	// Custom encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // Lowercase level names
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // Human-readable timestamp format
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// Create a custom pretty-print encoder
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// Use a custom console encoder for pretty-printing in the console
	consoleEncoderConfig := encoderConfig
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Capitalize and colorize level names
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// Configure zap core to write to both file and console
	core := zapcore.NewTee(
		zapcore.NewCore(
			jsonEncoder,
			zapcore.AddSync(ljLogger),
			zap.InfoLevel,
		),
		zapcore.NewCore(
			consoleEncoder,
			zapcore.AddSync(os.Stdout),
			zap.InfoLevel,
		),
	)

	// Create the logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}
