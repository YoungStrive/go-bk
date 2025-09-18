package routes

import (
	"github.com/gin-gonic/gin"
	"go-bk/internal/controller"
)

// 注册所有的路由
func RegisterRouter(r *gin.Engine) {
	userRouters := r.Group("/user")
	userRouters.POST("/registerUser", controller.RegisterUser)
	userRouters.GET("/listUser", controller.ListUser)

}
