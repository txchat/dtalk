package http

import (
	"fmt"
	"reflect"
	"strconv"
)

func ParseInterface(orign interface{}, ty string) (interface{}, error) {
	var result interface{}

	switch ty {
	case "":
		return nil, fmt.Errorf("invalid ty")
	case "int":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = int(tOrign)
		case uint:
			result = int(tOrign)
		case int32:
			result = int(tOrign)
		case uint32:
			result = int(tOrign)
		case int64:
			result = int(tOrign)
		case uint64:
			result = int(tOrign)
		case float32:
			result = int(tOrign)
		case float64:
			result = int(tOrign)
		case string:
			tm, err := strconv.ParseInt(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = int(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "uint":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = uint(tOrign)
		case uint:
			result = uint(tOrign)
		case int32:
			result = uint(tOrign)
		case uint32:
			result = uint(tOrign)
		case int64:
			result = uint(tOrign)
		case uint64:
			result = uint(tOrign)
		case float32:
			result = uint(tOrign)
		case float64:
			result = uint(tOrign)
		case string:
			tm, err := strconv.ParseUint(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = uint(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "int32":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = int32(tOrign)
		case uint:
			result = int32(tOrign)
		case int32:
			result = int32(tOrign)
		case uint32:
			result = int32(tOrign)
		case int64:
			result = int32(tOrign)
		case uint64:
			result = int32(tOrign)
		case float32:
			result = int32(tOrign)
		case float64:
			result = int32(tOrign)
		case string:
			tm, err := strconv.ParseInt(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = int32(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "uint32":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = uint32(tOrign)
		case uint:
			result = uint32(tOrign)
		case int32:
			result = uint32(tOrign)
		case uint32:
			result = uint32(tOrign)
		case int64:
			result = uint32(tOrign)
		case uint64:
			result = uint32(tOrign)
		case float32:
			result = uint32(tOrign)
		case float64:
			result = uint32(tOrign)
		case string:
			tm, err := strconv.ParseUint(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = uint32(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "int64":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = int64(tOrign)
		case uint:
			result = int64(tOrign)
		case int32:
			result = int64(tOrign)
		case uint32:
			result = int64(tOrign)
		case int64:
			result = int64(tOrign)
		case uint64:
			result = int64(tOrign)
		case float32:
			result = int64(tOrign)
		case float64:
			result = int64(tOrign)
		case string:
			tm, err := strconv.ParseInt(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = int64(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "uint64":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = uint64(tOrign)
		case uint:
			result = uint64(tOrign)
		case int32:
			result = uint64(tOrign)
		case uint32:
			result = uint64(tOrign)
		case int64:
			result = uint64(tOrign)
		case uint64:
			result = uint64(tOrign)
		case float32:
			result = uint64(tOrign)
		case float64:
			result = uint64(tOrign)
		case string:
			tm, err := strconv.ParseUint(tOrign, 10, 64)
			if nil != err {
				return nil, err
			}
			result = uint64(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "float32":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = float32(tOrign)
		case uint:
			result = float32(tOrign)
		case int32:
			result = float32(tOrign)
		case uint32:
			result = float32(tOrign)
		case int64:
			result = float32(tOrign)
		case uint64:
			result = float32(tOrign)
		case float32:
			result = float32(tOrign)
		case float64:
			result = float32(tOrign)
		case string:
			tm, err := strconv.ParseFloat(tOrign, 32)
			if nil != err {
				return nil, err
			}
			result = float32(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "float64":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = float64(tOrign)
		case uint:
			result = float64(tOrign)
		case int32:
			result = float64(tOrign)
		case uint32:
			result = float64(tOrign)
		case int64:
			result = float64(tOrign)
		case uint64:
			result = float64(tOrign)
		case float32:
			result = float64(tOrign)
		case float64:
			result = float64(tOrign)
		case string:
			tm, err := strconv.ParseFloat(tOrign, 64)
			if nil != err {
				return nil, err
			}
			result = float64(tm)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "string":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case int:
			result = fmt.Sprint(uint64(tOrign))
		case uint:
			result = fmt.Sprint(uint64(tOrign))
		case int32:
			result = fmt.Sprint(uint64(tOrign))
		case uint32:
			result = fmt.Sprint(uint64(tOrign))
		case int64:
			result = fmt.Sprint(uint64(tOrign))
		case uint64:
			result = fmt.Sprint(uint64(tOrign))
		case float32:
			// 这种只适合整数转字符串的情形, 也就是id那种情况, 带小数的转换不支持
			if float64(tOrign) > float64(uint64(tOrign)) {
				return nil, fmt.Errorf("not support the condition, float64(tOrign) > float64(uint64(tOrign))")
			}

			result = fmt.Sprint(uint64(tOrign))
		case float64:
			// 这种只适合整数转字符串的情形, 也就是id那种情况, 带小数的转换不支持
			if tOrign > float64(uint64(tOrign)) {
				return nil, fmt.Errorf("not support the condition, tOrign > float64(uint64(tOrign))")
			}
			result = fmt.Sprint(uint64(tOrign))
		case string:
			result = tOrign
		case bool:
			result = fmt.Sprint(tOrign)
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	case "bool":
		switch tOrign := orign.(type) {
		case nil:
			return nil, fmt.Errorf("tOrign is nil")
		case bool:
			result = tOrign
		case string:
			tm, err := strconv.ParseBool(tOrign)
			if nil != err {
				return nil, err
			}
			result = tm
		default:
			return nil, fmt.Errorf("unknow tOrign type, " + reflect.TypeOf(tOrign).String())
		}
	default:
		return nil, fmt.Errorf("unknow ty")
	}

	return result, nil
}
