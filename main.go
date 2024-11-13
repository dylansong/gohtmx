package main

import (
	"fmt"
	"net/http"

	"github.com/dylansong/gohtmx/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Hello, World!")
	r := chi.NewRouter()
	r.Get("/status", func ( w http.ResponseWriter, r *http.Request){
		w.Write([]byte("ok"))
	})
	// r.Use(cors.Handler(cors.Options{
	// 	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	// 	AllowedOrigins:   []string{"https://*", "http://*"},
	// 	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: false,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	//   }))

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/status", func ( w http.ResponseWriter, r *http.Request){
		w.Write([]byte("ok v2"))
	})
	r.Get("/", func ( w http.ResponseWriter, r *http.Request){
		// 修复错误：Index 未定义
		todos := []*views.Todo{
			{Id: "1", Description: "Buy milk223"},
			{Id: "2", Description: "Buy eggs"},
			{Id: "3", Description: "Buy bread"},
			{Id: "4", Description: "Buy eggs"},
			{Id: "1", Description: "Buy milk22"},
			{Id: "2", Description: "Buy eggs"},
		

		}
		views.Index(todos).Render(r.Context(),w)
	})
	http.ListenAndServe(":8080", r);
}
