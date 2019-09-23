package handler

import (
	"instagram/api/domain/model"
	"instagram/api/infrastructure/utils"
	"instagram/api/interface/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// handlerはcontrollerに依存
type userHandler struct {
	userController controllers.UserController
}

// インターフェース
type UserHandler interface {
	CreateUser(c echo.Context) error
	GetUserByID(c echo.Context) error
	GetLoginUser(c echo.Context) error
}

// コンストラクタ
func NewUserHandler(uc controllers.UserController) UserHandler {
	return &userHandler{userController: uc}
}

//ユーザーの作成
func (userHandler *userHandler) CreateUser(c echo.Context) error {

	// リクエスト用のentityを作成
	user := &model.User{}

	// bodyからデータ取得
	if err := c.Bind(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	// handlerにデータを渡す
	if err := userHandler.userController.CreateUserByModel(user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

// idからユーザーの取得
func (userHandler *userHandler) GetUserByID(c echo.Context) error {

	// idパラメータの取得
	id_int, _ := strconv.Atoi(c.Param("id"))
	id := uint(id_int)

	// コントローラに送信
	user, err := userHandler.userController.GetUserByID(id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, user)
}

//ログインユーザーの取得
func (userHandler *userHandler) GetLoginUser(c echo.Context) error {

	githubToken := utils.ReadGithuTokenCookie(c)

	// コントローラに送信
	user, _, err := userHandler.userController.GetLoginUser(githubToken)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, user)
}
