package repositorie

import (
	"go-bk/internal/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func CreateUser(user *model.User) error {
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
