package main

import (
	"github.com/gin-gonic/gin"
	"go-bk/configs"
	"go-bk/internal/routes"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	//读取数据库配置
	db = configs.InitDB()
	//gin的默认值
	r := gin.Default()
	//注册路由
	routes.RegisterRouter(r)
	//启动服务
	r.Run(configs.Config.ServerPort)
}
