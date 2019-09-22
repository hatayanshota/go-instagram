package repository

import "instagram/api/model"

//テーブル操作のインターフェース
type LikeRepository interface {
	Create(like *model.Like) error
	Delete(like *model.Like, userId uint, postId uint) error
}
