package api

import (
	"github.com/Shubhaankar-sharma/golang-blog-api/api/handlers"
	// "github.com/Shubhaankar-sharma/golang-blog-api/app"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("", handlers.Home).Methods("GET")
	r.HandleFunc("/blog/", handlers.CreateBlog(db)).Methods("POST")
	r.HandleFunc("/blog/{id}", handlers.GetBlog(db)).Methods("GET")
	r.HandleFunc("/blog/{id}", handlers.UpdateBlog(db)).Methods("PUT")
	r.HandleFunc("/blogs", handlers.GetListBlog(db)).Methods("GET")
}
