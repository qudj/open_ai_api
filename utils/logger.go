package utils

import (
	"github.com/qudj/open_ai_api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger      = initDefaultLogger()
	sugarLogger = logger.Sugar()
	loggers     = make([]*zap.Logger, 0)
)

func LoggerWithFields(fields ...zap.Field) *zap.Logger {
	l := logger.With(fields...)
	loggers = append(loggers, l)
	return logger
}

func SugarLogger() *zap.SugaredLogger {
	return sugarLogger
}

func CloseLogger() (err error) {
	for _, l := range loggers {
		if err = l.Sync(); err != nil {
			return
		}
	}
	if err = logger.Sync(); err != nil {
		return
	}
	if err = sugarLogger.Sync(); err != nil {
		return
	}
	return
}

func initDefaultLogger() (defaultLogger *zap.Logger) {
	if config.Global.Debug {
		defaultLogger, _ = zap.NewDevelopment()
	} else {
		conf := zap.Config{
			EncoderConfig: zapcore.EncoderConfig{
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				TimeKey:        "time",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				FunctionKey:    zapcore.OmitKey,
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
			Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
			Development:      false,
			Encoding:         "json",
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}
		defaultLogger, _ = conf.Build()
	}
	return
}
