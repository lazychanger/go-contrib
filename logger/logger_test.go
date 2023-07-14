package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetLevel(t *testing.T) {
	for _, lvl := range []Level{
		Disabled, NoLevel, TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel,
	} {
		SetLevel(lvl)
		assert.Equal(t, lvl, GetLogger().Level())
	}
}

func TestSetFormat(t *testing.T) {
	for _, fmt := range []Format{
		Json, Text, Color,
	} {
		SetFormat(fmt)
		assert.Equal(t, fmt, GetLogger().Format())
	}
}

//func TestLogger(t *testing.T) {
//log := GetLogger()
//
//buf := &bytes.Buffer{}
//
//log.SetWriter(buf)
//
//msg := "hello world"
//
//log.Info(msg)
//
//assert.Equal(t, buf.String(), msg)
//}
