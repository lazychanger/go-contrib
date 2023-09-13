package zaplog

import (
	"github.com/lazychanger/go-contrib/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var (
	globalLogger       *zap.Logger
	globalSurgerLogger *zap.SugaredLogger
	globalLvl          zapcore.Level
	format             = "color"
)

func init() {
	_ = zap.RegisterEncoder("color", func(config zapcore.EncoderConfig) (zapcore.Encoder, error) {
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncodeCaller = zapcore.ShortCallerEncoder
		return zapcore.NewConsoleEncoder(config), nil
	})
}

func Init() {
	var (
		cfg = zap.NewProductionConfig()
		err error
	)
	if common.DEBUG {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.Level = zap.NewAtomicLevelAt(globalLvl)
	cfg.Encoding = format

	globalLogger, err = cfg.Build()

	if err != nil {
		log.Panicf("zap logger initialization failed: %s", err)
		return
	}
}

func SetLevel(lvl string) {
	zapLvl, err := zapcore.ParseLevel(lvl)
	if err != nil {
		log.Panicf("zap logger level parse failed: %s", err)
		return
	}

	globalLvl = zapLvl
}

// SetFormat will set zap logger format. allowed json,console,color
func SetFormat(fmt string) {
	format = fmt
}

func GetLogger() *zap.Logger {
	if globalLogger == nil {
		Init()
	}
	return globalLogger
}

func GetSuperLogger() *zap.SugaredLogger {
	if globalSurgerLogger == nil {
		globalSurgerLogger = GetLogger().Sugar()
	}

	return globalSurgerLogger
}
