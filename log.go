package main

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DefaultLogFilePath = "/var/log/test-cni.log"
)

//goland:noinspection ALL
type Logger interface {
	Infof(format string, args ...any)
	Errorf(format string, any ...any)
	Warnf(format string, any ...any)
	Debugf(format string, any ...any)
	Fatalf(format string, any ...any)
	SetLvl(level int8)
}

func NewLogger(filePath string) (Logger, error) {
	var (
		outputPaths = []string{"stdout"}
	)
	filePath = strings.TrimSpace(filePath)
	if filePath != "" {
		outputPaths = append(outputPaths, filePath)
	}

	atomicLvl := zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg := zap.Config{
		Level:       atomicLvl,
		Development: false,
		Sampling:    nil,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			StacktraceKey:  "",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   nil,
		},
		OutputPaths:      outputPaths,
		ErrorOutputPaths: outputPaths,
	}

	base, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("create new logger: %w", err)
	}

	return &zapLoggerWrapper{
		SugaredLogger: base.Sugar(),
		AtomicLevel:   &atomicLvl,
	}, nil
}

type zapLoggerWrapper struct {
	*zap.SugaredLogger
	*zap.AtomicLevel
	initLvl zapcore.Level
}

func (lw *zapLoggerWrapper) SetLvl(level int8) {
	lw.AtomicLevel.SetLevel(zapcore.Level(level))
}
