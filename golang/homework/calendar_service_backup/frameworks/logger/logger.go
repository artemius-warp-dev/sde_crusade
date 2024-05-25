// logger/logger.go
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance *zap.Logger
)

// GetInstance returns the singleton instance of the logger
func GetInstance() *zap.Logger {
	return instance
}

func getAtomicLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// ConfigureLogger configures the logger with the specified level and output paths
func ConfigureLogger(level string, logFilePath string) {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(getAtomicLevel(level))

	config := zap.Config{
		Level:            atomicLevel,
		Encoding:         "json",
		OutputPaths:      []string{"stdout", logFilePath},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.LowercaseLevelEncoder, // Adds the log level to the output
			TimeKey:      "time",                        // Optional: adds a timestamp
			EncodeTime:   zapcore.ISO8601TimeEncoder,    // Optional: formats the timestamp
			CallerKey:    "caller",                      // Optional: adds caller info
			EncodeCaller: zapcore.ShortCallerEncoder,    // Optional: formats the caller info
		},
	}

	var err error
	instance, err = config.Build()
	if err != nil {
		// If logger initialization fails, log the error to stderr
		zap.NewExample().Error("failed to create logger", zap.Error(err))
	}
}

// Debug logs a message at Debug level
func Debug(fields ...interface{}) {
	instance.Sugar().Debug(fields...)
}

// Info logs a message at Info level
func Info(fields ...interface{}) {
	instance.Sugar().Info(fields...)
}

// Warning logs a message at Warning level
func Warning(fields ...interface{}) {
	instance.Sugar().Warn(fields...)
}

// Error logs a message at Error level
func Error(fields ...interface{}) {
	instance.Sugar().Error(fields...)
}

// Fatal logs a message at Fatal level and then calls os.Exit(1)
func Fatal(fields ...interface{}) {
	instance.Sugar().Fatal(fields...)
}
