package database

import (
	"instagram/api/model"

	"github.com/jinzhu/gorm"
)

//db接続情報を持つ
type postRepository struct {
	db *gorm.DB
}

//テーブル操作のインターフェース
type PostRepository interface {
	Create(post *model.Post) error
	GetForThisPage(pageNum int, count int, posts *model.Posts) (*model.Posts, error)
	GetCount(posts *[]model.Post, count int) (count, error)
	GetLastID(post *model.Post) (uint, error)
	Delete(postId uint) error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

// 投稿カラム新規作成
func (postRepository *postRepository) Create(post *model.Post) error {
	return postRepository.db.Create(post).Error
}

// 指定ページに対応した投稿を取得
func (postRepository *postRepository) GetForThisPage(pageNum int, count int, posts *model.Posts) (*model.Posts, error) {

	err := postRepository.db.Order("id desc").Limit(count)).Offset(offset).Preload("User").Preload("LikeUsers").Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// postカラムの個数を取得
func (postRepository *postRepository) GetCount(posts *[]model.Post, count int) (count, error) {
	if err := postRepository.db.Find(&posts).Count(&count).Error; err!= nil {
		return nil, err
	}
	return count, nil
}

// postテーブルから最後のidを取得
func (postRepository *postRepository) GetLastID(post *model.Post) (uint, error) {
	if err := postRepository.db.Last(post).Error; err != nil {
		return nil, err
	}
	return post.ID, err
}

func (postRepository *postRepository) GetByID(post *model.Post, postId uint) (*model.Post, error) {

	if err := postRepository.db.Where("id = ?", postId).First(&post).Error; err != nil {
		return nil, err
	}

	if err := postRepository.db.Model(&post).Association("LikeUsers").Find(&post.LikeUsers).Error; err != nil {
		return nil, err
	}

	return post, err
}

// カラム削除　
func (postRepository *postRepository) Delete(postId uint) error {
	
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

	if err := postRepository.db.Unscoped().Where("post_id = ?", post_id).Delete(&[]Like{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := postRepository.db.Unscoped().Where("id = ?", post_id).Delete(&Post{}).Error; err != nil {
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
