package model

import (
	"log"
)

//postsテーブル用構造体
type Post struct {
	Model
	UserID    uint   `gorm:"not null;default:0" json:"user_id"`
	ImageURL  string `gorm:"type:varchar(255);not null;default:''" json:"image"`
	Caption   string `gorm:"type:varchar(255);not null;default:''" json:"caption"`
	LikeUsers []User `gorm:"many2many:likes" json:"like_users"`
	Likes     []Like `json:"likes"`
	User      User   `json:"user"`
}

// post情報とユーザ情報を紐付ける
type PostRelatedUser struct {
	Post
	UserName string `json:"user_name"`
	IconURL  string `json:"icon_url"`
}

//投稿作成
func CreatePost(user_id uint, image_url string, caption string) {
	db := NewDb()
	post := &Post{}

	post.UserID = user_id
	post.ImageURL = image_url
	post.Caption = caption

	db.Create(post)

}

//投稿データ10件を返す
func GetPost(page_num int) *[]Post {
	db := NewDb()

	offset := 0

	if page_num != 0 {
		offset = (page_num - 1) * 10
	}

	var posts []Post
	db.Order("id desc").Limit(10).Offset(offset).Preload("User").Preload("LikeUsers").Find(&posts)

	return &posts
}

// 全投稿数を取得
func GetPostCount() int {
	db := NewDb()
	var posts []Post
	var count int
	db.Find(&posts).Count(&count)

	return count
}

//id指定して投稿データ返す
func GetPostByID(id uint) *Post {
	db := NewDb()
	post := Post{}

	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		log.Fatal(err)
	}

	db.Model(&post).Association("LikeUsers").Find(&post.LikeUsers)

	return &post
}

//一番最後のpost_idを返す
func GetLastPostID() uint {
	db := NewDb()
	post := Post{}

	db.Last(&post)

	return post.ID
}

//id指定して投稿削除
func DeletePost(post_id uint) bool {
	db := NewDb()

	// トランザクション開始
	tx := db.Begin()

	// トランザクション中にエラーが発生した場合はロールバック
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// トランザクション開始に失敗した場合はfalseを返す
	if err := tx.Error; err != nil {
		return false
	}

	if err := db.Unscoped().Where("post_id = ?", post_id).Delete(&[]Like{}).Error; err != nil {
		tx.Rollback()
		return false
	}

	if err := db.Unscoped().Where("id = ?", post_id).Delete(&Post{}).Error; err != nil {
		tx.Rollback()
		return false
	}

	// トランザクションをコミットする
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return false
	}

	return true
}
