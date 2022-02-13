package builder

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger instantiates a new zap.SugaredLogger.
func NewLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.TimeKey = "level"
	config.EncoderConfig.MessageKey = "msg"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	// config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.ConsoleSeparator = "  "
	config.Encoding = "console"
	config.DisableCaller = true
	config.DisableStacktrace = true
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Error Creating Logger. %v", err)
	}
	return logger.Sugar()
}

// Sync synchronizes a zap.SugaredLogger.
func Sync(log *zap.SugaredLogger) {
	_ = log.Sync()
}

// LogLevel returns the definition of levels in logs.
func LogLevel(config Configuration) zapcore.Level {
	var lvl zapcore.Level
	err := json.Unmarshal([]byte("\""+config.LogLevel+"\""), &lvl)
	if err != nil {
		return zapcore.InfoLevel
	}
	return lvl
}
