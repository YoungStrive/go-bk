package model

import (
	"go-bk/utils"
)

type RegisterUser struct {
	Username string     `form:"username" json:"username" binding:"required,max=20"`
	Password string     `form:"password" json:"password" binding:"required"`
	Email    string     `form:"email" json:"email" binding:"required,email"`
	Age      uint8      `form:"age" json:"age" binding:"required,gt=0"`
	Sex      *uint8     `form:"sex" json:"sex" binding:"required"`
	Birthday utils.Date `form:"birthday" json:"birthday" `
}

// AddPost 添加文章的结构体
type AddPost struct {
	Title   string `form:"title" json:"title" binding:"required,max=20"`
	Content string `form:"content" json:"content" binding:"required"`
	UserId  uint
}

type UpdatePost struct {
	Title   string `form:"title" json:"title" binding:"required,max=20"`
	Content string `form:"content" json:"content" binding:"required"`
	ID      int    `form:"id" json:"id" binding:"required"`
}

// 添加文章评论的结构体
type AddPostComment struct {
	PostId  int    `form:"postId" json:"postId" binding:"required"`
	Comment string `form:"comment" json:"comment" binding:"required"`
	UserId  uint
}
