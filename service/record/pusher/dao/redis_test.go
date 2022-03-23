package dao

import (
	"testing"
)

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
