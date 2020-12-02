package handlers

import (
	"encoding/json"
	"github.com/Shubhaankar-sharma/golang-blog-api/api/models"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"log"
	"strconv"

	// "github.com/Shubhaankar-sharma/golang-blog-api/app"
	"gorm.io/gorm"

	// "gorm.io/gorm"
	"net/http"
)

type CreateBlogStruct struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// type a app.App
func CreateBlog(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var b CreateBlogStruct

		err := json.NewDecoder(r.Body).Decode(&b)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		newblog := models.Blog{
			Author: b.Author,
			Title:  b.Title,
			Body:   b.Body}

		newB, err := newblog.Save(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// _, err = fmt.Fprintf(w, "New Blog Created: Title: %s ID: %v Author: %s Content: %s",newblog.Title,newblog.ID,newblog.Author,newblog.Body)
		err = json.NewEncoder(w).Encode(&newB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		return
	}
}

func GetBlog(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Blog
		vars := mux.Vars(r)
		valUint, _ := strconv.ParseInt(vars["id"], 1, 64)
		newB, err := b.Get(int(valUint), db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(&newB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		return
	}
}
func GetListBlog(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Blog
		newB, err := b.GetAll(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(&newB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func UpdateBlog(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data CreateBlogStruct

		vars := mux.Vars(r)
		valUint, _ := strconv.ParseInt(vars["id"], 10, 64)
		err := json.NewDecoder(r.Body).Decode(&data)
		log.Println(color.Debug.Render(valUint))
		log.Println(color.Debug.Render(vars))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		b := models.Blog{
			Author: data.Author,
			Title:  data.Title,
			Body:   data.Body}
		newB, err := b.Put(int(valUint), db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		err = json.NewEncoder(w).Encode(&newB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		return
	}
}
