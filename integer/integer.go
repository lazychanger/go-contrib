package integer

import (
	"encoding/binary"
	"golang.org/x/exp/constraints"
	"unsafe"
)

const (
	INT8Size    = 1
	INT16Size   = 2
	INT32Size   = 4
	INT64Size   = 8
	INTSize     = INT64Size
	UINT8Size   = 1
	UINT16Size  = 2
	UINT32Size  = 4
	UINT64Size  = 8
	UINTSize    = UINT64Size
	Float32Size = 4
	Float64Size = 8
)

// Integer is a generic interface that represents any integer or float type.
type Integer interface {
	constraints.Integer | constraints.Float
}

// GetSize returns the size in bytes of the given value, based on its type.
// It uses a type switch to handle different types.
func GetSize(v any) int {
	switch v.(type) {
	case int8:
		return INT8Size
	case int16:
		return INT16Size
	case int32:
		return INT32Size
	case int64:
		return INT64Size
	case int:
		return INTSize
	case uint8:
		return UINT8Size
	case uint16:
		return UINT16Size
	case uint32:
		return UINT32Size
	case uint64:
		return UINT64Size
	case uint:
		return UINTSize
	case float32:
		return Float32Size
	case float64:
		return Float64Size
	default:
		return int(unsafe.Sizeof(v))
	}
}

// GetSizeUnsafe returns the size in bytes of the given value, based on its type.
// It uses unsafe.Sizeof to get the size of any type that implements Integer.
func GetSizeUnsafe[T Integer](v T) int {
	return int(unsafe.Sizeof(v))
}

// Int64ToBytes converts an int64 value to a byte slice in big endian order.
func Int64ToBytes(i int64) []byte {
	b := make([]byte, INT64Size)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

// BytesToInt64 converts a byte slice in big endian order to an int64 value.
func BytesToInt64(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

// Int32ToBytes converts an int32 value to a byte slice in big endian order.
func Int32ToBytes(i int32) []byte {
	b := make([]byte, INT32Size)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

// BytesToInt32 converts a byte slice in big endian order to an int32 value.
func BytesToInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}

// IntToBytes converts an int value to a byte slice in big endian order.
// It is equivalent to Int64ToBytes with an int64 cast.
func IntToBytes(i int) []byte {
	return Int64ToBytes(int64(i))
}

// BytesToInt converts a byte slice in big endian order to an int value.
// It is equivalent to BytesToInt64 with an int cast.
func BytesToInt(b []byte) int {
	return int(BytesToInt64(b))
}

// Int16ToBytes converts an int16 value to a byte slice in big endian order.
func Int16ToBytes(i int16) []byte {
	b := make([]byte, INT16Size)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

// BytesToInt16 converts a byte slice in big endian order to an int16 value.
func BytesToInt16(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(b))
}

// Int8ToBytes converts an int8 value to a byte slice.
// It is equivalent to casting the value to uint8 and wrapping it in a slice.
func Int8ToBytes(i int8) []byte {
	return []byte{uint8(i)}
}

// BytesToInt8 converts a byte slice to an int8 value.
// It is equivalent to taking the first element of the slice and casting it to int8.
func BytesToInt8(b []byte) int8 {
	return int8(b[0])
}

// Uint64ToBytes converts a uint64 value to a byte slice in big endian order.
func Uint64ToBytes(i uint64) []byte {
	b := make([]byte, UINT64Size)
	binary.BigEndian.PutUint64(b, i)
	return b
}

// BytesToUint64 converts a byte slice in big endian order to a uint64 value.
func BytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Uint32ToBytes converts a uint32 value to a byte slice in big endian order.
func Uint32ToBytes(i uint32) []byte {
	b := make([]byte, UINT32Size)
	binary.BigEndian.PutUint32(b, i)
	return b
}

// BytesToUint32 converts a byte slice in big endian order to a uint32 value.
func BytesToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

// UintToBytes converts a uint value to a byte slice in big endian order.
// It is equivalent to Uint64ToBytes with a uint64 cast.
func UintToBytes(i uint) []byte {
	return Uint64ToBytes(uint64(i))
}

// BytesToUint converts a byte slice in big endian order to a uint value.
// It is equivalent to BytesToUint64 with a uint cast.
func BytesToUint(b []byte) uint {
	return uint(BytesToUint64(b))
}

// Uint16ToBytes converts a uint16 value to a byte slice in big endian order.
func Uint16ToBytes(i uint16) []byte {
	b := make([]byte, UINT16Size)
	binary.BigEndian.PutUint16(b, i)
	return b
}

// BytesToUint16 converts a byte slice in big endian order to a uint16 value.
func BytesToUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

// Uint8ToBytes converts a uint8 value to a byte slice.
// It is equivalent to wrapping the value in a slice.
func Uint8ToBytes(i uint8) []byte {
	return []byte{i}
}

// BytesToUint8 converts a byte slice to a uint8 value.
// It is equivalent to taking the first element of the slice.
func BytesToUint8(b []byte) uint8 {
	return b[0]
}

// No support float and bytes transfer

// Float64ToBytes converts a float64 value to a byte slice in big endian order.
// It uses math.Float64bits to get the binary representation of the float as a uint64,
// and then uses binary.BigEndian.PutUint64 to encode it as bytes.
//func Float64ToBytes(f float64) []byte {
//	b := make([]byte, Float64Size)
//	n := math.Float64bits(f)
//	binary.BigEndian.PutUint64(b, n)
//	return b
//}

// BytesToFloat64 converts a byte slice in big endian order to a float64 value.
// It uses binary.BigEndian.Uint64 to decode the bytes as a uint64,
// and then uses math.Float64frombits to convert it to a float.
//func BytesToFloat64(b []byte) float64 {
//	return math.Float64frombits(binary.BigEndian.Uint64(b))
//}

// Float32ToBytes converts a float32 value to a byte slice in big endian order.
// It uses math.Float32bits to get the binary representation of the float as a uint32,
// and then uses binary.BigEndian.PutUint32 to encode it as bytes.
//func Float32ToBytes(f float32) []byte {
//	b := make([]byte, Float32Size)
//	n := math.Float32bits(f)
//	binary.BigEndian.PutUint32(b, n)
//	return b
//}

// BytesToFloat32 converts a byte slice in big endian order to a float32 value.
// It uses binary.BigEndian.Uint32 to decode the bytes as a uint32,
// and then uses math.Float32frombits to convert it to a float.
//func BytesToFloat32(b []byte) float32 {
//	return math.Float32frombits(binary.BigEndian.Uint32(b))
//}
