package public

import (
	"time"

	"github.com/labstack/echo/v4"
)

// 認証の時にクッキーに設定したGitHubのトークンをハッシュ化したものを読み取るための関数
func ReadGithuTokenCookie(c echo.Context) string {
	cookie, err := c.Cookie("github_token_hash")
	if err != nil || time.Now().Sub(cookie.Expires) < 0 {
		return "cannot read github token cookie"
	}

	return cookie.Value
}
