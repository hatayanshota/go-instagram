package database

import (
	"instagram/api/domain/model"

	"github.com/jinzhu/gorm"
)

//db接続情報を持つ
type postRepository struct {
	db *gorm.DB
}

//テーブル操作のインターフェース
type PostRepository interface {
	Create(post *model.Post) error
	GetForThisPage(count int, offset int, posts *[]model.Post) (*[]model.Post, error)
	GetCount(posts *[]model.Post, count int) (int, error)
	GetLastID(post *model.Post) (uint, error)
	GetByID(post *model.Post, postID uint) (*model.Post, error)
	Delete(postID uint) error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

// 投稿カラム新規作成
func (postRepository *postRepository) Create(post *model.Post) error {
	return postRepository.db.Create(post).Error
}

// 指定ページに対応した投稿を取得
func (postRepository *postRepository) GetForThisPage(count int, offset int, posts *[]model.Post) (*[]model.Post, error) {

	err := postRepository.db.Order("id desc").Limit(count).Offset(offset).Preload("User").Preload("LikeUsers").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// postカラムの個数を取得
func (postRepository *postRepository) GetCount(posts *[]model.Post, count int) (int, error) {
	if err := postRepository.db.Find(&posts).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// postテーブルから最後のidを取得
func (postRepository *postRepository) GetLastID(post *model.Post) (uint, error) {
	if err := postRepository.db.Last(post).Error; err != nil {
		return 0, err
	}
	return post.ID, nil
}

func (postRepository *postRepository) GetByID(post *model.Post, postID uint) (*model.Post, error) {

	if err := postRepository.db.Where("id = ?", postID).First(&post).Error; err != nil {
		return nil, err
	}

	if err := postRepository.db.Model(&post).Association("LikeUsers").Find(&post.LikeUsers).Error; err != nil {
		return nil, err
	}

	return post, nil
}

// カラム削除
func (postRepository *postRepository) Delete(postID uint) error {

	// トランザクション開始
	tx := postRepository.db.Begin()

	// トランザクション中にエラーが発生した場合はロールバック
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// トランザクション開始に失敗した場合はfalseを返す
	if err := tx.Error; err != nil {
		return err
	}

	if err := postRepository.db.Unscoped().Where("post_id = ?", postID).Delete(&[]model.Like{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := postRepository.db.Unscoped().Where("id = ?", postID).Delete(&model.Post{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// トランザクションをコミットする
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
