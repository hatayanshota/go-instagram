package handler

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// type sessionHandler struct {
// }

// type SessionHandler interface {
// 	SetGithubTokenSession(e *echo.Echo) echo.HandlerFunc
// 	ClearSession(c echo.Context) bool
// }

// func NewSessionHandler() {
// 	return
// }

// GitHubのトークンをハッシュ化したものをsessionに保存する
func (sessionHandler *sessionHandler) SetGithubTokenSession(e *echo.Echo) echo.HandlerFunc {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7, // 1週間でセッションが消される
			HttpOnly: true,
		}

		// セッションに値を代入
		sess.Values[githubTokenHash] = c.Param("token_hash")

		// セッションを保存できたか確認
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		// クライアントのホーム画面へ移動
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/posts")

		return c.NoContent(http.StatusOK)
	}
}

// セッションの有効期限を切れるように設定する
func (sessionHandler *sessionHandler) ClearSession(c echo.Context) bool {
	sess, err := session.Get("session", c)
	if err != nil {
		return false
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return false
	}

	return true
}
