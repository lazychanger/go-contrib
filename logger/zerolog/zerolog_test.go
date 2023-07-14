package zerolog

import (
	"bytes"
	"fmt"
	loggerPkg "github.com/lazychanger/go-contrib/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"testing"
)

type testingLog struct {
	testing.TB
	buf bytes.Buffer
}

func (t *testingLog) Log(args ...interface{}) {
	if _, err := t.buf.WriteString(fmt.Sprint(args...)); err != nil {
		t.Error(err)
	}
}

func (t *testingLog) Logf(format string, args ...interface{}) {
	if _, err := t.buf.WriteString(fmt.Sprintf(format, args...)); err != nil {
		t.Error(err)
	}
}

func TestLogger(t *testing.T) {

	tests := []struct {
		name  string
		write []byte
		want  []byte
	}{{
		name:  "newline",
		write: []byte("newline\n"),
		want:  []byte("newline"),
	}, {
		name:  "oneline",
		write: []byte("oneline"),
		want:  []byte("oneline"),
	}, {
		name:  "twoline",
		write: []byte("twoline\n\n"),
		want:  []byte("twoline"),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := &testingLog{TB: t} // Capture TB log buffer.
			w := zerolog.TestWriter{T: tb}

			n, err := w.Write(tt.write)
			if err != nil {
				t.Error(err)
			}
			if n != len(tt.write) {
				t.Errorf("Expected %d write length but got %d", len(tt.write), n)
			}
			p := tb.buf.Bytes()
			if !bytes.Equal(tt.want, p) {
				t.Errorf("Expected %q, got %q.", tt.want, p)
			}

			log := New(zerolog.New(zerolog.NewConsoleWriter(zerolog.ConsoleTestWriter(t))))
			log.With("name", tt.name).Info("Success!")
			tb.buf.Reset()
		})
	}
}

func BenchmarkLogger(b *testing.B) {
	nlog := New(log.Logger)
	b.Run("logger", func(b *testing.B) {
		nlog.SetLevel(loggerPkg.ErrorLevel)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			nlog.Info(fmt.Sprintf("%d", i))
		}
	})

	b.Run("zerolog", func(b *testing.B) {
		log.Logger = log.Level(zerolog.ErrorLevel)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			log.Info().Msg(fmt.Sprintf("%d", i))
		}
	})
}
