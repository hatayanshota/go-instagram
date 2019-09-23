package repository

import "instagram/api/domain/model"

// インターフェース
type UserRepository interface {
	Create(user *model.User) error
	GetByGithubId(user *model.User, githubUserId uint) (*model.User, error)
	GetByID(id uint, user *model.User) (*model.User, error)
	GetLoginUser(loginUser *model.User, githubToken string) (*model.User, bool, error)
	UpdateField(user *model.User, oldField string, newField string) error
}
