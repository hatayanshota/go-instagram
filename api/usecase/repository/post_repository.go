package repository

import "instagram/api/domain/model"

//テーブル操作のインターフェース
type PostRepository interface {
	Create(post *model.Post) error
	GetForThisPage(count int, offset int, posts *[]model.Post) (*[]model.Post, error)
	GetCount(posts *[]model.Post, count int) (int, error)
	GetLastID(post *model.Post) (uint, error)
	GetByID(post *model.Post, postID uint) (*model.Post, error)
	Delete(postID uint) error
}
