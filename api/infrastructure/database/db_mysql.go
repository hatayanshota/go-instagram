package database

import (
	"os"

	"github.com/jinzhu/gorm"
)

// DB接続
func NewMysqlDB() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp(mysql:3306)/"+os.Getenv("MYSQL_DATABASE")+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)

	//マイグレート
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&Like{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Like{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")

	return db
}
