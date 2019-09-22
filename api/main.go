//go:generate swag init

package main

import (
	"instagram/api/infrastructure/api/middleware"
	"instagram/api/infrastructure/api/router"
	"instagram/api/infrastructure/database"
	"instagram/api/infrastructure/env"
	"instagram/api/infrastructure/storage"

	_ "instagram/api/docs"

	"github.com/labstack/echo"
)

// @title インスタグラムもどき課題 API サーバー
// @version 1.0.0
// @tag.name go-instagram
// @description
// @termsOfService localhost:8080

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name none

// @host localhost:8080
// @BasePath /
func main() {

	// godotenvの初期化
	env.SetEnv()

	// データベース設定
	db := database.NewMysqlDB()

	// aws sdk設定
	s3Config := storage.NewS3()

	// interacterの設定
	r := registry.Newinteracter(db, s3Config)

	// 依存解決
	h := r.NewAppHandler()

	// Echoのインスタンス作る
	e := echo.New()

	//router
	router.NewRouter(e, h)

	//ミドルウェア
	middleware.Middleware(e)

	// サーバー起動
	e.Start(":80")
}
