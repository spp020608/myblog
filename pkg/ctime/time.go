package ctime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
	timeZone = "Asia/Shanghai"
)

type Time time.Time

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	buf := make([]byte, 0, len(timeLayout))
	buf = append(buf, '"')
	buf = time.Time(t).AppendFormat(buf, timeLayout)
	buf = append(buf, '"')
	return buf, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeLayout+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) String() string {
	return time.Time(t).Format(timeLayout)
}

func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(timeZone)
	return time.Time(t).In(loc)
}

// Value insert timestamp into mysql need this function.
// 实现 /src/database/sql/driver/types.go 中的 "driver.Valuer" 接口
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan value of time.Time
// 实现 /src/database/sql/sql.go 中的 "sql.Scanner" 接口
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}