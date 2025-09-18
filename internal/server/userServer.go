package server

import (
	"fmt"
	"go-bk/internal/model"
	"go-bk/internal/repositorie"
	"go-bk/utils"
)

// 创建用户
func CreateUser(user *model.RegisterUser) (*model.User, error) {
	nameCount := repositorie.GetCountByUserName(user.Username)
	if nameCount > 0 {
		return nil, fmt.Errorf("user %s 已存在", user.Username)
	}

	addUser := &model.User{
		Name:     user.Username,
		Sex:      user.Sex,
		Email:    user.Email,
		Password: utils.Md5Str(user.Password),
		Birthday: user.Birthday,
		Age:      user.Age,
	}
	repositorie.CreateUser(addUser)
	return addUser, nil
}

// 根据用户名获取到用户列表
func ListUserByName(name string) ([]map[string]interface{}, error) {
	return repositorie.ListUserByName(name)
}
