package users

import (
	user_handler "github.com/Shubhaankar-sharma/golang-blog-api/users/user-handler"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Router(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("", user_handler.Home).Methods("GET")
	r.HandleFunc("/login", user_handler.Login(db)).Methods("POST")
	r.HandleFunc("/signup", user_handler.SignUp(db)).Methods("POST")
}
