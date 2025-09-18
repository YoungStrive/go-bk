package model

import (
	"go-bk/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Age      uint8
	Birthday utils.Date
	Sex      uint8
	Password string
	gorm.Model
}

func (User) TableName() string {
	return "user"
}
