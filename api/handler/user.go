package handler

import (
	"fmt"
	"instagram/api/model"
	"instagram/api/public"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//ユーザー作成(token, name, iconでPOST)
func CreateUser(c echo.Context) error {
	var user *model.User
	err := c.Bind(user)
	if err != nil {
		return err
	}
	model.CreateUser(user.GithubToken, user.Icon, user.Name, user.GithubId)

	return c.NoContent(http.StatusOK)
}

//ユーザー取得(paramにid)
func GetUser(c echo.Context) error {
	public.AllowCORS(c)
	var user *model.User

	iid, _ := strconv.Atoi(c.Param("id"))
	id := uint(iid)
	user = model.GetUserByID(id)

	return c.JSON(http.StatusOK, user)
}

func GetLoginUser(c echo.Context) error {
	public.AllowCORS(c)
	var user *model.User
	user, _ = model.LoginUser(c)
	fmt.Print(user)
	return c.JSON(http.StatusOK, user)
}
