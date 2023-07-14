package zerolog

import (
	loggerPkg "github.com/lazychanger/go-contrib/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
)

func init() {
	loggerPkg.RegisterDefaultLogger(New(log.Logger))
}

type logger struct {
	zlog *zerolog.Logger
	lvl  loggerPkg.Level
	fmt  loggerPkg.Format
}

func (l *logger) Write(p []byte) (n int, err error) {
	return l.zlog.Write(p)
}

func (l *logger) Log(lvl loggerPkg.Level, msg string) {
	l.LogF(lvl, msg)
}

func (l *logger) LogF(lvl loggerPkg.Level, msg string, a ...any) {
	l.zlog.WithLevel(ParseLevel(lvl)).Msgf(msg, a...)
}

func (l *logger) With(name string, value any) loggerPkg.Logger {
	return l.Withs(loggerPkg.Field{Name: name, Value: value})
}

func (l *logger) Withs(fields ...loggerPkg.Field) loggerPkg.Logger {
	fs := map[string]interface{}{}
	for _, field := range fields {
		fs[field.Name] = field.Value
	}

	return New(l.zlog.With().Fields(fs).Logger())
}

func (l *logger) WithLevel(lvl loggerPkg.Level) loggerPkg.Logger {
	return New(l.zlog.Level(ParseLevel(lvl)))
}
func (l *logger) WithFormat(format loggerPkg.Format) loggerPkg.Logger {
	newlogger := New(*l.zlog)
	newlogger.SetFormat(format)
	return newlogger
}

func (l *logger) Level() loggerPkg.Level {
	return l.lvl
}

func (l *logger) Format() loggerPkg.Format {
	return l.fmt
}

func (l *logger) SetWriter(writer io.Writer) {
	output := l.zlog.Output(writer)
	l.zlog = &output
}

func (l *logger) SetLevel(lvl loggerPkg.Level) {
	l.lvl = lvl
	newzlog := l.zlog.Level(ParseLevel(lvl))
	l.zlog = &newzlog
}

func (l *logger) SetFormat(format loggerPkg.Format) {
	var newzlog = *l.zlog
	switch format {
	case loggerPkg.Text:
		newzlog = l.zlog.Output(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.NoColor = true
		}))
		break
	case loggerPkg.Json:
		newzlog = l.zlog.Output(os.Stderr)
		break
	default:
		newzlog = l.zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		break
	}

	l.zlog = &newzlog
	l.fmt = format
}

func (l *logger) Trace(msg string) {
	l.LogF(loggerPkg.TraceLevel, msg)
}

func (l *logger) Debug(msg string) {
	l.LogF(loggerPkg.DebugLevel, msg)
}

func (l *logger) Info(msg string) {
	l.LogF(loggerPkg.InfoLevel, msg)
}

func (l *logger) Warn(msg string) {
	l.LogF(loggerPkg.WarnLevel, msg)
}

func (l *logger) Error(msg string) {
	l.LogF(loggerPkg.ErrorLevel, msg)
}

func (l *logger) Fatal(msg string) {
	l.LogF(loggerPkg.FatalLevel, msg)
}

func (l *logger) Panic(msg string) {
	l.LogF(loggerPkg.PanicLevel, msg)
}

func (l *logger) TraceF(msg string, a ...any) {
	l.LogF(loggerPkg.TraceLevel, msg, a...)
}

func (l *logger) DebugF(msg string, a ...any) {
	l.LogF(loggerPkg.DebugLevel, msg, a...)
}

func (l *logger) InfoF(msg string, a ...any) {
	l.LogF(loggerPkg.InfoLevel, msg, a...)
}

func (l *logger) WarnF(msg string, a ...any) {
	l.LogF(loggerPkg.WarnLevel, msg, a...)
}

func (l *logger) ErrorF(msg string, a ...any) {
	l.LogF(loggerPkg.ErrorLevel, msg, a...)
}

func (l *logger) FatalF(msg string, a ...any) {
	l.LogF(loggerPkg.FatalLevel, msg, a...)
}

func (l *logger) PanicF(msg string, a ...any) {
	l.LogF(loggerPkg.PanicLevel, msg, a...)
}

func New(zlog zerolog.Logger) loggerPkg.Logger {
	return &logger{
		zlog: &zlog,
		lvl:  loggerPkg.Level(zlog.GetLevel()),
		fmt:  loggerPkg.Color,
	}
}

func ParseLevel(lvl loggerPkg.Level) zerolog.Level {
	return zerolog.Level(lvl)
}
