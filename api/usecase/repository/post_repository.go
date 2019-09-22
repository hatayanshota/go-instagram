package repository

import "instagram/api/model"

//テーブル操作のインターフェース
type PostRepository interface {
	Create(post *model.Post) error
	GetForThisPage(pageNum int, count int, posts *model.Posts) (*model.Posts, error)
	GetCount(posts *[]model.Post, count int) (count, error)
	GetLastID(post *model.Post) (uint, error)
	Delete(postId uint) error
}
