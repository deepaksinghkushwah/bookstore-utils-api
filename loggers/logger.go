package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log is logger var
	log *zap.Logger
)

// GetLogger return logger
func GetLogger() *zap.Logger {
	return log
}

// Log to log msg
func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error

	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}

}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()

}

func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}
	log.Error(msg, tags...)
	log.Sync()
}
