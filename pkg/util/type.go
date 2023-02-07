package util

import (
	"encoding/json"
	"errors"
	"math"
	"runtime/debug"
	"strconv"
)

const (
	maxFloat32UintBit = 0x01 << 25
	maxFloat64UintBit = 0x01 << 53
)

var (
	ErrMissingPrecision     = errors.New("missing precision")
	ErrOutOfRange           = errors.New("value out of range")
	ErrInvalidSyntax        = errors.New("invalid syntax")
	ErrOverflow             = errors.New("")
	ErrEmptyValue           = errors.New("convert source data is empty")
	ErrDataTypeNotSupported = errors.New("data type not supported")
)

func MustToInt(src interface{}) int {
	return int(MustToInt64(src))
}

func MustToInt32(src interface{}) int32 {
	return int32(MustToInt64(src))
}

func MustToInt64(src interface{}) int64 {
	v, err := ToInt64(src)
	if err != nil {
		debug.PrintStack()
		panic(err)
		//.Fatal("string parse int err", err)
	}
	return v
}

func MustToUint(scr interface{}) uint {
	return uint(MustToInt64(scr))
}

func MustToUint32(scr interface{}) uint32 {
	return uint32(MustToInt64(scr))
}

func MustToUint64(scr interface{}) uint64 {
	v, err := ToUint64(scr)
	if err != nil {
		debug.PrintStack()
		panic(err)
		//.Fatal("string parse int err", err)
	}
	return v
}

func ToInt(src interface{}) (int, error) {
	v, err := ToInt64(src)
	return int(v), err
}

func ToInt32(src interface{}) (int32, error) {
	v, err := ToInt64(src)
	if v > math.MaxInt32 || v < math.MinInt32 {
		return 0, ErrOutOfRange
	}
	return int32(v), err
}

func ToInt64(src interface{}) (int64, error) {
	if src == nil {
		return 0, ErrEmptyValue
	}
	switch v := src.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if v > math.MaxInt64 {
			return 0, ErrOutOfRange
		}
		return int64(v), nil
	case uintptr:
		return int64(v), nil
	case float32:
		tmp := math.Floor(float64(v))
		if tmp != float64(v) {
			return 0, ErrInvalidSyntax
		}
		if tmp > math.MaxInt64 || tmp < math.MinInt64 {
			return 0, ErrOutOfRange
		}
		if tmp > maxFloat32UintBit {
			return 0, ErrMissingPrecision
		}
		return int64(v), nil
	case float64:
		tmp := math.Floor(v)
		if tmp != v {
			return 0, ErrInvalidSyntax
		}
		if tmp > math.MaxInt64 || tmp < math.MinInt64 {
			return 0, ErrOutOfRange
		}
		if tmp > maxFloat64UintBit {
			return 0, ErrMissingPrecision
		}
		return int64(v), nil
	case complex64:
		return 0, ErrDataTypeNotSupported
	case complex128:
		return 0, ErrDataTypeNotSupported
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if v == "" {
			return 0, nil
		}
		tmp, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return tmp, nil
	case json.Number:
		tmp, err := v.Int64()
		if err != nil {
			return 0, err
		}
		return tmp, nil
	default:
		return 0, ErrDataTypeNotSupported
	}
}

func ToUint(src interface{}) (uint, error) {
	v, err := ToUint64(src)
	return uint(v), err
}

func ToUint32(src interface{}) (uint32, error) {
	v, err := ToUint64(src)
	if v > math.MaxUint32 {
		return 0, ErrOutOfRange
	}
	return uint32(v), err
}

