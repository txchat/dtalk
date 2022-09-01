package util

import (
	"encoding/hex"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/inconshreveable/log15"

	"github.com/haltingstate/secp256k1-go"
)

const (
	Invalid   = 0
	SECP256K1 = 1
	ED25519   = 2
	SM2       = 3
)

//tm 时间，单位毫秒
func CheckTimeOut(tm int64, timeout time.Duration) bool {
	sec := tm / 1000
	nsec := tm % 1000
	t := time.Unix(sec, nsec).Add(timeout)
	return time.Now().After(t)
}

//func CheckTimeOut(tm int64, timeout time.Duration) bool {
//	sec := tm / 1000
//	nsec := tm % 1000
//	t := time.Unix(sec, nsec)
//	tExp := t.Add(timeout)
//	return time.Now().After(tExp) || time.Now().Before(t)
//}

//得到摘要
func GetSummary(i interface{}) ([]byte, error) {
	return summaryNoEncode(i)
}

//得到摘要
func summaryUrlEncode(i interface{}) ([]byte, error) {
	params := paramsMap(i)
	delete(params, "sign")
	delete(params, "-")

	u := url.Values{}
	for k, v := range params {
		u.Set(k, v)
	}
	//按字典升序排序并URL编码
	secretStr := u.Encode()

	return []byte(secretStr), nil
}

func summaryNoEncode(i interface{}) ([]byte, error) {
	params := paramsMap(i)
	delete(params, "sign")
	delete(params, "-")

	u := url.Values{}
	for k, v := range params {
		u.Set(k, v)
	}

	var buf strings.Builder
	keys := make([]string, 0, len(u))
	for k := range u {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := u[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}

	//按字典升序排序
	secretStr := buf.String()

	return []byte(secretStr), nil
}

func Secp256k1Verify(msg, sig, pubKey []byte) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			err, _ := r.(error)
			log15.Error("secp256k1 VerifySignature failed", "err", err)
			b = false
		}
	}()
	return 1 == secp256k1.VerifySignature(msg, sig, pubKey)
}

func paramsMap(i interface{}) map[string]string {
	ret := make(map[string]string)

	st := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	ss := st.Elem()
	vv := v.Elem()
	for i := 0; i < ss.NumField(); i++ {
		field := ss.Field(i)
		val := vv.Field(i).Interface()
		keyName := MustToString(field.Tag.Get("json"))
		keyName = strings.Replace(keyName, ",omitempty", "", 1)
		ret[keyName] = MustToString(val)
	}
	return ret
}

//兼容0x格式
func HexDecode(in string) ([]byte, error) {
	return hex.DecodeString(strings.Replace(in, "0x", "", 1))
}
