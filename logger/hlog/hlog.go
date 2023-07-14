package hlog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	loggerPkg "github.com/lazychanger/go-contrib/logger"
	"io"
	"os"
)

func init() {
	loggerPkg.RegisterDefaultLogger(New())
}

type logger struct {
	io.Writer
	hlog.FullLogger
	fields []loggerPkg.Field

	lvl loggerPkg.Level
}

func (l *logger) Write(p []byte) (n int, err error) {
	return l.Writer.Write(p)
}

func (l *logger) With(name string, value any) loggerPkg.Logger {
	return l.Withs(loggerPkg.Field{Name: name, Value: value})
}

func (l *logger) Withs(_ ...loggerPkg.Field) loggerPkg.Logger {
	panic("no implement Withs")
}

func (l *logger) WithLevel(lvl loggerPkg.Level) loggerPkg.Logger {
	l.SetLevel(lvl)
	return l
}

func (l *logger) WithFormat(_ loggerPkg.Format) loggerPkg.Logger {
	panic("no implement WithFormat")
}

func (l *logger) Level() loggerPkg.Level {
	return l.lvl
}

func (l *logger) Format() loggerPkg.Format {
	panic("no implement Format")
}

func (l *logger) SetLevel(lvl loggerPkg.Level) {
	l.FullLogger.SetLevel(ParseLevel(lvl))
	l.lvl = lvl
}

func (l *logger) SetFormat(_ loggerPkg.Format) {
	panic("no implement SetFormat")
}

func (l *logger) SetWriter(writer io.Writer) {
	l.FullLogger.SetOutput(writer)
	l.Writer = writer
}

func (l *logger) Log(lvl loggerPkg.Level, msg string) {
	l.LogF(lvl, msg)
}

func (l *logger) LogF(lvl loggerPkg.Level, msg string, a ...any) {
	switch lvl {
	case loggerPkg.TraceLevel:
		l.TraceF(msg, a...)
	case loggerPkg.DebugLevel:
		l.DebugF(msg, a...)
	case loggerPkg.InfoLevel:
		l.InfoF(msg, a...)
	case loggerPkg.WarnLevel:
		l.WarnF(msg, a...)
	case loggerPkg.ErrorLevel:
		l.ErrorF(msg, a...)
	case loggerPkg.FatalLevel:
		l.FatalF(msg, a...)
	default:
		l.InfoF(msg, a...)
	}

}

func (l *logger) Trace(msg string) {
	l.FullLogger.Trace(msg)
}

func (l *logger) Debug(msg string) {
	l.FullLogger.Debug(msg)
}

func (l *logger) Info(msg string) {
	l.FullLogger.Info(msg)
}

func (l *logger) Warn(msg string) {
	l.FullLogger.Warn(msg)
}

func (l *logger) Error(msg string) {
	l.FullLogger.Error(msg)
}

func (l *logger) Fatal(msg string) {
	l.FullLogger.Fatal(msg)
}

func (l *logger) Panic(msg string) {
	l.FullLogger.Error(msg)
}

func (l *logger) TraceF(msg string, a ...any) {
	l.FullLogger.Tracef(msg, a...)
}

func (l *logger) DebugF(msg string, a ...any) {
	l.FullLogger.Debugf(msg, a...)
}

func (l *logger) InfoF(msg string, a ...any) {
	l.FullLogger.Infof(msg, a...)
}

func (l *logger) WarnF(msg string, a ...any) {
	l.FullLogger.Warnf(msg, a...)
}

func (l *logger) ErrorF(msg string, a ...any) {
	l.FullLogger.Errorf(msg, a...)
}

func (l *logger) FatalF(msg string, a ...any) {
	l.FullLogger.Fatalf(msg, a...)
}

func (l *logger) PanicF(msg string, a ...any) {
	l.FullLogger.Errorf(msg, a...)
}

func New() loggerPkg.Logger {
	hl := hlog.DefaultLogger()
	hl.SetLevel(hlog.LevelInfo)
	hl.SetOutput(os.Stdout)
	return &logger{
		FullLogger: hl,
		Writer:     os.Stdout,
		lvl:        loggerPkg.InfoLevel,
	}
}

func ParseLevel(lvl loggerPkg.Level) hlog.Level {
	switch lvl {
	case loggerPkg.TraceLevel:
		return hlog.LevelTrace
	case loggerPkg.DebugLevel:
		return hlog.LevelDebug
	case loggerPkg.InfoLevel:
		return hlog.LevelInfo
	case loggerPkg.WarnLevel:
		return hlog.LevelWarn
	case loggerPkg.ErrorLevel:
		return hlog.LevelError
	case loggerPkg.FatalLevel:
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
