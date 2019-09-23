package model

import "time"

//likesテーブル用構造体
type Like struct {
	ID        uint       `gorm:"primary_key; auto_increment" json:"id"`
	UserID    uint       `gorm:"not null;default:0" json:"user_id"`
	PostID    uint       `gorm:"not null;default:0" json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
