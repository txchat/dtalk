package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

var (
	ErrTransToInt64 = errors.New("transfer type int64 error")
	ErrTransToInt32 = errors.New("transfer type int32 error")
)

func ToInt(val interface{}) int {
	return int(ToInt32(val))
}

func ToUInt32(val interface{}) uint32 {
	return uint32(ToInt32(val))
}

func ToUInt64(val interface{}) uint64 {
	return uint64(ToInt64(val))
}

func ToInt32(o interface{}) int32 {
	if o == nil {
		debug.PrintStack()
		log.Fatal("nil value")
		return 0
	}
	switch t := o.(type) {
	case int:
		return int32(t)
	case int32:
		return t
	case int64:
		return int32(t)
	case float64:
		return int32(t)
	case string:
		if o == "" {
			debug.PrintStack()
			log.Fatal("empty string")
			return 0
		}
		temp, err := strconv.ParseInt(o.(string), 10, 32)
		if err != nil {
			debug.PrintStack()
			log.Fatal("string parse int err", err)
			return 0
		}
		return int32(temp)
	default:
		debug.PrintStack()
		log.Fatal("unknown type", fmt.Sprintf("%T", o))
		return 0
	}
}

func ToInt32E(o interface{}) (int32, error) {
	if o == nil {
		return 0, ErrTransToInt32
	}
	switch t := o.(type) {
	case int:
		return int32(t), nil
	case int32:
		return t, nil
	case int64:
		return int32(t), nil
	case float64:
		return int32(t), nil
	case string:
		if o == "" {
			return 0, ErrTransToInt32
		}
		temp, err := strconv.ParseInt(o.(string), 10, 32)
		if err != nil {
			return 0, ErrTransToInt32
		}
		return int32(temp), nil
	default:
		return 0, ErrTransToInt32
	}
}

func ToInt64(val interface{}) int64 {
	if val == nil {
		debug.PrintStack()
		log.Fatal("nil value")
		return 0
	}
	switch val.(type) {
	case int:
		return int64(val.(int))
	case string:
		if val.(string) == "" {
			debug.PrintStack()
			log.Fatal("empty string")
			return 0
		}
		ret, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			debug.PrintStack()
			log.Fatal("string parse int err", err)
			return 0
		}
		return ret
	case float64:
		return int64(val.(float64))
	case int64:
		return val.(int64)
	case json.Number:
		v := val.(json.Number)
		ret, err := v.Int64()
		if err != nil {
			debug.PrintStack()
			log.Fatal("unknown json number")
			return 0
		}
		return ret
	default:
		debug.PrintStack()
		log.Fatal("unknown type", fmt.Sprintf("%T", val))
		return 0
	}
}

func ToInt64E(val interface{}) (int64, error) {
	if val == nil {
		return 0, ErrTransToInt64
	}
	switch val.(type) {
	case int:
		return int64(val.(int)), nil
	case string:
		if val.(string) == "" {
			return 0, ErrTransToInt64
		}
		ret, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return 0, ErrTransToInt64
		}
		return ret, nil
	case float64:
		return int64(val.(float64)), nil
	case int64:
		return val.(int64), nil
	case json.Number:
		v := val.(json.Number)
		ret, err := v.Int64()
		if err != nil {
			return 0, ErrTransToInt64
		}
		return ret, nil
	default:
		return 0, ErrTransToInt64
	}
}

func ToFloat64(val interface{}) float64 {
	if val == nil {
		debug.PrintStack()
		log.Fatal("nil value")
		return 0
	}
	switch val.(type) {
	case string:
		ret, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			debug.PrintStack()
			log.Fatal("string parse float err", err)
		}
		return ret
	default:
		if v, ok := val.(float64); ok {
			return v
		}
		debug.PrintStack()
		log.Fatal("unknown type", fmt.Sprintf("%T", val))
		return 0
	}
}

func TypeToString(val interface{}) string {
	if val == nil {
		debug.PrintStack()
		log.Fatal("nil value")
		return ""
	}
	switch val.(type) {
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case int64:
		return strconv.FormatInt(val.(int64), 10)
	}
	return fmt.Sprintf("%v", val)
}

func ToBool(val interface{}) bool {
	return ToUInt32(val) != 0
}
