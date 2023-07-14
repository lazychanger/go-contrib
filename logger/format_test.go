package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFormat(t *testing.T) {
	for _, fmt := range []Format{
		Json, Text, Color,
	} {
		assert.Equal(t, ParseFormat(fmt.String()), fmt)
	}
}
