package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Age      uint8
	Birthday time.Time
	Sex      uint8
	Password string
	gorm.Model
}
