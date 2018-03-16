package main

import (
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
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}

	templates := template.Must(template.ParseFiles(files...))
	// todo
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
