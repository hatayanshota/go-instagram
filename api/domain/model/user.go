package model

import "time"

//usersテーブル用構造体
type User struct {
	ID          uint       `gorm:"primary_key; auto_increment" json:"id"`
	GithubToken string     `gorm:"type:varchar(255);unique;not null;default:''" json:"token"`
	Icon        string     `gorm:"type:varchar(255);not null;default:''" json:"icon"`
	Name        string     `gorm:"type:varchar(255);not null;default:''" json:"name"`
	GithubId    uint       `gorm:"type:int;not null;default:0" json:"github_id"`
	Posts       []Post     `json:"posts"`
	LikePosts   []Post     `gorm:"many2many:likes" json:"like_posts"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type Users []*User
