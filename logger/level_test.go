package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLevel(t *testing.T) {
	assert.Equal(t, ParseLevel(Disabled.String()), Disabled)
	assert.Equal(t, ParseLevel(NoLevel.String()), NoLevel)
	assert.Equal(t, ParseLevel(TraceLevel.String()), TraceLevel)
	assert.Equal(t, ParseLevel(DebugLevel.String()), DebugLevel)
	assert.Equal(t, ParseLevel(InfoLevel.String()), InfoLevel)
	assert.Equal(t, ParseLevel(WarnLevel.String()), WarnLevel)
	assert.Equal(t, ParseLevel(ErrorLevel.String()), ErrorLevel)
	assert.Equal(t, ParseLevel(FatalLevel.String()), FatalLevel)
	assert.Equal(t, ParseLevel(PanicLevel.String()), PanicLevel)

	for _, lvl := range []Level{
		Disabled, NoLevel, TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel,
	} {
		assert.Equal(t, ParseLevel(lvl.String()), lvl)
	}
}
