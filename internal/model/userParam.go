package model

import "time"

type RegisterUser struct {
	Username string `form:"username" json:"username" binding:"required,max=20"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Age      int    `form:"age" json:"age" binding:"required,gt=0"`
	Sex      uint8  `form:"sex" json:"sex" binding:"required"`
	Birthday time.Time
}
