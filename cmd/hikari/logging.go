package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/olehvolynets/hikari"
)

const logLevelEnvKey = "HIKARI_LOG_LEVEL"

var ErrUnknownLogLevel = hikari.HikariError{Err: errors.New("unknown log level")}

func init() {
	slog.SetLogLoggerLevel(slog.LevelWarn)

	logLevel, ok := os.LookupEnv(logLevelEnvKey)
	if ok {
		switch logLevel {
		case "debug":
			slog.SetLogLoggerLevel(slog.LevelDebug)
		case "info":
			slog.SetLogLoggerLevel(slog.LevelInfo)
		case "warn":
			// Default level.
		case "error":
			slog.SetLogLoggerLevel(slog.LevelError)
		default:
			slog.Warn(ErrUnknownLogLevel.Error())
		}
	}
}