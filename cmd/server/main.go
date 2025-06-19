package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// 生成されたOpenAPIハンドラパッケージ

	openapi "example.com/dill3102/game_server/api"         // 実際のパスに応じて変更
	"example.com/dill3102/game_server/internal/handler"    // DDDのハンドラ
	"example.com/dill3102/game_server/internal/repository" // DDDのリポジトリ
	"example.com/dill3102/game_server/internal/usecase"    // DDDのユースケース
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// インフラ層（永続化）などの初期化
	userRepo := repository.NewUserRepository()

	// ユースケースの初期化
	userUsecase := usecase.NewUserUsecase(userRepo)

	// ハンドラ層の初期化（依存注入）
	userHandler := handler.NewUserHandler(userUsecase)

	// OpenAPIのハンドラに自作ハンドラを接続
	api := openapi.NewStrictHandler(userHandler, nil)

	// Echoルーターにマウント
	openapi.RegisterHandlers(e, api)

	// サーバー起動
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("shutting down the server")
	}
}
