package handler

import (
	"instagram/api/model"
	"instagram/api/public"
	"net/http"
	"os"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

//oauth2の設定
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

func GetStateString() string {
	var oauthStateString = "scucess"
	return oauthStateString
}

//ユーザーをguthub認証ページにリダイレクトする
func GithubAuth(c echo.Context) error {
	oauthConf := GetOauthConfig()
	oauthStateString := GetStateString()
	url := oauthConf.AuthCodeURL(oauthStateString)
	return c.Redirect(http.StatusTemporaryRedirect, url) //ステータスコード307でリダイレクト
}

//githubがcallbackするメソッド
func GithubCallback(c echo.Context) error {
	oauthConf := GetOauthConfig()
	oauthStateString := GetStateString()
	state := c.FormValue("state")

	if state != oauthStateString {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	//認証コードを取得してトークンに変換するよ
	code := c.FormValue("code")
	github_token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	access_token := github_token.AccessToken

	// GitHubAPIからユーザ情報を取得
	github_user_icon, github_user_name, github_user_id := public.GetGithubUser(access_token)

	// トークンからクッキーとセッションに保存するハッシュを作成
	token_hash := public.TokenHash(access_token)

	// ユーザが存在しなければ作成
	if model.UserIsCreated(token_hash, github_user_icon, github_user_name, github_user_id) == false {
		model.CreateUser(token_hash, github_user_icon, github_user_name, github_user_id)
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
func AuthUserLogin(c echo.Context) error {
	public.AllowCORS(c)

	sess, err := session.Get("session", c)
	github_token_session, ok := sess.Values[githubTokenHash]
	if err != nil || !ok || github_token_session.(string) != public.ReadGithuTokenCookie(c) {
		return c.NoContent(http.StatusUnauthorized)
	} else {
		return c.NoContent(http.StatusOK)
	}
}

// ログアウト機能
func UserLogout(c echo.Context) error {
	// セッション削除とクッキー削除が完了したら、サインアップ画面へ遷移
	if ClearSession(c) && ClearCookie(c) {
		c.Redirect(http.StatusTemporaryRedirect, "http://localhost:3000/sign_up")
		return c.NoContent(http.StatusOK)
	} else {
		return c.NoContent(http.StatusInternalServerError)
	}
}
