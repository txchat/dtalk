package util

import (
	"errors"
	"time"
)

var (
	shanghai = Location("Asia/Shanghai")  // Shanghai *time.Location
	hongkong = Location("Asia/Hong_Kong") // Hong Kong *time.Location
	local    = Location("Local")          // Local *time.Location
	utc      = Location("UTC")            // UTC *time.Location
)

// Location returns *time.Location by location name.
func Location(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}

// Shanghai returns Shanghai *time.Location.
func Shanghai() *time.Location {
	return shanghai
}

// HongKong returns Hong Kong *time.Location.
func HongKong() *time.Location {
	return hongkong
}

// Local returns Local *time.Location.
func Local() *time.Location {
	return local
}

// UTC returns UTC *time.Location.
func UTC() *time.Location {
	return utc
}

// TimeNowDate returns a date representation of now time value.
func TimeNowDate(location ...*time.Location) string {
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	return time.Now().In(loc).Format("2006-01-02")
}

// TimeNowDateTime returns a datetime representation of now time value.
func TimeNowDateTime(location ...*time.Location) string {
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

// TimeNowFormat returns a textual representation of now time value formatted according to layout.
func TimeNowFormat(layout string, location ...*time.Location) string {
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	return time.Now().In(loc).Format(layout)
}

// TimeNowUnix returns now unix second timestamp.
func TimeNowUnix(location ...*time.Location) int64 {
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	return time.Now().In(loc).Unix()
}

// TimeNowUnixMilli returns now unix millisecond timestamp.
func TimeNowUnixMilli(location ...*time.Location) int64 {
	return TimeNowUnixNano(location...) / int64(time.Millisecond)
}

// TimeNowUnixMicro returns now unix microsecond timestamp.
func TimeNowUnixMicro(location ...*time.Location) int64 {
	return TimeNowUnixNano(location...) / int64(time.Microsecond)
}

// TimeNowUnixNano returns now unix nanosecond timestamp.
func TimeNowUnixNano(location ...*time.Location) int64 {
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	return time.Now().In(loc).UnixNano()
}

// UnixToTime returns time.Time by unix timestamp.
func UnixToTime(timestamp int64) time.Time {
	if timestamp < 1e10 {
		return time.Unix(timestamp, 0)
	} else if timestamp < 1e13 {
		return time.Unix(0, timestamp*int64(time.Millisecond))
	} else if timestamp < 1e16 {
		return time.Unix(0, timestamp*int64(time.Microsecond))
	} else {
		return time.Unix(0, timestamp)
	}
}

// UnixAddDate returns time.Time after unix timestamp has been added date.
func UnixAddDate(timestamp int64, years, months, days int) time.Time {
	t := UnixToTime(timestamp)
	return t.AddDate(years, months, days)
}

// UnixAddYears returns time.Time after unix timestamp has been added years.
func UnixAddYears(timestamp int64, years int) time.Time {
	return UnixAddDate(timestamp, years, 0, 0)
}

// UnixAddMonths returns time.Time after unix timestamp has been added months.
func UnixAddMonths(timestamp int64, months int) time.Time {
	return UnixAddDate(timestamp, 0, months, 0)
}

// UnixAddDays returns time.Time after unix timestamp has been added days.
func UnixAddDays(timestamp int64, days int) time.Time {
	return UnixAddDate(timestamp, 0, 0, days)
}

// UnixEqual reports whether timestamp1 is equal timestamp2.
func UnixEqual(timestamp1 int64, timestamp2 int64) bool {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.Equal(t2)
}

// UnixBefore reports whether timestamp1 is before timestamp2.
func UnixBefore(timestamp1 int64, timestamp2 int64) bool {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.Before(t2)
}

// UnixAfter reports whether timestamp1 is after timestamp2.
func UnixAfter(timestamp1 int64, timestamp2 int64) bool {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.After(t2)
}

// UnixDifferDays returns the number of days between two timestamp.
func UnixDifferDays(timestamp1 int64, timestamp2 int64) int {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return int(t1.Sub(t2).Hours() / 24)
}

// UnixDifferHours returns the number of hours between two timestamp.
func UnixDifferHours(timestamp1 int64, timestamp2 int64) float64 {
	t1, t2 := UnixToTime(timestamp1), UnixToTime(timestamp2)
	return t1.Sub(t2).Hours()
}

// StringToTime returns time.Time representation of str value parsed according to layout.
func StringToTime(str, format string, location ...*time.Location) (time.Time, error) {
	// format example:
	//     20060102150405
	//     2006-01-02 15:04:05
	//     2006/01/02 15/04/05
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	if len(str) != len(format) {
		return time.Now(), errors.New("input does not match format")
	}
	return time.ParseInLocation(format, str, loc)
}

// StringToUnix returns unix second timestamp representation of str value parsed according to layout.
// If str parsed err, it returns now unix second timestamp.
func StringToUnix(str, format string, location ...*time.Location) int64 {
	// format example:
	//     20060102150405
	//     2006-01-02 15:04:05
	//     2006/01/02 15/04/05
	loc := Shanghai()
	if len(location) != 0 {
		loc = location[0]
	}
	t, err := StringToTime(str, format, loc)
	if err != nil {
		return TimeNowUnix()
	}
	return t.In(loc).Unix()
}

// Sleep pauses the current goroutine for at least n second.
func Sleep(n int64) {
	time.Sleep(time.Duration(n) * time.Second)
}

// SleepMilli pauses the current goroutine for at least n millisecond.
func SleepMilli(n int64) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// SleepMicro pauses the current goroutine for at least n microsecond.
func SleepMicro(n int64) {
	time.Sleep(time.Duration(n) * time.Microsecond)
}

func StringToTimeDefault(str string) time.Time {
	tm, err := StringToTime(str, "2006-01-02 15:04:05")
	if err != nil {
		return time.Time{}
	}
	return tm
}
