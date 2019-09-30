package handler

import (
	"instagram/api/infrastructure/utils"
	"instagram/api/interface/controllers"
	"time"

	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

// このクラスの構造体宣言
type authHandler struct {
	userController controllers.UserController
}

// インターフェース宣言
type AuthHandler interface {
	GithubAuth(c echo.Context) error
	GithubCallback(c echo.Context) error
	AuthUserLogin(c echo.Context) error
	UserLogout(c echo.Context) error
	SetGithubTokenCookie(c echo.Context) error
	SetGithubTokenSession(e *echo.Echo) echo.HandlerFunc
}

// このクラスのゲッター
func NewAuthHandler(uc controllers.UserController) AuthHandler {
	return &authHandler{userController: uc}
}

// github認証時の状態確認用文字列
var oauthStateString = "scucess"

// oauth2configのゲッター
func GetOauthConfig() *oauth2.Config {
	var oauthConf = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:8080/login/github/callback",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"read:user"},
		Endpoint:     githuboauth.Endpoint,
	}
	return oauthConf
}

// ユーザーをguthub認証ページにリダイレクト
func (authHandler *authHandler) GithubAuth(c echo.Context) error {
	oauthConf := GetOauthConfig()
	url := oauthConf.AuthCodeURL(oauthStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// githubがcallbackするメソッド
// 正しい遷移先かを確認し、認証コードからアクセストークンを取得してハッシュに変換
// githubAPIからのユーザー情報をDBに保存
func (authHandler *authHandler) GithubCallback(c echo.Context) error {
	oauthConf := GetOauthConfig()
	state := c.FormValue("state")

	if state != oauthStateString {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	//認証コードを取得してトークンに変換
	code := c.FormValue("code")
	githubToken, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	accessToken := githubToken.AccessToken

	// GitHubAPIからユーザ情報を取得
	githubUserIcon, githubUserName, githubUserID := utils.GetGithubUser(accessToken)

	// トークンからクッキーとセッションに保存するハッシュを作成
	tokenHash := utils.TokenToHash(accessToken)

	// ユーザが存在しなければ作成
	exist, err := authHandler.userController.ExistsUser(tokenHash, githubUserIcon, githubUserName, githubUserID)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	} else if exist == false {
		if err := authHandler.userController.CreateUser(tokenHash, githubUserIcon, githubUserName, githubUserID); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}

	// ハッシュをcookieに設定するためにリダイレクト
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/set/cookie/"+tokenHash)

	// ホーム画面へ移動
	return c.String(http.StatusOK, tokenHash)
}

// AuthUserLogin ユーザーのログイン状態を判定
// @ID authUserLogin
// @Summary ユーザーのログイン状態を判定
// @Description sessionとcookieの値を比較し、ログイン状態を判定
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Failure 401 {string} Unauthorized
// @Router /login [get]
func (authHandler *authHandler) AuthUserLogin(c echo.Context) error {
	sess, err := session.Get("session", c)
	githubTokenSession, ok := sess.Values[githubTokenHash]
	if err != nil || !ok || githubTokenSession.(string) != utils.ReadGithuTokenCookie(c) {
		return c.NoContent(http.StatusUnauthorized)
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// ログアウト機能
func (authHandler *authHandler) UserLogout(c echo.Context) error {
	// セッション削除とクッキー削除が完了したら、サインアップ画面へ遷移
	if ClearSession(c) && ClearCookie(c) {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/sign_up")
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusInternalServerError)
	}
}

// クッキーとセッションに保存するハッシュの名前を定数で宣言
const githubTokenHash = "github_token_hash"

// Githubのトークンをハッシュ化したものをクッキーに設定する
func (authHandler *authHandler) SetGithubTokenCookie(c echo.Context) error {
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

// GitHubのトークンをハッシュ化したものをsessionに保存する
func (authHandler *authHandler) SetGithubTokenSession(e *echo.Echo) echo.HandlerFunc {
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

// クッキー削除
func ClearCookie(c echo.Context) bool {
	cookie, err := c.Cookie(githubTokenHash)
	if err != nil {
		return false
	}
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return true
}

// セッション削除
func ClearSession(c echo.Context) bool {
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
