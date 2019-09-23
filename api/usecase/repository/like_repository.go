package repository

import "instagram/api/domain/model"

//テーブル操作のインターフェース
type LikeRepository interface {
	Create(like *model.Like) error
	Delete(like *model.Like, userID uint, postID uint) error
}
