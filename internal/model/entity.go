package model

import (
	"fmt"
	"go-bk/utils"
	"gorm.io/gorm"
)

// User 这个是用户的结构体
type User struct {
	ID       uint
	Name     string
	Email    string
	Age      uint8
	Birthday utils.Date
	Sex      *uint8
	Password string
	gorm.Model
}

// Post 这个是文章的结构体
type Post struct {
	ID uint
	//标题
	Title string
	//内容
	Content string
	//关联的用户id
	RefUserID uint

	//评论状态
	IsComment bool
	gorm.Model
}
type PostComment struct {
	ID uint
	//内容
	Comment string
	//文章id
	PostId int

	UserId uint

	gorm.Model
}

func (User) TableName() string {
	return "user"
}

// TableName TableName()方法返回数据库表名
func (Post) TableName() string {
	return "post"
}

func (PostComment) TableName() string {
	return "post_comment"
}

// AfterCreate 钩子函数创建文章之后
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	postCount, ok := tx.InstanceGet("postCount")
	postCountNum := postCount.(uint64)
	//成功
	if ok {
		fmt.Printf("postCount==%d", postCountNum)
		postCountNum++
		tx.Model(&User{}).Where(&User{ID: p.RefUserID}).Update("post_count", postCountNum)
	}

	return
}

// 钩子函数删除文章之后  更新用户发表的文章数
func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	usrId := p.RefUserID
	var postCount int64
	user := &User{ID: usrId}
	tx.Debug().Model(&User{}).Select("post_count").
		Where(user).Find(&postCount)
	if postCount > 0 {
		postCount--
		tx.Model(&User{}).Where(&User{ID: usrId}).Update("post_count", postCount)
	}

	return
}
