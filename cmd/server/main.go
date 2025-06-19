package main

import (
	"log"
	"net/http"

	openapi "github.com/dill3102/game_server_guild/api"
	handler "github.com/dill3102/game_server_guild/internal/controller"
	"github.com/dill3102/game_server_guild/internal/db"
	"github.com/dill3102/game_server_guild/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// インフラ層（永続化）などの初期化
	repo := repository.NewRepository(db.NewDB())
	handler := handler.NewHandler(repo)

	// OpenAPIのハンドラに自作ハンドラを接続
	// api := openapi.NewStrictHandler(handler, nil)

	// Echoルーターにマウント
	openapi.RegisterHandlers(e, handler)

	// サーバー起動
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("shutting down the server")
	}
}
