package logger

var globalLogger Logger

func RegisterDefaultLogger(logger Logger) {
	globalLogger = logger
}

func GetLogger() Logger {
	if globalLogger == nil {
		globalLogger = &emptyLogger{}
	}

	return globalLogger
}

type emptyLogger struct {
}

func (e *emptyLogger) With(_ string, _ any) Logger {
	return e
}

func (e *emptyLogger) Withs(_ ...Field) Logger {
	return e
}

func (e *emptyLogger) WithLevel(_ Level) Logger {
	return e
}

func (e *emptyLogger) WithFormat(_ Format) Logger {
	return e
}

func (e *emptyLogger) Level() Level {
	return NoLevel
}

func (e *emptyLogger) SetLevel(_ Level) {
}

func (e *emptyLogger) SetFormat(_ Format) {
}

func (e *emptyLogger) Log(_ Level, _ string) {
}

func (e *emptyLogger) LogF(_ Level, _ string, _ ...any) {
}

func (e *emptyLogger) Trace(_ string) {
}

func (e *emptyLogger) Debug(_ string) {
}

func (e *emptyLogger) Info(_ string) {
}

func (e *emptyLogger) Warn(_ string) {
}

func (e *emptyLogger) Error(_ string) {
}

func (e *emptyLogger) Fatal(_ string) {
}

func (e *emptyLogger) Panic(_ string) {
}

func (e *emptyLogger) TraceF(_ string, _ ...any) {

}

func (e *emptyLogger) DebugF(_ string, _ ...any) {

}

func (e *emptyLogger) InfoF(_ string, _ ...any) {

}

func (e *emptyLogger) WarnF(_ string, _ ...any) {
}

func (e *emptyLogger) ErrorF(_ string, _ ...any) {

}

func (e *emptyLogger) FatalF(_ string, _ ...any) {

}

func (e *emptyLogger) PanicF(_ string, _ ...any) {

}
