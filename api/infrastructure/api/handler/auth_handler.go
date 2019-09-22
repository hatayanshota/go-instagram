package handler

import (
	"instagram/api/infrastructure/utils"
	"instagram/api/model"
	"instagram/api/public"
	"net/http"
	"os"

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
	github_token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	access_token := github_token.AccessToken

	// GitHubAPIからユーザ情報を取得
	github_user_icon, github_user_name, github_user_id := utils.GetGithubUser(access_token)

	// トークンからクッキーとセッションに保存するハッシュを作成
	token_hash := utils.TokenToHash(access_token)

	// ユーザが存在しなければ作成
	exist, err := authHandler.userController.ExistsUser(token_hash, github_user_icon, github_user_name, github_user_id)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	else if exist == false {
		if err := authHandler.userController.CreateUser(token_hash, github_user_icon, github_user_name, github_user_id); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
	}

	// ハッシュをcookieに設定するためにリダイレクト
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/set/cookie/"+token_hash)

	return c.String(http.StatusOK, token_hash)
}

// AuthUserLogin ユーザーのログイン状態を判定
// @ID authUserLogin
// @Summary ユーザーのログイン状態を判定
// @Description sessionとcookieの値を比較し、ログイン状態を判定
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Failure 401 {string} Unauthorized
// @Router /auth [get]
func (authHandler *authHandler) AuthUserLogin(c echo.Context) error {

	sess, err := session.Get("session", c)
	github_token_session, ok := sess.Values[githubTokenHash]
	if err != nil || !ok || github_token_session.(string) != utils.ReadGithuTokenCookie(c) {
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
