package logger

import (
	"context"
	"errors"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type loader struct {
	*confita.Loader
}

func NewConfig(configPath string) loader {
	loader1 := confita.NewLoader(
		env.NewBackend(),
		file.NewBackend(configPath),
		flags.NewBackend(),
	)
	return loader{loader1}
}

func (l *loader) GetConfig(ctx context.Context, dst interface{}) error {
	err := l.Load(ctx, dst)
	if err != nil {
		return err
	}
	return nil
}

type Logger struct {
	logger *zap.Logger
}

func NewLogger(logLevel string, output []string) (*zap.Logger, error) {
	if logLevel == "Debug" || logLevel == "debug" {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.RFC3339TimeEncoder
		config.ConsoleSeparator = " | "
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		defaultLogLevel := zapcore.DebugLevel
		core := zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)

		logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

		return logger, nil
	}
	if logLevel == "release" || logLevel == "Release" {
		logger, err := zap.Config{
			Level:       zap.NewAtomicLevelAt(zapcore.ErrorLevel),
			OutputPaths: output,
			Encoding:    "console",
			EncoderConfig: zapcore.EncoderConfig{
				MessageKey: "message", // <--
			},
		}.Build()
		return logger, err
	}
	return nil, errors.New("log level not set")
}
