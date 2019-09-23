package middleware

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
func Middleware(e *echo.Echo) {

	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
}
