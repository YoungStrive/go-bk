package main

import (
	"github.com/gin-gonic/gin"
	"go-bk/configs"
	"go-bk/internal/routes"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func main() {

	// 读取 YAML 文件
	data, err := os.ReadFile("configs/api.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// 解析 YAML 数据到结构体
	var config configs.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//读取数据库配置
	db = configs.InitDB(&config)
	//gin的默认值
	r := gin.Default()
	//注册路由
	routes.RegisterRouter(r)
	//启动服务
	r.Run(config.Server.Port)
}
