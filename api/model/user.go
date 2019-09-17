package model

import (
	"instagram/api/public"
	"log"

	"github.com/labstack/echo/v4"
)

//usersテーブル用構造体
type User struct {
	Model
	GithubToken string `gorm:"type:varchar(255);unique;not null;default:''" json:"token"`
	Icon        string `gorm:"type:varchar(255);not null;default:''" json:"icon"`
	Name        string `gorm:"type:varchar(255);not null;default:''" json:"name"`
	GithubId    uint   `gorm:"type:int;not null;default:0" json:"github_id"`
	Posts       []Post `json:"posts"`
	LikePosts   []Post `gorm:"many2many:likes" json:"like_posts"`
}

//ユーザーの作成(データ投入時に起こったエラーを返す)
func CreateUser(token string, icon string, name string, github_id uint) {
	db := NewDb()
	user := &User{}

	user.GithubToken = token
	user.Icon = icon
	user.Name = name
	user.GithubId = github_id

	db.Create(user)
}

//ユーザー取得(idから)
func GetUserByID(id uint) *User {
	db := NewDb()
	user := &User{}

	if err := db.Where("id = ?", id).Find(user).Error; err != nil {
		log.Fatal(err)
	}

	// ユーザの投稿を取得
	db.Model(&user).Association("Posts").Find(&user.Posts)

	for index, user_post := range user.Posts {
		// 投稿に対するいいね情報を取得
		db.Model(&user_post).Association("LikeUsers").Find(&user.Posts[index].LikeUsers)
	}

	return user
}

// GithubのIDでユーザの一意性を確保しつつ検索をかける
func UserIsCreated(github_token, github_user_icon, github_user_name string, github_id uint) bool {
	db := NewDb()
	user := User{}

	if err := db.Where("github_id = ?", github_id).First(&user).Error; err != nil {
		return false
	} else {
		// ハッシュが更新されている場合はデータベースを更新
		if user.GithubToken != github_token {
			db.Model(&user).Update("github_token", github_token)
		}
		// アイコンが変更されている場合はデータベースを更新
		if user.Icon != github_user_icon {
			db.Model(&user).Update("icon", github_user_icon)
		}
		// 名前が変更されている場合はデータベースを更新
		if user.Name != github_user_name {
			db.Model(&user).Update("name", github_user_name)
		}
		return true
	}
}

func LoginUser(c echo.Context) (*User, bool) {
	db := NewDb()
	login_user := User{}

	if err := db.Where("github_token = ?", public.ReadGithuTokenCookie(c)).First(&login_user).Error; err != nil {
		return &login_user, false
	}

	return &login_user, true
}
