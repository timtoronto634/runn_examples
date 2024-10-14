package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// エンドポイント1: 単純な挨拶
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// エンドポイント2: パスパラメータを使用
	mux.HandleFunc("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Fprintf(w, "Greetings, %s!", name)
	})

	// エンドポイント3: クエリパラメータを使用
	mux.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		if message == "" {
			message = "No message provided"
		}
		fmt.Fprintf(w, "Echo: %s", message)
	})

	// GETメソッドの処理
	mux.HandleFunc("GET /resource", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "GET request received")
	})

	// POSTメソッドの処理
	mux.HandleFunc("POST /resource", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "POST request received")
	})

	// PUTメソッドの処理
	mux.HandleFunc("PUT /resource", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "PUT request received")
	})

	// DELETEメソッドの処理
	mux.HandleFunc("DELETE /resource", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "DELETE request received")
	})

	// 全てのメソッドを処理する汎用ハンドラ
	mux.HandleFunc("/all-methods", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s request received", r.Method)
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
