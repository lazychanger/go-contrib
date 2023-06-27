package stringer

import (
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"testing"
)

func TestRandom(t *testing.T) {
	assert.Len(t, Random(32), 32)
	assert.Len(t, Random(32, HexChars), 32)
}

func TestShuffle(t *testing.T) {
	assert.NotEqual(t, Base64Chars, Shuffle(Base64Chars))
}

func BenchmarkRandom(b *testing.B) {
	const length = 32
	b.Run("rune", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Random(length)
		}
	})

	b.Run("third_party", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			randstr.String(length)
		}
	})
}

func TestToShort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "short",
			args: args{
				s: "aaaaa-bb_bbb",
			},
			want: "abb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToShort(tt.args.s); got != tt.want {
				t.Errorf("ToShort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperToShort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "upper to short",
			args: args{
				s: "aBbbbCcccDD",
			},
			want: "ABCDD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperToShort(tt.args.s); got != tt.want {
				t.Errorf("UpperToShort() = %v, want %v", got, tt.want)
			}
		})
	}
}
