package integer

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGetTypeSize(t *testing.T) {
	assert.Equal(t, GetSize(varT[int8]()), 1)
	assert.Equal(t, GetSize(varT[int16]()), 2)
	assert.Equal(t, GetSize(varT[int32]()), 4)
	assert.Equal(t, GetSize(varT[int]()), 8)
	assert.Equal(t, GetSize(varT[int64]()), 8)
	assert.Equal(t, GetSize(varT[uint8]()), 1)
	assert.Equal(t, GetSize(varT[uint16]()), 2)
	assert.Equal(t, GetSize(varT[uint32]()), 4)
	assert.Equal(t, GetSize(varT[uint]()), 8)
	assert.Equal(t, GetSize(varT[uint64]()), 8)
	assert.Equal(t, GetSize(varT[float32]()), 4)
	assert.Equal(t, GetSize(varT[float64]()), 8)
	//assert.Equal(t, GetSize("unknown type"), 0)

	assert.Equal(t, GetSizeUnsafe(varT[int8]()), 1)
	assert.Equal(t, GetSizeUnsafe(varT[int16]()), 2)
	assert.Equal(t, GetSizeUnsafe(varT[int32]()), 4)
	assert.Equal(t, GetSizeUnsafe(varT[int]()), 8)
	assert.Equal(t, GetSizeUnsafe(varT[int64]()), 8)
	assert.Equal(t, GetSizeUnsafe(varT[uint8]()), 1)
	assert.Equal(t, GetSizeUnsafe(varT[uint16]()), 2)
	assert.Equal(t, GetSizeUnsafe(varT[uint32]()), 4)
	assert.Equal(t, GetSizeUnsafe(varT[uint]()), 8)
	assert.Equal(t, GetSizeUnsafe(varT[uint64]()), 8)
	assert.Equal(t, GetSizeUnsafe(varT[float32]()), 4)
	assert.Equal(t, GetSizeUnsafe(varT[float64]()), 8)

	assert.NotEqual(t, GetSize(""), 0)
}

func TestIntTransfer(t *testing.T) {
	assert.Equal(t, math.MaxInt, BytesToInt(IntToBytes(math.MaxInt)))
	assert.Equal(t, math.MinInt, BytesToInt(IntToBytes(math.MinInt)))
	assert.Equal(t, int64(math.MaxInt64), BytesToInt64(Int64ToBytes(math.MaxInt64)))
	assert.Equal(t, int64(math.MinInt64), BytesToInt64(Int64ToBytes(math.MinInt64)))
	assert.Equal(t, int16(math.MaxInt16), BytesToInt16(Int16ToBytes(math.MaxInt16)))
	assert.Equal(t, int16(math.MinInt16), BytesToInt16(Int16ToBytes(math.MinInt16)))
	assert.Equal(t, int32(math.MaxInt32), BytesToInt32(Int32ToBytes(math.MaxInt32)))
	assert.Equal(t, int32(math.MinInt32), BytesToInt32(Int32ToBytes(math.MinInt32)))
	assert.Equal(t, int8(math.MaxInt8), BytesToInt8(Int8ToBytes(math.MaxInt8)))
	assert.Equal(t, int8(math.MinInt8), BytesToInt8(Int8ToBytes(math.MinInt8)))

	assert.True(t, math.MaxUint == BytesToUint(UintToBytes(math.MaxUint)))
	assert.True(t, math.MaxUint64 == BytesToUint64(Uint64ToBytes(math.MaxUint64)))
	assert.Equal(t, uint16(math.MaxUint16), BytesToUint16(Uint16ToBytes(math.MaxUint16)))
	assert.Equal(t, uint32(math.MaxUint32), BytesToUint32(Uint32ToBytes(math.MaxUint32)))
	assert.Equal(t, uint8(math.MaxUint8), BytesToUint8(Uint8ToBytes(math.MaxUint8)))

	// No support float and bytes transfer
	//assert.Equal(t, math.MaxFloat32, BytesToFloat32(Float32ToBytes(math.MaxFloat32)))
	//assert.Equal(t, math.SmallestNonzeroFloat32, BytesToFloat32(Float32ToBytes(math.SmallestNonzeroFloat32)))
	//assert.Equal(t, math.MaxFloat64, BytesToFloat64(Float64ToBytes(math.MaxFloat64)))
	//assert.Equal(t, math.SmallestNonzeroFloat64, BytesToFloat64(Float64ToBytes(math.SmallestNonzeroFloat64)))
}
func BenchmarkGetSize(b *testing.B) {

	b.Run("switch", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			GetSize(varT[int]())
			GetSize(varT[int32]())
			GetSize(varT[uint]())
			GetSize(varT[int64]())
			GetSize(varT[float64]())
		}
	})

	b.Run("unsafe", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			GetSizeUnsafe(varT[int]())
			GetSizeUnsafe(varT[int32]())
			GetSizeUnsafe(varT[uint]())
			GetSizeUnsafe(varT[int64]())
			GetSizeUnsafe(varT[float64]())
		}

	})
}

func varT[T any]() T {
	var t T
	return t
}
