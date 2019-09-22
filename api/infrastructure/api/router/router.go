package router

import (
	"instagram/api/infrastructure/api/handler"

	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, handler handler.AppHandler) {

	// swaggerを使う
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/auth", handler.AuthUserLogin)

	e.GET("/logout", handler.UserLogout)

	// サインアップ
	e.GET("/signup", handler.GithubAuth)
	e.GET("/login/github/callback", handler.GithubCallback)
	e.GET("/set/cookie/:token_hash", handler.SetGithubTokenCookie)
	e.GET("/set/session/:token_hash", handler.SetGithubTokenSession(e))

	// userデータベース操作
	e.POST("/user", handler.CreateUser)
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users/mydata", handler.GetLoginUser)

	// postデータベース操作
	e.POST("/posts/new", handler.CreatePost)
	e.GET("/posts", handler.GetPostIndex)
	e.POST("/posts/:post_id/delete", handler.DeletePost)

	// likeデータベース操作
	e.POST("/posts/like", handler.CreateLike)
	e.POST("/posts/delete/like", handler.DeleteLike)
	e.GET("/posts/:post_id/likes", handler.GetLike)
}
