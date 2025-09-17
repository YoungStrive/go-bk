package repositorie

import (
	"go-bk/internal/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(db *gorm.DB) {
	db = db
}

func CreateUser(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户名查找数据判断是否重名
func GetCountByUserName(name string) int64 {
	var nameCount int64
	db.Model(&model.User{}).Where(&model.User{Name: name}).Count(&nameCount)
	return nameCount
}
