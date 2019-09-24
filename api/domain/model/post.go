package model

import "time"

// postテーブル用構造体
type Post struct {
	ID        uint       `gorm:"primary_key; auto_increment" json:"id"`
	UserID    uint       `gorm:"not null;default:0" json:"user_id"`
	ImageURL  string     `gorm:"type:varchar(255);not null;default:''" json:"image"`
	Caption   string     `gorm:"type:varchar(255);not null;default:''" json:"caption"`
	LikeUsers []User     `json:"like_users"`
	Likes     []Like     `json:"likes"`
	User      User       `json:"user"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
