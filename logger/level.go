package logger

type Level int8

//go:generate stringer -type=Level -linecomment
const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota // debug
	// InfoLevel defines info log level.
	InfoLevel // info
	// WarnLevel defines warn log level.
	WarnLevel // warn
	// ErrorLevel defines error log level.
	ErrorLevel // error
	// FatalLevel defines fatal log level.
	FatalLevel // fatal
	// PanicLevel defines panic log level.
	PanicLevel // panic
	// NoLevel defines an absent log level.
	NoLevel // no
	// Disabled disables the logger.
	Disabled // disabled

	// TraceLevel defines trace log level.
	TraceLevel Level = -1 // trace
	// Values less than TraceLevel are handled as numbers.
)

// ParseLevel will transfer string to Level
func ParseLevel(lvl string) Level {
	switch lvl {
	case DebugLevel.String():
		return DebugLevel
	case InfoLevel.String():
		return InfoLevel
	case WarnLevel.String():
		return WarnLevel
	case ErrorLevel.String():
		return ErrorLevel
	case FatalLevel.String():
		return FatalLevel
	case PanicLevel.String():
		return PanicLevel
	case Disabled.String():
		return Disabled
	case TraceLevel.String():
		return TraceLevel
	default:
		return NoLevel
	}
}
