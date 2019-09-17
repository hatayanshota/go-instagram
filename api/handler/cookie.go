package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// クッキーとセッションに保存するハッシュの名前を定数で宣言
const githubTokenHash = "github_token_hash"

// Githubのトークンをハッシュ化したものをクッキーに設定する
func SetGithubTokenCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = githubTokenHash
	cookie.Value = c.Param("token_hash")
	cookie.Expires = time.Now().Add(24 * time.Hour) // １日でセッションがなくなる
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	// クッキーを設定できたらセッションにハッシュを保存するように遷移
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/set/session/"+c.Param("token_hash"))

	return c.String(http.StatusOK, cookie.Value)
}

// クッキーの有効期限を切れるように設定する
func ClearCookie(c echo.Context) bool {
	cookie, err := c.Cookie(githubTokenHash)
	if err != nil {
		return false
	}
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return true
}
