package utils

import (
	"crypto/md5"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

var jwtSecret []byte

func InitJwtSecret(secret string) {
	jwtSecret = []byte(secret)
}

// 将str md5加密
func Md5Str(str string) string {
	data := []byte(str)
	sum := md5.Sum(data)
	md5Str := hex.EncodeToString(sum[:])
	return md5Str

}

// 创建token  id 用户id name 用户名 expirSeconds 过期的秒数
func CreateToke(id uint, name string, expirSeconds uint) (string, error) {
	expirTime := time.Now().Add(time.Duration(expirSeconds) * time.Second)
	claims := &Claims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "yl",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)

}

type Date time.Time

// Claims 自定义的声明
type Claims struct {
	UserId   uint
	UserName string
	jwt.RegisteredClaims
}

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

// 解析token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
