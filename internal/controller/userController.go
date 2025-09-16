package controller

import (
	"github.com/gin-gonic/gin"
	"go-bk/internal/model"
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

}
