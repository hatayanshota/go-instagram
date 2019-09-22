package repository

import "instagram/api/model"

// frastructure/repositoryのインターフェース
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uint, user *model.User) (*model.User, error)
	GetLoginUser(loginUser *model.User, githubToken string) (*model.User, bool, error)
}