func ToUint64(src interface{}) (uint64, error) {
	if src == nil {
		return 0, ErrEmptyValue
	}
	switch v := src.(type) {
	case int:
		if v < 0 {
			return 0, ErrOutOfRange
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, ErrOutOfRange
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, ErrOutOfRange
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, ErrOutOfRange
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, ErrOutOfRange
		}
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case uintptr:
		return uint64(v), nil
	case float32:
		tmp := math.Floor(float64(v))
		if tmp != float64(v) {
			return 0, ErrInvalidSyntax
		}
		if tmp > math.MaxUint64 || tmp < 0 {
			return 0, ErrOutOfRange
		}
		if tmp > maxFloat32UintBit {
			return 0, ErrMissingPrecision
		}
		return uint64(v), nil
	case float64:
		tmp := math.Floor(v)
		if tmp != v {
			return 0, ErrInvalidSyntax
		}
		if tmp > math.MaxUint64 || tmp < 0 {
			return 0, ErrOutOfRange
		}
		if tmp > maxFloat64UintBit {
			return 0, ErrMissingPrecision
		}
		return uint64(v), nil
	case complex64:
		return 0, ErrDataTypeNotSupported
	case complex128:
		return 0, ErrDataTypeNotSupported
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if v == "" {
			return 0, nil
		}
		tmp, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return tmp, nil
	case json.Number:
		tmp, err := strconv.ParseUint(string(v), 10, 64)
		if err != nil {
			return 0, err
		}
		return tmp, nil
	default:
		return 0, ErrDataTypeNotSupported
	}
}

// ToFloat returns the float64 result converted by src.
func ToFloat(src interface{}) (float64, error) {
	return ToFloat64(src)
}

// ToFloat32 returns the float32 result converted by src.
func ToFloat32(src interface{}) (float32, error) {
	v, err := ToFloat64(src)
	if v > math.MaxFloat32 {
		return 0, ErrOutOfRange
	}
	return float32(v), err
}

func ToFloat64(src interface{}) (float64, error) {
	if src == nil {
		return 0, ErrEmptyValue
	}

	switch v := src.(type) {
	case int, int8, int16, int32, int64:
		tmp, err := ToInt64(v)
		return float64(tmp), err
	case uint, uint8, uint16, uint32, uint64, uintptr:
		tmp, err := ToUint64(v)
		return float64(tmp), err
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case complex64:
		return 0, ErrDataTypeNotSupported
	case complex128:
		return 0, ErrDataTypeNotSupported
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		if v == "" {
			return 0, nil
		}
		return strconv.ParseFloat(v, 64)
	case json.Number:
		return v.Float64()
	default:
		return 0, ErrDataTypeNotSupported
	}
}

func MustToString(src interface{}) string {
	v, err := ToString(src)
	if err != nil {
		debug.PrintStack()
		panic(err)
	}
	return v
}

func MustToBool(src interface{}) bool {
	v, err := ToBool(src)
	if err != nil {
		debug.PrintStack()
		panic(err)
	}
	return v
}

// ToString returns the string result converted by src.
func ToString(src interface{}) (string, error) {
	switch v := src.(type) {
	case int, int8, int16, int32, int64:
		tmp, err := ToInt64(v)
		return strconv.FormatInt(tmp, 10), err
	case uint, uint8, uint16, uint32, uint64, uintptr:
		tmp, err := ToUint64(v)
		return strconv.FormatUint(tmp, 10), err
	case float32, float64, complex64, complex128:
		tmp, err := ToFloat64(v)
		return strconv.FormatFloat(tmp, 'f', -1, 64), err
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case []rune:
		return string(v), nil
	case bool:
		return strconv.FormatBool(v), nil
	case nil:
		return "", ErrEmptyValue
	default:
		return "", ErrDataTypeNotSupported
	}
}

// ToBool returns the bool result converted by src.
func ToBool(src interface{}) (bool, error) {
	switch v := src.(type) {
	case int, int8, int16, int32, int64:
		tmp, err := ToInt64(v)
		return tmp > 0, err
	case uint, uint8, uint16, uint32, uint64, uintptr:
		tmp, err := ToUint64(v)
		return tmp > 0, err
	case float32, float64, complex64, complex128:
		tmp, err := ToFloat64(v)
		return tmp > 0, err
	case bool:
		return v, nil
	case string, []byte, []rune:
		s, err := ToString(v)
		if err != nil {
			return false, err
		}
		return strconv.ParseBool(s)
	default:
		return false, ErrDataTypeNotSupported
	}
}
