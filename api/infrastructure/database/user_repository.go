package database

import (
	"instagram/api/domain/model"
	"log"

	"github.com/jinzhu/gorm"
)

// db接続情報を持つリポジトリ
type userRepository struct {
	db *gorm.DB
}

// インターフェース
type UserRepository interface {
	Create(user *model.User) error
	GetByGithubId(user *model.User, githubUserId uint) (*model.User, error)
	GetByID(id uint, user *model.User) (*model.User, error)
	GetLoginUser(loginUser *model.User, githubToken string) (*model.User, bool, error)
	UpdateField(user *model.User, oldField string, newField string) error
}

// コンストラクタ
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// ユーザーカラムの新規作成
func (userRepository *userRepository) Create(user *model.User) error {
	return userRepository.db.Create(user).Error
}

// ユーザーカラムの取得(githubIdから)
func (userRepository *userRepository) GetByGithubId(user *model.User, githubUserId uint) (*model.User, error) {
	if err := userRepository.db.Where("github_id = ?", githubUserId).First(user).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return user, nil
}

// ユーザーカラムの取得(idから)
func (userRepository *userRepository) GetByID(id uint, user *model.User) (*model.User, error) {

	// 投稿といいねした投稿とともに取得
	if err := userRepository.db.Where("id = ?", id).Preload("Posts").Preload("LikePosts").Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// ユーザーカラムの取得(現在ログインしているユーザー)
func (userRepository *userRepository) GetLoginUser(loginUser *model.User, githubToken string) (*model.User, bool, error) {
	if err := userRepository.db.Where("github_token = ?", githubToken).First(loginUser).Error; err != nil {
		return loginUser, false, err
	}
	return loginUser, true, nil
}

// 指定したフィールドの更新
func (userRepository *userRepository) UpdateField(user *model.User, oldField string, newField string) error {
	if err := userRepository.db.Model(&user).Update(oldField, newField).Error; err != nil {
		return err
	}
	return nil
}
