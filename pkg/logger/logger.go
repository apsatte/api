package logger

import (
	"api/pkg/configuration"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(cfg *configuration.Logger) (*zap.Logger, error) {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "timestamp",
			LevelKey:     "level",
			MessageKey:   "message",
			CallerKey:    "caller",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.FullCallerEncoder,
		},
		OutputPaths:      []string{cfg.OutputPaths},
		ErrorOutputPaths: []string{cfg.ErrorOutputPaths},
		InitialFields: map[string]interface{}{
			"service": "backend",
		},
	}

	return config.Build()
}
