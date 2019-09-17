package model

import (
	"log"
)

//likesテーブル用構造体
type Like struct {
	Model
	UserID uint `gorm:"not null;default:0" json:"user_id"`
	PostID uint `gorm:"not null;default:0" json:"post_id"`
}

// いいねデータ作成
func CreateLike(user_id uint, post_id uint) {
	db := NewDb()
	like := Like{}

	like.UserID = user_id
	like.PostID = post_id

	db.Create(&like)
}

// いいね削除
func DeleteLike(user_id uint, post_id uint) bool {
	db := NewDb()
	like := Like{}

	if err := db.Where("user_id = ? AND post_id = ?", user_id, post_id).First(&like).Error; err != nil {
		return false
	}
	if err := db.Unscoped().Delete(&like).Error; err != nil {
		return false
	}
	return true
}

//PostIDがpost_idであるlikesテーブル取得
func GetLikeByPostID(post_id uint) []Like {
	db := NewDb()
	likes := []Like{}

	if err := db.Where("PostID = ?", post_id).Find(&likes).Error; err != nil {
		log.Fatal(err)
	}

	return likes
}

//likesのuser_idから該当userを取得
func GetUserByPostID(l Like) User {
	db := NewDb()
	user := User{}

	if err := db.Where("ID = ?", l.UserID).Find(user).Error; err != nil {
		log.Fatal(err)
	}

	return user
}
