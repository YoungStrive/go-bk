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
	sexMap := map[int32]string{0: "男", 1: "女"}
	userList, err := repositorie.ListUserByName(name)
	// 返回错误而不进行后续操作
	if err != nil {
		return nil, err
	}

	sexFlag := "sex"
	sexName := "未知"
	for _, obj := range userList {
		//性别的值
		sexValue, ok := obj[sexFlag]
		if ok {
			sex, ok := sexMap[sexValue.(int32)]
			if ok {
				sexName = sex
			}
		}
		obj["sexName"] = sexName
	}
	return userList, nil
}

func LoginUser(name string, pwd string) (string, error) {
	//根据名字查找用户
	user := repositorie.GetByUserName(name)
	if user == nil {
		return "", fmt.Errorf("用户不存在")
	}
	//验证密码
	if user.Password != utils.Md5Str(pwd) {
		return "", fmt.Errorf("密码错误")
	}

	tokerStr, err := utils.CreateToke(user.ID, user.Name, 60*60*24)
	return tokerStr, err

}
