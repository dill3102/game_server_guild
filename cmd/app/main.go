package main

import (
	"log"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/go-chi/chi/v5"

	// 生成されたパッケージ名に合わせて import
	"api"
)

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("failed to get swagger: %v", err)
	}

	r := chi.NewRouter()

	// OpenAPIのバリデーションミドルウェア
	r.Use(middleware.OapiRequestValidator(swagger))

	// ハンドラの実装を渡す
	handler := &APIHandler{}
	api.RegisterHandlers(r, handler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
