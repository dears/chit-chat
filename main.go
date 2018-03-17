package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	// デフォルトのマルチプレクサ
	mux := http.NewServeMux()

	// 静的ファイルの返送
	files := http.FileServer(http.Dir("/public")) // publicディレクトリからファイルを配信するハンドラを作成
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	// mux.HandleFunc("/err", err)

	// mux.HandleFunc("/login", login)
	// mux.HandleFunc("/signup", signup)
	// mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// mux.HandleFunc("/thread/new", newThread)
	// mux.HandleFunc("/thread/create", createThread)
	// mux.HandleFunc("/thread/post", postThread)
	// mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		// ログインしてなければクッキーが作られていないので、エラーが返ってくる
		_, err := session(w, r)

		if err != nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}

// data interface{} はどんな型でも渡せる
// fn ...string でfnという名のstringを可変長で受ける
func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
