package controllers

import (
	"instagram/api/usecase/usecase/service"

	"github.com/labstack/echo"
)

// controlleはserviceに依存
type userController struct {
	userService service.UserService
}

// インターフェース宣言
type UserController interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) *model.User error
	GetLoginUser(c echo.Context) *model.User bool error
}

// serviceに依存したcontrollerを生成
func NewUserController(us service.UserService) UserController {
	return &userController{us}
}

// ユーザーの作成(modelから)
func (userController *userController) CreateUser(user *model.User) error {
	return userController.userService.Create(user)
}

// ユーザーの作成
func (userController *userController) CreateUser(token_hash, github_user_icon, github_user_name, github_user_id) error {
	
	user := &model.User{
		GithubToken: token_hash, 
		Icon: github_user_icon,
		Name: github_user_name,
		GithubID: github_user_id
	}

	return userController.userService.Create(user)
}

// ユーザーの取得
func (userController *userController) GetUserByID(id uint) (*model.User error) {

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