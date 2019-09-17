package route

import (
	"instagram/api/handler"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//echoインスタンス初期化
func NewEcho() *echo.Echo {
	e := echo.New()
	return e
}

// ルーティング
func Route(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/auth", handler.AuthUserLogin)

	e.GET("/set/cookie/:token_hash", handler.SetGithubTokenCookie)

	e.GET("/set/session/:token_hash", handler.SetGithubTokenSession(e))

	e.GET("/logout", handler.UserLogout)

	e.GET("/signup", handler.GithubAuth)

	e.GET("/login/github/callback", handler.GithubCallback)

	//userデータベース操作
	e.POST("/user", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users/mydata", handler.GetLoginUser)

	//postデータベース操作
	e.POST("/posts/new", handler.CreatePost)
	e.GET("/posts", handler.GetPostIndex)
	e.POST("/posts/:post_id/delete", handler.DeletePost)

	//likeデータベース操作
	e.POST("/posts/like", handler.CreateLike)
	e.POST("/posts/delete/like", handler.DeleteLike)
	e.GET("/posts/:post_id/likes", handler.GetLike)

}
