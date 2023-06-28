package logger

import "fmt"

// Logger is an interface that defines the logging functionality of a program.
type Logger interface {

	// With returns a new Logger with the specified field added to the context.
	With(name string, value any) Logger

	// Withs returns a new Logger with the specified fields added to the context.
	Withs(fields ...Field) Logger

	// WithLevel returns a new Logger with the specified level.
	WithLevel(lvl Level) Logger

	// WithFormat returns a new Logger with the specified format.
	WithFormat(format Format) Logger

	// Level returns the current level of the logger.
	Level() Level

	// SetLevel sets the current level of the logger.
	SetLevel(lvl Level)

	// SetFormat sets the format of the logger.
	SetFormat(format Format)

	// Log logs a message at the specified level.
	Log(lvl Level, msg string)

	// LogF logs a formatted message at the specified level.
	LogF(lvl Level, msg string, a ...any)

	// Trace logs a trace message.
	Trace(msg string)

	// Debug logs a debug message.
	Debug(msg string)

	// Info logs an info message.
	Info(msg string)

	// Warn logs a warning message.
	Warn(msg string)

	// Error logs an error message.
	Error(msg string)

	// Fatal logs a fatal message and then panics.
	Fatal(msg string)

	// Panic logs a panic message and then panics.
	Panic(msg string)

	// TraceF logs a formatted trace message.
	TraceF(msg string, a ...any)

	// DebugF logs a formatted debug message.
	DebugF(msg string, a ...any)

	// InfoF logs a formatted info message.
	InfoF(msg string, a ...any)

	// WarnF logs a formatted warning message.
	WarnF(msg string, a ...any)

	// ErrorF logs a formatted error message.
	ErrorF(msg string, a ...any)

	// FatalF logs a formatted fatal message and then panics.
	FatalF(msg string, a ...any)

	// PanicF logs a formatted panic message and then panics.
	PanicF(msg string, a ...any)
}

func With(name string, value any) Logger {
	return GetLogger().With(name, value)
}

func Withs(fields ...Field) Logger {
	return GetLogger().Withs(fields...)
}

func Trace(msg string) {
	GetLogger().Trace(msg)
}

func Debug(msg string) {
	GetLogger().Debug(msg)
}

func Info(msg string) {
	GetLogger().Info(msg)
}

func Warn(msg string) {
	GetLogger().Warn(msg)
}

func Error(msg string) {
	GetLogger().Error(msg)
}

func Fatal(msg string) {
	GetLogger().Fatal(msg)
}

func Panic(msg string) {
	GetLogger().Panic(msg)
}

func TraceF(msg string, a ...any) {
	GetLogger().TraceF(msg, a...)
}

func DebugF(msg string, a ...any) {
	GetLogger().DebugF(msg, a...)
}

func InfoF(msg string, a ...any) {
	GetLogger().InfoF(msg, a...)
}

func WarnF(msg string, a ...any) {
	GetLogger().WarnF(msg, a...)
}

func ErrorF(msg string, a ...any) {
	GetLogger().ErrorF(msg, a...)
}

func FatalF(msg string, a ...any) {
	GetLogger().FatalF(msg, a...)
}

func PanicF(msg string, a ...any) {
	GetLogger().PanicF(msg, a...)
}

func Print(a ...any) {
	GetLogger().Debug(fmt.Sprint(a...))
}

func Printf(format string, a ...any) {
	GetLogger().Debug(fmt.Sprintf(format, a...))
}

func SetLevel(lvl Level) {
	GetLogger().SetLevel(lvl)
}

func SetFormat(format Format) {
	GetLogger().SetFormat(format)
}

func WithLevel(lvl Level) Logger {
	return GetLogger().WithLevel(lvl)
}

func WithFormat(format Format) Logger {
	return GetLogger().WithFormat(format)
}
