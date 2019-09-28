//go:generate swag init

package main

import (
	"instagram/api/infrastructure/api/middleware"
	"instagram/api/infrastructure/api/router"
	"instagram/api/infrastructure/database"
	"instagram/api/infrastructure/env"
	"instagram/api/infrastructure/storage"
	"instagram/api/registry"
	"log"

	_ "instagram/api/docs"

	"github.com/labstack/echo/v4"
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

	// データベース接続
	db := database.NewMysqlDB()

	// データベース切断
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// aws sdk設定
	s3Config := storage.NewS3()

	// コンストラクタインジェクションによるDI
	r := registry.NewInteractor(db, s3Config)

	// ハンドラ設定
	h := r.NewAppHandler()

	// Echoのインスタンス作成
	e := echo.New()

	// routerの設定
	router.NewRouter(e, h)

	// ミドルウェア
	middleware.Middleware(e)

	// サーバー起動
	e.Start(":80")
}
