package dao

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func connSeqIndexValConvert(tp string, logs []string) (string, error) {
	val := fmt.Sprintf("type=%s;", tp)
	val += strings.Join(logs, ",")
	//check
	reg := regexp.MustCompile(`^type=(\S+);`)
	if !reg.MatchString(val) {
		return val, errors.New("conn seq index convert err")
	}
	return val, nil
}

func connSeqIndexValParse(val string) ([]string, error) {
	reg := regexp.MustCompile(`^type=(\S+);`)
	if !reg.MatchString(val) {
		return nil, errors.New("conn seq index parse err")
	}
	//找出type子串
	typeCnt := reg.FindStringSubmatch(val)
	if len(typeCnt) < 2 {
		return nil, errors.New("conn seq index parse err: cnt len < 2")
	}
	//找出;所在的尾部下标
	typeIdx := reg.FindStringIndex(val)
	if len(typeIdx) < 2 {
		return nil, errors.New("conn seq index parse err: idx len < 2")
	}
	log := strings.Split(val[typeIdx[1]:], ",")
	var ret = make([]string, len(log)+1)
	copy(ret[1:], log)
	ret[0] = typeCnt[1]
	return ret, nil
}

func Test_connSeqIndexValConvert(t *testing.T) {
	got, err := connSeqIndexValConvert("common", []string{"1", "2"})
	if err != nil {
		t.Errorf("connSeqIndexValConvert() error = %v", err)
		return
	}
	t.Logf("got=%s", got)

	got2, err2 := connSeqIndexValConvert("common", []string{})
	t.Logf("got2=%s, err= %v", got2, err2)
}

func Test_connSeqIndexValParse(t *testing.T) {
	{
		got, err := connSeqIndexValParse("type=common;1,2")
		t.Logf("index=%d,got=%s,err=%v\n", 1, got, err)
	}
	{
		got, err := connSeqIndexValParse("type=common;1")
		t.Logf("index=%d,got=%s,err=%v\n", 2, got, err)
	}
	{
		got, err := connSeqIndexValParse("type=common;")
		t.Logf("index=%d,got=%s,err=%v\n", 3, got, err)
	}
	{
		got, err := connSeqIndexValParse("type=common1,2")
		t.Logf("index=%d,got=%s,err=%v\n", 4, got, err)
	}
	{
		got, err := connSeqIndexValParse("type=;1,2")
		t.Logf("index=%d,got=%s,err=%v\n", 4, got, err)
	}
}
