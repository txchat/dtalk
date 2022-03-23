package util

import (
	"testing"
	"time"
)

func TestLocation(t *testing.T) {
	chongqing := Location("Asia/Chongqing")
	t.Log(TimeNowDateTime(chongqing))
	t.Log(TimeNowDateTime(Shanghai()))
	t.Log(TimeNowDateTime(HongKong()))
	t.Log(TimeNowDateTime(Local()))
	t.Log(TimeNowDateTime(UTC()))
}

func TestNowUnix(t *testing.T) {
	t.Log(TimeNowUnix())
	t.Log(TimeNowUnixMilli())
	t.Log(TimeNowUnixMicro())
	t.Log(TimeNowUnixNano())
}

func TestNowFormat(t *testing.T) {
	t.Log(TimeNowDate())
	t.Log(TimeNowDateTime())
	t.Log(TimeNowFormat("2006/01/02 15:04:05"))
}

func TestUnixToTime(t *testing.T) {
	nowUnixSecond := TimeNowUnix()
	nowUnixMillisecond := TimeNowUnixMilli()
	nowUnixMicrosecond := TimeNowUnixMicro()
	nowUnixNanosecond := TimeNowUnixNano()

	t.Log(nowUnixSecond, UnixToTime(nowUnixSecond))
	t.Log(nowUnixMillisecond, UnixToTime(nowUnixMillisecond))
	t.Log(nowUnixMicrosecond, UnixToTime(nowUnixMicrosecond))
	t.Log(nowUnixNanosecond, UnixToTime(nowUnixNanosecond))
}

func TestUnixAddDate(t *testing.T) {
	now := TimeNowUnix()
	t.Log(UnixAddDate(now, 1, 1, 1))
	t.Log(UnixAddYears(now, 1))
	t.Log(UnixAddMonths(now, 1))
	t.Log(UnixAddDays(now, 1))
}

func TestUnixCompare(t *testing.T) {
	now := TimeNowUnix()
	nowAdd := UnixAddDays(now, 1).In(Shanghai()).Unix()
	if !UnixBefore(now, nowAdd) {
		t.Error(now, nowAdd)
	}
	if UnixAfter(now, nowAdd) {
		t.Error(now, nowAdd)
	}
	if UnixEqual(now, nowAdd) {
		t.Error(now, nowAdd)
	}
	if !UnixEqual(now, now) {
		t.Error(now, now)
	}
	if !UnixEqual(nowAdd, nowAdd) {
		t.Error(nowAdd, nowAdd)
	}
}

func TestUnixDiffer(t *testing.T) {
	now := TimeNowUnix()
	nowAdd := UnixAddDate(now, 1, 1, 1).In(Shanghai()).Unix()
	t.Log(UnixDifferDays(nowAdd, now))
	t.Log(UnixDifferHours(nowAdd, now))
}

func TestSleep(t *testing.T) {
	Sleep(1)
	t.Log("sleep 1 second")
	SleepMilli(1)
	t.Log("sleep 1 millisecond")
	SleepMicro(1)
	t.Log("sleep 1 microsecond")
}

func TestStringToTime(t *testing.T) {
	s2t, _ := StringToTime("2020-06-01 00:00:00", "2006-01-02 15:04:05")
	t.Log(s2t)
	s2t, _ = StringToTime("2020/06/18 18:00:00", "2006/01/02 15:04:05")
	t.Log(s2t)
	s2t, _ = StringToTime("20200618180000", "20060102150405")
	t.Log(s2t)
}

func TestStringToUnix(t *testing.T) {
	t.Log(StringToUnix("20091225091010", "20060102150405"))
	t.Log(StringToUnix("2020/06/18 18:00:00", "2006/01/02 15:04:05"))
	t.Log(StringToUnix("2020-06-18 18:00:00", "2006-01-02 15:04:05"))
}

func TestTimeUnixStampToFormatString(t *testing.T) {
	t.Log(UnixToTime(time.Now().Add(time.Minute * 7).Unix()).In(Shanghai()).Format("2006-01-02 15:04:05"))
}
