package routes

import (
	"github.com/gin-gonic/gin"
	"go-bk/internal/controller"
	"go-bk/middlewares"
)

// 注册所有的路由
func RegisterRouter(r *gin.Engine) {
	userRouters := r.Group("/user")
	userRouters.POST("/registerUser", controller.RegisterUser)
	userRouters.POST("/login", controller.LoginUser)
	userRouters.GET("/listUser", middlewares.AuthMiddleware(), controller.ListUser)

	postRouters := r.Group("/post")
	postRouters.POST("/add", middlewares.AuthMiddleware(), controller.AddPost)
	postRouters.PUT("/update", middlewares.AuthMiddleware(), controller.UpdatePost)
	postRouters.GET("/list", middlewares.AuthMiddleware(), controller.ListPost)
	postRouters.GET("/getPost", middlewares.AuthMiddleware(), controller.GetPost)
	postRouters.DELETE("/del", middlewares.AuthMiddleware(), controller.DeletePost)

}
