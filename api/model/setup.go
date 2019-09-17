package model

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID        uint       `gorm:"primary_key; auto_increment" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewDb() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp(mysql:3306)/"+os.Getenv("MYSQL_DATABASE")+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)

	return db
}

func Migrate(db *gorm.DB) {
	//マイグレート
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&Like{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Like{}).AddForeignKey("post_id", "posts(id)", "RESTRICT", "RESTRICT")

	//Db.DropTable(&Like{}, &Post{}, &User{})
}
