package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Config struct {
	ServerPort string `yaml:"server.port"`
	DBUser     string `json:"database.user"`
	DBPassword string `json:"database.password"`
	DBName     string `json:"database.name"`
	DBHost     string `json:"database.host"`
	DBPort     string `json:"database.port"`
}

func InitDB() *gorm.DB {
	dsn := Config.DBUser + ":" + Config.DBPassword + "@tcp(" + Config.DBHost + ":" + Config.DBPort + ")/" + Config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("链接数据库出错了:", err.Error())
	}

	return db

}
