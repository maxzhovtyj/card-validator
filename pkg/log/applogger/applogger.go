package applogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

func New() *zap.SugaredLogger {
	once.Do(func() {
		logger = createLogger()
	})

	return logger
}

func createLogger() *zap.SugaredLogger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	outputDir := "./logs"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "console",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr", outputDir + "/error.log",
		},
	}

	lg := zap.Must(config.Build())

	zap.ReplaceGlobals(lg)

	return lg.Sugar()
}
