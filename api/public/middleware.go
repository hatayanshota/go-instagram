package public

import (
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
func Middleware(e *echo.Echo) {

	e.Use(middleware.Logger())

	e.Use(middleware.Recover())
}
