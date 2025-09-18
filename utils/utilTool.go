package utils

import (
	"crypto/md5"
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

// 将str md5加密
func Md5Str(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	md5Str := hex.EncodeToString(sum[:])
	return md5Str

}

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(d).Format(DateFormat))), nil
}
func (d *Date) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+DateFormat+`"`, string(b), time.Local)
	if err != nil {
		return fmt.Errorf("can not convert %v to date,must like format:yyyy-MM-dd,simple example : %v", string(b), DateFormat)
	}
	*d = Date(now)
	return nil
}
func (d Date) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(d).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(d), nil
}
func (d *Date) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*d = Date(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (d Date) String() string {
	return time.Time(d).Format(DateFormat)
}

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Time(t).Format(TimeFormat))), nil
}
func (t *DateTime) UnmarshalJSON(b []byte) error {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(b), time.Local)
	if err != nil {
		return fmt.Errorf("can not convert %v to date,must like format:yyyy-MM-dd HH:mm:ss,simple example : %v", string(b), TimeFormat)
	}
	*t = DateTime(now)
	return nil
}
func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}
func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (t DateTime) String() string {
	return time.Time(t).Format(TimeFormat)
}
