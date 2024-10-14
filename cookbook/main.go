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

	// サーバーの起動
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
