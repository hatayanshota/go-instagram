package database

import (
	"instagram/api/domain/model"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB接続
func NewMysqlDB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp(mysql:3306)/"+os.Getenv("MYSQL_DATABASE")+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	//マイグレート
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&model.Like{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&model.Like{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")

	return db
}
