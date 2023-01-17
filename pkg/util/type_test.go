package util

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect string
		err    error
	}{
		{name: "src: 0", src: 0, expect: "0"},
		{name: "src: empty", src: "", expect: ""},
		{name: "src: nil", src: nil, expect: "", err: ErrEmptyValue},
		//int
		{name: "src: math.MaxInt", src: math.MaxInt, expect: "9223372036854775807"},
		{name: "src: math.MinInt", src: math.MinInt, expect: "-9223372036854775808"},
		{name: "src: math.MaxInt8", src: int8(math.MaxInt8), expect: "127"},
		{name: "src: math.MinInt8", src: int8(math.MinInt8), expect: "-128"},
		{name: "src: math.MaxInt16", src: int16(math.MaxInt16), expect: "32767"},
		{name: "src: math.MinInt16", src: int16(math.MinInt16), expect: "-32768"},
		{name: "src: math.MaxInt32", src: int32(math.MaxInt32), expect: "2147483647"},
		{name: "src: math.MinInt32", src: int32(math.MinInt32), expect: "-2147483648"},
		{name: "src: math.MaxInt64", src: int64(math.MaxInt64), expect: "9223372036854775807"},
		{name: "src: math.MinInt64", src: int64(math.MinInt64), expect: "-9223372036854775808"},
		//uint
		{name: "src: math.MaxUint", src: uint(math.MaxUint), expect: "18446744073709551615"},
		{name: "src: math.MaxUint8", src: uint8(math.MaxUint8), expect: "255"},
		{name: "src: math.MaxUint16", src: uint16(math.MaxUint16), expect: "65535"},
		{name: "src: math.MaxUint32", src: uint32(math.MaxUint32), expect: "4294967295"},
		{name: "src: math.MaxUint64", src: uint64(math.MaxUint64), expect: "18446744073709551615"},
		//float
		{name: "src: math.MaxFloat32", src: float32(math.MaxFloat32), expect: "340282346638528860000000000000000000000"},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: "179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{name: "src: -math.MaxFloat32", src: float32(-math.MaxFloat32), expect: "-340282346638528860000000000000000000000"},
		{name: "src: -math.MaxFloat64", src: -math.MaxFloat64, expect: "-179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{name: "src: 1234.5678", src: 1234.5678, expect: "1234.5678"},
		//bool
		{src: true, expect: "true"},
		{src: false, expect: "false"},
		//string
		{name: "src: \"to string\"", src: "to string", expect: "to string"},
		{name: "src: []byte(\"to string\")", src: []byte("to string"), expect: "to string"},
		{name: "src: []rune(\"to string\")", src: []rune("to string"), expect: "to string"},
		{name: "src: []string{\"a\", \"b\", \"c\", \"d\"}", src: []string{"a", "b", "c", "d"}, expect: "", err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToString(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToBool(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect bool
		err    error
	}{
		{name: "123456", src: 123456, expect: true, err: nil},
		{name: "int64(13579)", src: int64(13579), expect: true, err: nil},
		{name: "-123456", src: -123456, expect: false, err: nil},
		{name: "int64(-13579)", src: int64(-13579), expect: false, err: nil},
		{name: "uint(123456)", src: uint(123456), expect: true, err: nil},
		{name: "uint64(13579)", src: uint64(13579), expect: true, err: nil},
		{name: "1234.5678", src: 1234.5678, expect: true, err: nil},
		{name: "float32(1234.5)", src: float32(1234.5), expect: true, err: nil},
		{name: "-1234.5678", src: -1234.5678, expect: false, err: nil},
		{name: "float32(-1234.5)", src: float32(-1234.5), expect: false, err: nil},

		{name: "true", src: true, expect: true, err: nil},
		{name: "false", src: false, expect: false, err: nil},
		{name: "TOO BOOL", src: "TOO BOOL", expect: false, err: strconv.ErrSyntax},
		{name: "TRUE", src: "TRUE", expect: true, err: nil},
		{name: "false", src: "false", expect: false, err: nil},
		//{src: []byte("true"), expect: true},
		//{src: []byte("to string"), expect: false},
		//{src: []rune("to string"), expect: false},
		//{src: []string{"a", "b", "c", "d"}, expect: false},
		//{src: nil, expect: false},

		//error
		{name: "12345.12345 + 6666i", src: 12345.12345 + 6666i, expect: false, err: ErrDataTypeNotSupported},
		{name: "complex64(1234.5 + 6666i)", src: complex64(1234.5 + 6666i), expect: false, err: ErrDataTypeNotSupported},
		{name: "-12345.12345 + 6666i", src: -12345.12345 + 6666i, expect: false, err: ErrDataTypeNotSupported},
		{name: "complex64(-1234.5 + 6666i)", src: complex64(-1234.5 + 6666i), expect: false, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToBool(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToInt(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect int
		err    error
	}{
		{name: "src: 0", src: 0, expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: int8(math.MaxInt8)", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: int8(math.MinInt8)", src: int8(math.MinInt8), expect: math.MinInt8},
		{name: "src: int16(math.MaxInt16)", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: int16(math.MinInt16)", src: int16(math.MinInt16), expect: math.MinInt16},
		{name: "src: int32(math.MaxInt32)", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: int32(math.MinInt32)", src: int32(math.MinInt32), expect: math.MinInt32},
		{name: "src: int64(math.MaxInt64)", src: int64(math.MaxInt64), expect: math.MaxInt64},
		{name: "src: int64(math.MinInt64)", src: int64(math.MinInt64), expect: math.MinInt64},
		//uint
		{name: "src: uint8(math.MaxUint8)", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: uint16(math.MaxUint16)", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: uint32(math.MaxUint32)", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: uint64(math.MaxUint64)", src: uint64(math.MaxUint64), expect: 0, err: ErrOutOfRange}, // overflow
		//todo uintptr
		{name: "src: uintptr(0x7fffffffffffffff)", src: uintptr(0x7fffffffffffffff), expect: math.MaxInt64},
		{name: "src: uintptr(0x8000000000000000)", src: uintptr(0x8000000000000000), expect: math.MinInt64},
		{name: "src: uintptr(0xffffffffffffffff)", src: uintptr(0xffffffffffffffff), expect: -1}, // overflow
		//float
		{name: "src: float32(math.MaxInt64)", src: float32(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MaxInt64)", src: float64(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MinInt64)", src: float32(math.MinInt64), expect: math.MinInt64},
		{name: "src: float64(math.MinInt64)", src: float64(math.MinInt64), expect: math.MinInt64},
		{name: "src: float32(math.MaxFloat32)", src: float32(math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: float32(-math.MaxFloat32)", src: float32(-math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: -math.MaxFloat64", src: -math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: math.Pi", src: math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(math.Pi)", src: float32(math.Pi), expect: 0, err: ErrInvalidSyntax},
		{name: "src: -math.Pi", src: -math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(-math.Pi)", src: float32(-math.Pi), expect: 0, err: ErrInvalidSyntax},
		//complex64
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//boolean
		{name: "src: true", src: true, expect: 1},
		{name: "src: false", src: false, expect: 0},
		//string
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: math.MaxInt64},           //max int64
		{name: "src: 9223372036854775808", src: "-9223372036854775808", expect: math.MinInt64},          //min int64
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: 0, err: ErrOutOfRange}, //overflow
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 0, err: ErrInvalidSyntax},
		{name: "src: 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//json Number
		{name: "src: json.Number(9223372036854775807)", src: json.Number("9223372036854775807"), expect: math.MaxInt64},           //max int64
		{name: "src: json.Number(9223372036854775808)", src: json.Number("-9223372036854775808"), expect: math.MinInt64},          //min int64
		{name: "src: json.Number(18446744073709551615)", src: json.Number("18446744073709551615"), expect: 0, err: ErrOutOfRange}, //overflow
		{name: "src: json.Number(PI)", src: json.Number("3.14159265358979323846264338327950288419716939937510582097494459"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(1)", src: json.Number(" 1"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number( 3.14)", src: json.Number(" 3.14"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(empty)", src: json.Number(""), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(string)", src: json.Number("string"), expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToInt(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToInt32(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect int32
		err    error
	}{
		{name: "src: 0", src: 0, expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: int8(math.MaxInt8)", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: int8(math.MinInt8)", src: int8(math.MinInt8), expect: math.MinInt8},
		{name: "src: int16(math.MaxInt16)", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: int16(math.MinInt16)", src: int16(math.MinInt16), expect: math.MinInt16},
		{name: "src: int32(math.MaxInt32)", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: int32(math.MinInt32)", src: int32(math.MinInt32), expect: math.MinInt32},
		{name: "src: int64(math.MaxInt32)", src: int64(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: int64(math.MinInt32)", src: int64(math.MinInt32), expect: math.MinInt32},
		{name: "src: int64(math.MaxInt64)", src: int64(math.MaxInt64), expect: 0, err: ErrOverflow},
		{name: "src: int64(math.MinInt64)", src: int64(math.MinInt64), expect: 0, err: ErrOverflow},
		//uint
		{name: "src: uint8(math.MaxUint8)", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: uint16(math.MaxUint16)", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: uint32(math.MaxUint32)", src: uint32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: uint32(math.MaxUint32)", src: uint32(math.MaxUint32), expect: 0, err: ErrOverflow},
		{name: "src: uint64(math.MaxUint64)", src: uint64(math.MaxUint64), expect: 0, err: ErrOutOfRange}, // overflow
		//todo uintptr
		{name: "src: uintptr(0x7fffffffffffffff)", src: uintptr(0x7fffffffffffffff), expect: 0, err: ErrOutOfRange}, // overflow
		{name: "src: uintptr(0x8000000000000000)", src: uintptr(0x8000000000000000), expect: 0, err: ErrOutOfRange}, // overflow
		{name: "src: uintptr(0xffffffffffffffff)", src: uintptr(0xffffffffffffffff), expect: -1},                    // overflow
		//float
		{name: "src: float32(math.MaxInt32)", src: float32(math.MaxInt32), expect: 0, err: ErrMissingPrecision}, // float32类型大于 33554432 精度缺失
		{name: "src: float64(math.MaxInt32)", src: float64(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: float32(math.MinInt32)", src: float32(math.MinInt32), expect: math.MinInt32},
		{name: "src: float64(math.MinInt32)", src: float64(math.MinInt32), expect: math.MinInt32},

		{name: "src: float32(math.MaxInt64)", src: float32(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MaxInt64)", src: float64(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MinInt64)", src: float32(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float64(math.MinInt64)", src: float64(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float32(math.MaxFloat32)", src: float32(math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: float32(-math.MaxFloat32)", src: float32(-math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: -math.MaxFloat64", src: -math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: math.Pi", src: math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(math.Pi)", src: float32(math.Pi), expect: 0, err: ErrInvalidSyntax},
		{name: "src: -math.Pi", src: -math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(-math.Pi)", src: float32(-math.Pi), expect: 0, err: ErrInvalidSyntax},
		//complex64
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt32 + 2147483647i)", src: complex64(math.MaxInt32 + 2147483647i), expect: 0, err: ErrDataTypeNotSupported}, // float32类型大于 33554432 精度缺失
		{name: "src: complex64(math.MinInt32 + 2147483647i)", src: complex64(math.MinInt32 + 2147483647i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt32 - 2147483647i)", src: complex64(math.MaxInt32 - 2147483647i), expect: 0, err: ErrDataTypeNotSupported}, // float32类型大于 33554432 精度缺失
		{name: "src: complex64(math.MinInt32 - 2147483647i)", src: complex64(math.MinInt32 - 2147483647i), expect: 0, err: ErrDataTypeNotSupported},
		//boolean
		{name: "src: true", src: true, expect: 1},
		{name: "src: false", src: false, expect: 0},
		//string
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: 2147483647", src: "2147483647", expect: math.MaxInt32},
		{name: "src: -2147483648", src: "-2147483648", expect: math.MinInt32},
		{name: "src: 4294967295", src: "4294967295", expect: 0, err: ErrOutOfRange}, //overflow
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 0, err: ErrInvalidSyntax},
		{name: "src: 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//json Number
		{name: "src: json.Number(2147483647)", src: json.Number("2147483647"), expect: math.MaxInt32},
		{name: "src: json.Number(-2147483648)", src: json.Number("-2147483648"), expect: math.MinInt32},
		{name: "src: json.Number(4294967295)", src: json.Number("4294967295"), expect: 0, err: ErrOutOfRange}, //overflow
		{name: "src: json.Number(PI)", src: json.Number("3.14159265358979323846264338327950288419716939937510582097494459"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(1)", src: json.Number(" 1"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number( 3.14)", src: json.Number(" 3.14"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(empty)", src: json.Number(""), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(string)", src: json.Number("string"), expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToInt32(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToInt64(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect int64
		err    error
	}{
		{name: "src: 0", src: 0, expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: int8(math.MaxInt8)", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: int8(math.MinInt8)", src: int8(math.MinInt8), expect: math.MinInt8},
		{name: "src: int16(math.MaxInt16)", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: int16(math.MinInt16)", src: int16(math.MinInt16), expect: math.MinInt16},
		{name: "src: int32(math.MaxInt32)", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: int32(math.MinInt32)", src: int32(math.MinInt32), expect: math.MinInt32},
		{name: "src: int64(math.MaxInt64)", src: int64(math.MaxInt64), expect: math.MaxInt64},
		{name: "src: int64(math.MinInt64)", src: int64(math.MinInt64), expect: math.MinInt64},
		//uint
		{name: "src: uint8(math.MaxUint8)", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: uint16(math.MaxUint16)", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: uint32(math.MaxUint32)", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: uint64(math.MaxUint64)", src: uint64(math.MaxUint64), expect: 0, err: ErrOutOfRange}, // overflow
		//todo uintptr
		{name: "src: uintptr(0x7fffffffffffffff)", src: uintptr(0x7fffffffffffffff), expect: math.MaxInt64},
		{name: "src: uintptr(0x8000000000000000)", src: uintptr(0x8000000000000000), expect: math.MinInt64},
		{name: "src: uintptr(0xffffffffffffffff)", src: uintptr(0xffffffffffffffff), expect: -1}, // overflow
		//float
		{name: "src: float32(math.MaxInt64)", src: float32(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MaxInt64)", src: float64(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MinInt64)", src: float32(math.MinInt64), expect: math.MinInt64},
		{name: "src: float64(math.MinInt64)", src: float64(math.MinInt64), expect: math.MinInt64},
		{name: "src: float32(math.MaxFloat32)", src: float32(math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: float32(-math.MaxFloat32)", src: float32(-math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: -math.MaxFloat64", src: -math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: math.Pi", src: math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(math.Pi)", src: float32(math.Pi), expect: 0, err: ErrInvalidSyntax},
		{name: "src: -math.Pi", src: -math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(-math.Pi)", src: float32(-math.Pi), expect: 0, err: ErrInvalidSyntax},
		//complex64
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//boolean
		{name: "src: true", src: true, expect: 1},
		{name: "src: false", src: false, expect: 0},
		//string
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: math.MaxInt64},           //max int64
		{name: "src: 9223372036854775808", src: "-9223372036854775808", expect: math.MinInt64},          //min int64
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: 0, err: ErrOutOfRange}, //overflow
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 0, err: ErrInvalidSyntax},
		{name: "src: 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//json Number
		{name: "src: json.Number(9223372036854775807)", src: json.Number("9223372036854775807"), expect: math.MaxInt64},           //max int64
		{name: "src: json.Number(9223372036854775808)", src: json.Number("-9223372036854775808"), expect: math.MinInt64},          //min int64
		{name: "src: json.Number(18446744073709551615)", src: json.Number("18446744073709551615"), expect: 0, err: ErrOutOfRange}, //overflow
		{name: "src: json.Number(PI)", src: json.Number("3.14159265358979323846264338327950288419716939937510582097494459"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(1)", src: json.Number(" 1"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number( 3.14)", src: json.Number(" 3.14"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(empty)", src: json.Number(""), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(string)", src: json.Number("string"), expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToInt64(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToUint(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect uint
		err    error
	}{
		{name: "src: 0", src: 0, expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: int8(math.MaxInt8)", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: int8(math.MinInt8)", src: int8(math.MinInt8), expect: 0, err: ErrOutOfRange},
		{name: "src: int16(math.MaxInt16)", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: int16(math.MinInt16)", src: int16(math.MinInt16), expect: 0, err: ErrOutOfRange},
		{name: "src: int32(math.MaxInt32)", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: int32(math.MinInt32)", src: int32(math.MinInt32), expect: 0, err: ErrOutOfRange},
		{name: "src: int64(math.MaxInt64)", src: int64(math.MaxInt64), expect: math.MaxInt64},
		{name: "src: int64(math.MinInt64)", src: int64(math.MinInt64), expect: 0, err: ErrOutOfRange},
		//uint
		{name: "src: uint8(math.MaxUint8)", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: uint16(math.MaxUint16)", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: uint32(math.MaxUint32)", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: uint64(math.MaxUint64)", src: uint64(math.MaxUint64), expect: math.MaxUint64},
		//todo uintptr
		{name: "src: uintptr(0x7fffffffffffffff)", src: uintptr(0x7fffffffffffffff), expect: math.MaxInt64},
		{name: "src: uintptr(0x8000000000000000)", src: uintptr(0x8000000000000000), expect: math.MaxInt64 + 1},
		{name: "src: uintptr(0xffffffffffffffff)", src: uintptr(0xffffffffffffffff), expect: math.MaxUint64},
		//float
		{name: "src: float32(maxFloat32UintBit)", src: float32(maxFloat32UintBit), expect: maxFloat32UintBit},
		{name: "src: float64(maxFloat64UintBit)", src: float64(maxFloat64UintBit), expect: maxFloat64UintBit},
		{name: "src: float32(math.MaxInt64)", src: float32(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MaxInt64)", src: float64(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MinInt64)", src: float32(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float64(math.MinInt64)", src: float64(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float32(math.MinInt64)", src: float32(math.MaxUint64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MinInt64)", src: float64(math.MaxUint64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MaxFloat32)", src: float32(math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: float32(-math.MaxFloat32)", src: float32(-math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: -math.MaxFloat64", src: -math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: math.Pi", src: math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(math.Pi)", src: float32(math.Pi), expect: 0, err: ErrInvalidSyntax},
		{name: "src: -math.Pi", src: -math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(-math.Pi)", src: float32(-math.Pi), expect: 0, err: ErrInvalidSyntax},
		//complex64
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//boolean
		{name: "src: true", src: true, expect: 1},
		{name: "src: false", src: false, expect: 0},
		//string
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: math.MaxUint64},
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: math.MaxInt64},
		{name: "src: -9223372036854775808", src: "-9223372036854775808", expect: 0, err: ErrInvalidSyntax},
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 0, err: ErrInvalidSyntax},
		{name: "src: 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//json Number
		{name: "src: json.Number(18446744073709551615)", src: json.Number("18446744073709551615"), expect: math.MaxUint64},
		{name: "src: json.Number(9223372036854775807)", src: json.Number("9223372036854775807"), expect: math.MaxInt64},
		{name: "src: json.Number(-9223372036854775808)", src: json.Number("-9223372036854775808"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(PI)", src: json.Number("3.14159265358979323846264338327950288419716939937510582097494459"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(1)", src: json.Number(" 1"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number( 3.14)", src: json.Number(" 3.14"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(empty)", src: json.Number(""), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(string)", src: json.Number("string"), expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToUint(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToUint32(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect uint32
		err    error
	}{
		{name: "src: 0", src: 0, expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: int8(math.MaxInt8)", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: int8(math.MinInt8)", src: int8(math.MinInt8), expect: 0, err: ErrOutOfRange},
		{name: "src: int16(math.MaxInt16)", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: int16(math.MinInt16)", src: int16(math.MinInt16), expect: 0, err: ErrOutOfRange},
		{name: "src: int32(math.MaxInt32)", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: int32(math.MinInt32)", src: int32(math.MinInt32), expect: 0, err: ErrOutOfRange},
		{name: "src: int64(math.MaxInt64)", src: int64(math.MaxInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: int64(math.MinInt64)", src: int64(math.MinInt64), expect: 0, err: ErrOutOfRange},
		//uint
		{name: "src: uint8(math.MaxUint8)", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: uint16(math.MaxUint16)", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: uint32(math.MaxUint32)", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: uint64(math.MaxUint64)", src: uint64(math.MaxUint64), expect: 0, err: ErrOutOfRange},
		//todo uintptr
		{name: "src: uintptr(0x7fffffff)", src: uintptr(0x7fffffff), expect: math.MaxInt32},
		{name: "src: uintptr(0x80000000)", src: uintptr(0x80000000), expect: math.MaxInt32 + 1},
		{name: "src: uintptr(0xffffffff)", src: uintptr(0xffffffff), expect: math.MaxUint32},
		//float
		{name: "src: float32(maxFloat32UintBit)", src: float32(maxFloat32UintBit), expect: maxFloat32UintBit},
		{name: "src: float64(maxFloat64UintBit)", src: float64(maxFloat64UintBit), expect: 0, err: ErrOutOfRange},
		{name: "src: float32(math.MaxInt64)", src: float32(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MaxInt64)", src: float64(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MinInt64)", src: float32(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float64(math.MinInt64)", src: float64(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float32(math.MinInt64)", src: float32(math.MaxUint64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MinInt64)", src: float64(math.MaxUint64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MaxFloat32)", src: float32(math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: float32(-math.MaxFloat32)", src: float32(-math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: -math.MaxFloat64", src: -math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: math.Pi", src: math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(math.Pi)", src: float32(math.Pi), expect: 0, err: ErrInvalidSyntax},
		{name: "src: -math.Pi", src: -math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(-math.Pi)", src: float32(-math.Pi), expect: 0, err: ErrInvalidSyntax},
		//complex64
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//boolean
		{name: "src: true", src: true, expect: 1},
		{name: "src: false", src: false, expect: 0},
		//string
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: 4294967295", src: "4294967295", expect: math.MaxUint32},
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: 0, err: ErrOutOfRange},
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: 0, err: ErrOutOfRange},
		{name: "src: -2147483648", src: "-2147483648", expect: 0, err: ErrInvalidSyntax},
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 0, err: ErrInvalidSyntax},
		{name: "src: 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//json Number
		{name: "src: json.Number(4294967295)", src: json.Number("4294967295"), expect: math.MaxUint32},
		{name: "src: json.Number(18446744073709551615)", src: json.Number("18446744073709551615"), expect: 0, err: ErrOutOfRange},
		{name: "src: json.Number(9223372036854775807)", src: json.Number("9223372036854775807"), expect: 0, err: ErrOutOfRange},
		{name: "src: json.Number(-2147483648)", src: json.Number("-2147483648"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(PI)", src: json.Number("3.14159265358979323846264338327950288419716939937510582097494459"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(1)", src: json.Number(" 1"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number( 3.14)", src: json.Number(" 3.14"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(empty)", src: json.Number(""), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(string)", src: json.Number("string"), expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToUint32(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToUint64(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect uint64
		err    error
	}{
		{name: "src: 0", src: 0, expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: int8(math.MaxInt8)", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: int8(math.MinInt8)", src: int8(math.MinInt8), expect: 0, err: ErrOutOfRange},
		{name: "src: int16(math.MaxInt16)", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: int16(math.MinInt16)", src: int16(math.MinInt16), expect: 0, err: ErrOutOfRange},
		{name: "src: int32(math.MaxInt32)", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: int32(math.MinInt32)", src: int32(math.MinInt32), expect: 0, err: ErrOutOfRange},
		{name: "src: int64(math.MaxInt64)", src: int64(math.MaxInt64), expect: math.MaxInt64},
		{name: "src: int64(math.MinInt64)", src: int64(math.MinInt64), expect: 0, err: ErrOutOfRange},
		//uint
		{name: "src: uint8(math.MaxUint8)", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: uint16(math.MaxUint16)", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: uint32(math.MaxUint32)", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: uint64(math.MaxUint64)", src: uint64(math.MaxUint64), expect: math.MaxUint64},
		//todo uintptr
		{name: "src: uintptr(0x7fffffffffffffff)", src: uintptr(0x7fffffffffffffff), expect: math.MaxInt64},
		{name: "src: uintptr(0x8000000000000000)", src: uintptr(0x8000000000000000), expect: math.MaxInt64 + 1},
		{name: "src: uintptr(0xffffffffffffffff)", src: uintptr(0xffffffffffffffff), expect: math.MaxUint64},
		//float
		{name: "src: float32(maxFloat32UintBit)", src: float32(maxFloat32UintBit), expect: maxFloat32UintBit},
		{name: "src: float64(maxFloat64UintBit)", src: float64(maxFloat64UintBit), expect: maxFloat64UintBit},
		{name: "src: float32(math.MaxInt64)", src: float32(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MaxInt64)", src: float64(math.MaxInt64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MinInt64)", src: float32(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float64(math.MinInt64)", src: float64(math.MinInt64), expect: 0, err: ErrOutOfRange},
		{name: "src: float32(math.MinInt64)", src: float32(math.MaxUint64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float64(math.MinInt64)", src: float64(math.MaxUint64), expect: 0, err: ErrMissingPrecision},
		{name: "src: float32(math.MaxFloat32)", src: float32(math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: float32(-math.MaxFloat32)", src: float32(-math.MaxFloat32), expect: 0, err: ErrOutOfRange},
		{name: "src: -math.MaxFloat64", src: -math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: math.Pi", src: math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(math.Pi)", src: float32(math.Pi), expect: 0, err: ErrInvalidSyntax},
		{name: "src: -math.Pi", src: -math.Pi, expect: 0, err: ErrInvalidSyntax},
		{name: "src: float32(-math.Pi)", src: float32(-math.Pi), expect: 0, err: ErrInvalidSyntax},
		//complex64
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//boolean
		{name: "src: true", src: true, expect: 1},
		{name: "src: false", src: false, expect: 0},
		//string
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: math.MaxUint64},
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: math.MaxInt64},
		{name: "src: -9223372036854775808", src: "-9223372036854775808", expect: 0, err: ErrInvalidSyntax},
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 0, err: ErrInvalidSyntax},
		{name: "src: 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//json Number
		{name: "src: json.Number(18446744073709551615)", src: json.Number("18446744073709551615"), expect: math.MaxUint64},
		{name: "src: json.Number(9223372036854775807)", src: json.Number("9223372036854775807"), expect: math.MaxInt64},
		{name: "src: json.Number(-9223372036854775808)", src: json.Number("-9223372036854775808"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(PI)", src: json.Number("3.14159265358979323846264338327950288419716939937510582097494459"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(1)", src: json.Number(" 1"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number( 3.14)", src: json.Number(" 3.14"), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(empty)", src: json.Number(""), expect: 0, err: ErrInvalidSyntax},
		{name: "src: json.Number(string)", src: json.Number("string"), expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToUint64(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToFloat(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect float64
		err    error
	}{
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: math.MaxInt", src: math.MaxInt, expect: math.MaxInt},
		{name: "src: math.MinInt", src: math.MinInt, expect: math.MinInt},
		{name: "src: math.MaxInt8", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: math.MinInt8", src: int8(math.MinInt8), expect: math.MinInt8},
		{name: "src: math.MaxInt16", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: math.MinInt16", src: int16(math.MinInt16), expect: math.MinInt16},
		{name: "src: math.MaxInt32", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: math.MinInt32", src: int32(math.MinInt32), expect: math.MinInt32},
		{name: "src: math.MaxInt64", src: int64(math.MaxInt64), expect: math.MaxInt64},
		{name: "src: math.MinInt64", src: int64(math.MinInt64), expect: math.MinInt64},
		//uint
		{name: "src: math.MaxInt", src: uint(math.MaxUint), expect: math.MaxUint},
		{name: "src: math.MaxInt8", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: math.MaxInt16", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: math.MaxInt32", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: math.MaxInt64", src: uint64(math.MaxUint64), expect: math.MaxUint64},
		//float
		{name: "src: math.MaxFloat32", src: math.MaxFloat32, expect: math.MaxFloat32},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: math.MaxFloat64},
		{name: "src: 1234.5678", src: 1234.5678, expect: 1234.5678},
		//complex
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//bool
		{src: true, expect: 1},
		{src: false, expect: 0},
		//string
		{name: "src: 3.4028234663852886E38", src: "3.4028234663852886E38", expect: math.MaxFloat32},   //max int64
		{name: "src: 1.7976931348623157E308", src: "1.7976931348623157E308", expect: math.MaxFloat64}, //max int64
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: math.MaxInt64},         //max int64
		{name: "src: 9223372036854775808", src: "-9223372036854775808", expect: math.MinInt64},        //min int64
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: math.MaxUint64},      //overflow
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 3.14159265358979323846264338327950288419716939937510582097494459},
		{name: "src: space 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToFloat(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToFloat32(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect float32
		err    error
	}{
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: math.MaxInt", src: math.MaxInt, expect: math.MaxInt},
		{name: "src: math.MinInt", src: math.MinInt, expect: math.MinInt},
		{name: "src: math.MaxInt8", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: math.MinInt8", src: int8(math.MinInt8), expect: math.MinInt8},
		{name: "src: math.MaxInt16", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: math.MinInt16", src: int16(math.MinInt16), expect: math.MinInt16},
		{name: "src: math.MaxInt32", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: math.MinInt32", src: int32(math.MinInt32), expect: math.MinInt32},
		{name: "src: math.MaxInt64", src: int64(math.MaxInt64), expect: math.MaxInt64},
		{name: "src: math.MinInt64", src: int64(math.MinInt64), expect: math.MinInt64},
		//uint
		{name: "src: math.MaxInt", src: uint(math.MaxUint), expect: math.MaxUint},
		{name: "src: math.MaxInt8", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: math.MaxInt16", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: math.MaxInt32", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: math.MaxInt64", src: uint64(math.MaxUint64), expect: math.MaxUint64},
		//float
		{name: "src: math.MaxFloat32", src: math.MaxFloat32, expect: math.MaxFloat32},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: 0, err: ErrOutOfRange},
		{name: "src: 1234.5678", src: 1234.5678, expect: 1234.5678},
		//complex
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//bool
		{src: true, expect: 1},
		{src: false, expect: 0},
		//string
		{name: "src: 3.4028234663852886E38", src: "3.4028234663852886E38", expect: math.MaxFloat32},         //max int64
		{name: "src: 1.7976931348623157E308", src: "1.7976931348623157E308", expect: 0, err: ErrOutOfRange}, //max int64
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: math.MaxInt64},               //max int64
		{name: "src: 9223372036854775808", src: "-9223372036854775808", expect: math.MinInt64},              //min int64
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: math.MaxUint64},            //overflow
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 3.14159265358979323846264338327950288419716939937510582097494459},
		{name: "src: space 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToFloat32(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}

func TestToFloat64(t *testing.T) {
	cases := []struct {
		name   string
		src    interface{}
		expect float64
		err    error
	}{
		{name: "src: 0", src: "0", expect: 0},
		{name: "src: nil", src: nil, expect: 0, err: ErrEmptyValue},
		//int
		{name: "src: math.MaxInt", src: math.MaxInt, expect: math.MaxInt},
		{name: "src: math.MinInt", src: math.MinInt, expect: math.MinInt},
		{name: "src: math.MaxInt8", src: int8(math.MaxInt8), expect: math.MaxInt8},
		{name: "src: math.MinInt8", src: int8(math.MinInt8), expect: math.MinInt8},
		{name: "src: math.MaxInt16", src: int16(math.MaxInt16), expect: math.MaxInt16},
		{name: "src: math.MinInt16", src: int16(math.MinInt16), expect: math.MinInt16},
		{name: "src: math.MaxInt32", src: int32(math.MaxInt32), expect: math.MaxInt32},
		{name: "src: math.MinInt32", src: int32(math.MinInt32), expect: math.MinInt32},
		{name: "src: math.MaxInt64", src: int64(math.MaxInt64), expect: math.MaxInt64},
		{name: "src: math.MinInt64", src: int64(math.MinInt64), expect: math.MinInt64},
		//uint
		{name: "src: math.MaxInt", src: uint(math.MaxUint), expect: math.MaxUint},
		{name: "src: math.MaxInt8", src: uint8(math.MaxUint8), expect: math.MaxUint8},
		{name: "src: math.MaxInt16", src: uint16(math.MaxUint16), expect: math.MaxUint16},
		{name: "src: math.MaxInt32", src: uint32(math.MaxUint32), expect: math.MaxUint32},
		{name: "src: math.MaxInt64", src: uint64(math.MaxUint64), expect: math.MaxUint64},
		//float
		{name: "src: math.MaxFloat32", src: math.MaxFloat32, expect: math.MaxFloat32},
		{name: "src: math.MaxFloat64", src: math.MaxFloat64, expect: math.MaxFloat64},
		{name: "src: 1234.5678", src: 1234.5678, expect: 1234.5678},
		//complex
		{name: "src: complex64(math.MaxInt64 + 9223372036854775807i)", src: complex64(math.MaxInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 + 9223372036854775807i)", src: complex64(math.MinInt64 + 9223372036854775807i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MaxInt64 - 9223372036854775808i)", src: complex64(math.MaxInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: complex64(math.MinInt64 - 9223372036854775808i)", src: complex64(math.MinInt64 - 9223372036854775808i), expect: 0, err: ErrDataTypeNotSupported},
		//bool
		{src: true, expect: 1},
		{src: false, expect: 0},
		//string
		{name: "src: 3.4028234663852886E38", src: "3.4028234663852886E38", expect: math.MaxFloat32},   //max int64
		{name: "src: 1.7976931348623157E308", src: "1.7976931348623157E308", expect: math.MaxFloat64}, //max int64
		{name: "src: 9223372036854775807", src: "9223372036854775807", expect: math.MaxInt64},         //max int64
		{name: "src: 9223372036854775808", src: "-9223372036854775808", expect: math.MinInt64},        //min int64
		{name: "src: 18446744073709551615", src: "18446744073709551615", expect: math.MaxUint64},      //overflow
		{name: "src: PI", src: "3.14159265358979323846264338327950288419716939937510582097494459", expect: 3.14159265358979323846264338327950288419716939937510582097494459},
		{name: "src: space 1", src: " 1", expect: 0, err: ErrInvalidSyntax},
		{name: "src: space 3.14", src: " 3.14", expect: 0, err: ErrInvalidSyntax},
		{name: "src: empty", src: "", expect: 0, err: nil},
		{name: "src: string", src: "string", expect: 0, err: ErrInvalidSyntax},
		//not supported
		{name: "src: []rune", src: []rune("to string"), expect: 0, err: ErrDataTypeNotSupported},
		{name: "src: []string", src: []string{"a", "b", "c", "d"}, expect: 0, err: ErrDataTypeNotSupported},
	}

	for _, c := range cases {
		get, err := ToFloat64(c.src)
		if c.err != nil {
			assert.ErrorContainsf(t, err, c.err.Error(), "test name %s --- source value %v", c.name, c.src)
		} else {
			assert.NoErrorf(t, err, "test name %s --- source value %v", c.name, c.src)
		}
		assert.Equalf(t, c.expect, get, "test name %s --- source value %v", c.name, c.src)
	}
}
