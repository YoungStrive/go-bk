package routes

import "github.com/gin-gonic/gin"

// 注册所有的路由
func RegisterRouter(r *gin.Engine) {
	userRouters := r.Group("/user")
	userRouters.POST("/")
	userRouters.GET("/")

}
