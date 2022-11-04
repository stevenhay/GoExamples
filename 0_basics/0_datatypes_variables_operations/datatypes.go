package main

import "math/cmplx"

var (
	BoolType   bool   = false           // zero value = false
	StringType string = "Hello, World!" // zero value = ""
)

// Varies
var (
	MaxUnsignedInt uint = 18446744073709551615 // zero value = 0
	MaxSignedInt   int  = 9223372036854775807
	MinSignedInt   int  = -9223372036854775808
)

// Unsigned types
var (
	MaxUnsignedInt64 uint64 = 18446744073709551615
	MaxUnsignedInt32 uint32 = 4294967295
	MaxUnsignedInt16 uint16 = 65535
	MaxUnsignedInt8  uint8  = 255
	ByteAlias        byte   = 255 // byte is an alias for uint8
)

// Signed types
var (
	MaxSignedInt64 int64 = 9223372036854775807
	MinSignedInt64 int64 = -9223372036854775808

	MaxSignedInt32 int32 = 2147483647
	MinSignedInt32 int32 = -2147483648
	MaxRuneAlias   rune  = 2147483647
	MinRuneAlias   rune  = -2147483648

	MaxSignedInt16 int16 = 32767
	MinSignedInt16 int16 = -32768

	MaxSignedInt8 int8 = 127
	MinSignedInt8 int8 = -128
)

// Floats
var (
	MaxFloat64  float64 = 0x1p1023 * (1 + (1 - 0x1p-52))
	MinFloat64  float64 = 0x1p-1022 * 0x1p-52
	SomeFloat64 float64 = 9999999999999999999999999999999999999999999999999999999999999999999999999999999999999.999999999999999999999999999999999999999999999999999999999

	MaxFloat32  float32 = 0x1p127 * (1 + (1 - 0x1p-23))
	MinFloat32  float32 = 0x1p-126 * 0x1p-23
	SomeFloat32 float32 = 99999999999999999999999999999999999999.9999999999999999999999999999999999999999999999
)

// Complex
var (
	SomeComplex complex128 = cmplx.Sqrt(-5 + 12i)
)
