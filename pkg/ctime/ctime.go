package ctime

import (
	"database/sql/driver"
	"fmt"
	"myblog/pkg/config"

	"time"
)

//
// CTime 自定义 time 类型
//
// 由于: Go 自身的 time.Time 类型默认解析的日期格式是 RFC3339 标准， 也就是 2006-01-02T15:04:05Z07:00 的格式。
//      所以需要自定义 time 类型, 然后重写 json 的序列化和反序列化方法
//      implements the json.Marshaler interface.
//      implements the json.Unmarshaler interface.
//
// 1. 在使用 gorm 框架时，用到 MarshalJSON / Scan / Value 方法。
// 2. 在使用 Gin 框架时使用 ShouldBindJSON 绑定参数会调用 UnmarshalJSON，
//    在 context.ShouldBindJSON 时，会调用 field.UnmarshalJSON方法，
//    context.JSON 返回 json 时，会调用 field.MarshalJSON 方法。
// 3. 在使用 go-redis 从缓存里取 2021-08-26 12:00:00 格式的字符串时会调用 UnmarshalJSON
//    不需要重新定义 string 类型的字段，解析到缓存数据之后再转成自定义 time 类型。
//
type CTime struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (t CTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(""), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(config.DateTimeLayout))
	return []byte(formatted), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *CTime) UnmarshalJSON(data []byte) error {
	if len(data) == 2 {
		*t = CTime{Time: time.Time{}}
		return nil
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}
	now, err := time.ParseInLocation(`"`+config.DateTimeLayout+`"`, string(data), location)
	if err != nil {
		return err
	}
	*t = CTime{Time: now}

	return nil
}

// Value insert timestamp into mysql need this function.
// 实现 /src/database/sql/driver/types.go 中的 "driver.Valuer" 接口
func (t CTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
// 实现 /src/database/sql/sql.go 中的 "sql.Scanner" 接口
func (t *CTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = CTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

