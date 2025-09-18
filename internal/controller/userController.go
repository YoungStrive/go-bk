package controller

import (
	"github.com/gin-gonic/gin"
	"go-bk/internal/model"
	"go-bk/internal/server"
	"go-bk/pkg/response"
	"net/http"
)

// 注册用户
func RegisterUser(c *gin.Context) {
	userRegister := model.RegisterUser{}
	//参数不对
	if err := c.ShouldBind(&userRegister); err != nil {
		response.Error(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	user, err := server.CreateUser(&userRegister)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 400, err.Error())
		return
	}
	response.Success(c, user)
}

// 获取所有的用户
func ListUser(c *gin.Context) {
	name := c.Query("name")
	allUser, err := server.ListUserByName(name)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, 400, err.Error())
		return
	}
	response.Success(c, allUser)
}
