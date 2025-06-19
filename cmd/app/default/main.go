package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8081", nil)
}

/*
go run cmd/app/main.go
localhost:8081でアクセスする
*/
