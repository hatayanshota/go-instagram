package service

import (
	"instagram/api/model"
	"instagram/api/usecase/repository"
)

// serviceはrepositoryのインターフェースとpresenterのインターフェースに依存
type userService struct {
	UserRepository repository.UserRepository
}

// インターフェース
type UserService interface {
	Create(user *model.User) error
	GetByID(id uint, user *model.User) *model.User, error
	GetLoginUser(u *model.User, githubToken string) (*model.User, bool, error)
	Exists(user *model.User, githubToken, githubUserIcon, githubUserName string, githubUserId uint) (bool, error)
}

// コンストラクタ
func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{repo, pre}
}

// ユーザー作成
func (userService *userService) Create(user *model.User) error {
	return userService.UserRepository.Create(user)
}

// idからユーザー取得
func (userService *userService) GetByID(id uint, user *model.User) (*model.User, error) {
	return userService.UserRepository.GetByID(id, user)
}

//ログインユーザーの取得
func (userService *userService) GetLoginUser(u *model.User, githubToken string) (*model.User, bool, error) {
	loginUser, isLogin, err := userService.UserRepository.GetLoginUser(u, githuToken)

	return loginUser, isLogin, err
}

// GithubのIDでユーザの一意性を確保しつつ検索をかける
func (userService *userService) Exists(user *model.User, githubToken, githubUserIcon, githubUserName string, githubUserId uint) (bool, error) {

	if err := userService.UserRepository.GetByGithubId(user, githubUserId); err != nil {
		return false, err
	} else {
		// ハッシュが更新されている場合はデータベースを更新
		if user.GithubToken != githubToken {
			if err = userService.UserRepository.UpdateField(user, "github_token", githubToken); err != nil {
				return false, err 
			}
		}
		// アイコンが変更されている場合はデータベースを更新
		if user.Icon != githubUserIcon {
			if err = userService.UserRepository.UpdateField(user, "icon", githubUserIcon); err != nil {
				return false, err 
			}
		}
		// 名前が変更されている場合はデータベースを更新
		if user.Name != githubUserName {
			if err = userService.UserRepository.UpdateField(user, "name", githubUserName); err != nil {
				return false, err 
			}
		}
		return true, nil
	}
}