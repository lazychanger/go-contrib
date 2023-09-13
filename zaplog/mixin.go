package zaplog

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"syscall"
)

func Sync() error {
	group := &errgroup.Group{}

	group.Go(GetLogger().Sync)
	group.Go(GetSuperLogger().Sync)

	if err := group.Wait(); errors.Is(err, syscall.ENOTTY) {
		return nil
	} else {
		return err
	}
}

func Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	GetLogger().Log(lvl, msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Log(zapcore.DebugLevel, msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Log(zapcore.InfoLevel, msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log(zapcore.WarnLevel, msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log(zapcore.ErrorLevel, msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	Log(zapcore.DPanicLevel, msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Log(zapcore.PanicLevel, msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log(zapcore.FatalLevel, msg, fields...)
}

func Debugf(template string, args ...any) {
	GetSuperLogger().Debugf(template, args...)
}

func Infof(template string, args ...any) {
	GetSuperLogger().Infof(template, args...)
}

func Warnf(template string, args ...any) {
	GetSuperLogger().Warnf(template, args...)
}

func Errorf(template string, args ...any) {
	GetSuperLogger().Errorf(template, args...)
}

func DPanicf(template string, args ...any) {
	GetSuperLogger().DPanicf(template, args...)
}

func Panicf(template string, args ...any) {
	GetSuperLogger().Panicf(template, args...)
}

func Fatalf(template string, args ...any) {
	GetSuperLogger().Fatalf(template, args...)
}

func Debugw(msg string, keyAndValues ...any) {
	GetSuperLogger().Debugw(msg, keyAndValues...)
}

func Infow(msg string, keyAndValues ...any) {
	GetSuperLogger().Infow(msg, keyAndValues...)
}

func Warnw(msg string, keyAndValues ...any) {
	GetSuperLogger().Warnw(msg, keyAndValues...)
}

func Errorw(msg string, keyAndValues ...any) {
	GetSuperLogger().Errorw(msg, keyAndValues...)
}

func DPanicw(msg string, keyAndValues ...any) {
	GetSuperLogger().DPanicw(msg, keyAndValues...)
}

func Panicw(msg string, keyAndValues ...any) {
	GetSuperLogger().Panicw(msg, keyAndValues...)
}

func Fatalw(msg string, keyAndValues ...any) {
	GetSuperLogger().Fatalw(msg, keyAndValues...)
}

func Debugln(args ...any) {
	GetSuperLogger().Debugln(args...)
}

func Infoln(args ...any) {
	GetSuperLogger().Infoln(args...)
}

func Warnln(args ...any) {
	GetSuperLogger().Warnln(args...)
}

func Errorln(args ...any) {
	GetSuperLogger().Errorln(args...)
}

func DPanicln(args ...any) {
	GetSuperLogger().DPanicln(args...)
}

func Panicln(args ...any) {
	GetSuperLogger().Panicln(args...)
}

func Fatalln(args ...any) {
	GetSuperLogger().Fatalln(args...)
}
