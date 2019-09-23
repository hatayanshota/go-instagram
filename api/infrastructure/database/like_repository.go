package database

import (
	"instagram/api/domain/model"

	"github.com/jinzhu/gorm"
)

//db接続情報を持つ
type likeRepository struct {
	db *gorm.DB
}

//テーブル操作のインターフェース
type LikeRepository interface {
	Create(like *model.Like) error
	Delete(like *model.Like, userId uint, postID uint) error
}

// コンストラクタ
func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{db}
}

// カラム作成
func (likeRepository *likeRepository) Create(like *model.Like) error {
	return likeRepository.db.Create(like).Error
}

// カラム削除
func (likeRepository *likeRepository) Delete(like *model.Like, userId uint, postID uint) error {

	if err := likeRepository.db.Where("user_id = ? AND post_id = ?", userId, postID).First(like).Error; err != nil {
		return err
	}
	if err := likeRepository.db.Unscoped().Delete(like).Error; err != nil {
		return err
	}
	return nil
}
