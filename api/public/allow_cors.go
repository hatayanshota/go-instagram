package public

import "github.com/labstack/echo/v4"

func AllowCORS(c echo.Context) {
	// クライアントサーバからの受け入れ許可
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:3000")
	// クライアントサーバからのクッキー受け入れに必要
	c.Response().Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
}
