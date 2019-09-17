//go:generate swag init

package main

import (
	"instagram/api/model"
	"instagram/api/public"
	"instagram/api/route"

	_ "instagram/api/docs"
)

func init() {
	//godotenvの初期化
	public.SetEnv()
}

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
	// Echoのインスタンス作る
	e := route.NewEcho()

	//データベース接続
	db := model.NewDb()
	model.Migrate(db)

	//ルーティング
	route.Route(e)

	//ミドルウェア
	public.Middleware(e)

	// サーバー起動
	e.Start(":80") //ポート番号指定してね
}
