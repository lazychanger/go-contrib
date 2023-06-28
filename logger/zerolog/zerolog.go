package zerolog

import (
	"github.com/lazychanger/go-contrib/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	logger.RegisterDefaultLogger(New(log.Logger))
}

type Logger struct {
	zlog *zerolog.Logger
	lvl  logger.Level
}

func (l *Logger) Log(lvl logger.Level, msg string) {
	l.LogF(lvl, msg)
}

func (l *Logger) LogF(lvl logger.Level, msg string, a ...any) {
	l.zlog.WithLevel(ParseLevel(lvl)).Msgf(msg, a...)
}

func (l *Logger) With(name string, value any) logger.Logger {
	return l.Withs(logger.Field{Name: name, Value: value})
}

func (l *Logger) Withs(fields ...logger.Field) logger.Logger {
	fs := map[string]interface{}{}
	for _, field := range fields {
		fs[field.Name] = field.Value
	}

	return New(l.zlog.With().Fields(fs).Logger())
}

func (l *Logger) WithLevel(lvl logger.Level) logger.Logger {
	return New(l.zlog.Level(ParseLevel(lvl)))
}
func (l *Logger) WithFormat(format logger.Format) logger.Logger {
	newlogger := New(*l.zlog)
	newlogger.SetFormat(format)
	return newlogger
}

func (l *Logger) Level() logger.Level {
	l.zlog.Info()
	return l.lvl
}

func (l *Logger) SetLevel(lvl logger.Level) {
	l.lvl = lvl
	newzlog := l.zlog.Level(ParseLevel(lvl))
	l.zlog = &newzlog
}

func (l *Logger) SetFormat(format logger.Format) {
	var newzlog = *l.zlog
	switch format {
	case logger.Text:
		newzlog = l.zlog.Output(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.NoColor = true
		}))
		break
	case logger.Json:
		newzlog = l.zlog.Output(os.Stderr)
		break
	default:
		newzlog = l.zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		break
	}

	l.zlog = &newzlog
}

func (l *Logger) Trace(msg string) {
	l.LogF(logger.TraceLevel, msg)
}

func (l *Logger) Debug(msg string) {
	l.LogF(logger.DebugLevel, msg)
}

func (l *Logger) Info(msg string) {
	l.LogF(logger.InfoLevel, msg)
}

func (l *Logger) Warn(msg string) {
	l.LogF(logger.WarnLevel, msg)
}

func (l *Logger) Error(msg string) {
	l.LogF(logger.ErrorLevel, msg)
}

func (l *Logger) Fatal(msg string) {
	l.LogF(logger.FatalLevel, msg)
}

func (l *Logger) Panic(msg string) {
	l.LogF(logger.PanicLevel, msg)
}

func (l *Logger) TraceF(msg string, a ...any) {
	l.LogF(logger.TraceLevel, msg, a...)
}

func (l *Logger) DebugF(msg string, a ...any) {
	l.LogF(logger.DebugLevel, msg, a...)
}

func (l *Logger) InfoF(msg string, a ...any) {
	l.LogF(logger.InfoLevel, msg, a...)
}

func (l *Logger) WarnF(msg string, a ...any) {
	l.LogF(logger.WarnLevel, msg, a...)
}

func (l *Logger) ErrorF(msg string, a ...any) {
	l.LogF(logger.ErrorLevel, msg, a...)
}

func (l *Logger) FatalF(msg string, a ...any) {
	l.LogF(logger.FatalLevel, msg, a...)
}

func (l *Logger) PanicF(msg string, a ...any) {
	l.LogF(logger.PanicLevel, msg, a...)
}

func (l *Logger) Register() {
	logger.RegisterDefaultLogger(l)
}

func New(zlog zerolog.Logger) *Logger {
	return &Logger{&zlog, logger.Level(zlog.GetLevel())}
}

func ParseLevel(lvl logger.Level) zerolog.Level {
	return zerolog.Level(lvl)
}
