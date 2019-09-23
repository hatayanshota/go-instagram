package controllers

import (
	"instagram/api/domain/model"
	"instagram/api/usecase/service"
)

// controlleはserviceに依存
type userController struct {
	userService service.UserService
}

// インターフェース宣言
type UserController interface {
	CreateUserByModel(user *model.User) error
	CreateUser(tokenHash string, githubUserIcon string, githubUserName string, githubUserId uint) error
	GetUserByID(id uint) (*model.User, error)
	GetLoginUser(githubToken string) (*model.User, bool, error)
	ExistsUser(githubToken, githubUserIcon, githubUserName string, githubUserId uint) (bool, error)
}

// serviceに依存したcontrollerを生成
func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

// ユーザーの作成(modelから)
func (userController *userController) CreateUserByModel(user *model.User) error {
	return userController.userService.Create(user)
}

// ユーザーの作成
func (userController *userController) CreateUser(tokenHash string, githubUserIcon string, githubUserName string, githubUserId uint) error {

	user := &model.User{
		GithubToken: tokenHash,
		Icon:        githubUserIcon,
		Name:        githubUserName,
		GithubID:    githubUserId,
	}

	return userController.userService.Create(user)
}

// ユーザーの取得
func (userController *userController) GetUserByID(id uint) (*model.User, error) {

	u := &model.User{}

	user, err := userController.userService.GetByID(id, u)
	if err != nil {
		return nil, err
	}
	return user, err
}

// ログインユーザーの取得
func (userController *userController) GetLoginUser(githubToken string) (*model.User, bool, error) {

	u := &model.User{}

	loginUser, isLogin, err := userController.userService.GetLoginUser(u, githubToken)

	return loginUser, isLogin, err
}

// ユーザーが存在しているか
func (userController *userController) ExistsUser(githubToken, githubUserIcon, githubUserName string, githubUserId uint) (bool, error) {
	user := &model.User{}
	return userController.userService.Exists(user, githubToken, githubUserIcon, githubUserName, githubUserId)
}
