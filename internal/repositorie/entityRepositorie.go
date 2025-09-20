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

// 根据用户名查找数据判断是否重名
func GetCountByUserName(name string) int64 {
	var nameCount int64
	DB.Model(&model.User{}).Where(&model.User{Name: name}).Count(&nameCount)
	return nameCount
}

func ListUserByName(name string) ([]map[string]interface{}, error) {
	var resultList []map[string]interface{}
	user := &model.User{}
	if name == "" {
		DB.Debug().Model(user).Select("id", "name", "sex").Scan(&resultList)
	} else {
		DB.Debug().Model(user).Select("id", "name", "sex").Where("name LIKE ? ", "%"+name+"%").Find(&resultList)
	}
	return resultList, nil
}

// 根据用户名获取用户信息
func GetByUserName(name string) *model.User {
	user := &model.User{}
	DB.Debug().Model(user).Where(&model.User{Name: name}).First(user)
	return user
}
