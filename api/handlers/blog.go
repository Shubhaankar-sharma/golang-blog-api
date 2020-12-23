package handlers

import (
	"encoding/json"
	"github.com/Shubhaankar-sharma/golang-blog-api/api/models"
	user_models "github.com/Shubhaankar-sharma/golang-blog-api/users/user-models"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"log"
	"strconv"

	// "github.com/Shubhaankar-sharma/golang-blog-api/app"
	"gorm.io/gorm"

	// "gorm.io/gorm"
	"net/http"
)


// serializer structs
type CreateBlogStruct struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// utils
func GetUser(r *http.Request, db *gorm.DB) (*models.ApiUser, error){
	usrid := r.Context().Value("userID").(float64)
	user := user_models.User{}
	newUsr, err := user.GetUserID(uint(usrid), db)
	if err != nil {
		return &models.ApiUser{}, err
	}
	apiusr := models.ApiUser{
		UId:    newUsr.ID,
		Email: newUsr.Email,
		Name:  newUsr.Name,
	}
	return &apiusr,nil
}

// Request Handlers

// GET REQUESTS

func GetBlog(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var b models.Blog

		vars := mux.Vars(r)

		valUint, _ := strconv.ParseInt(vars["id"], 10, 64)

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
		blogs, err := b.GetAll(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(&blogs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}


func GetUserListBlogs(db *gorm.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var b models.Blog
		newUsr, err := GetUser(r,db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		b.Author = *newUsr
		blogs, err := b.GetAllFromUser(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(&blogs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}



// PUT REQUESTS

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

		apiusr, err := GetUser(r, db)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		b := models.Blog{
			Author: *apiusr,
			Title: 	data.Title,
			Body:  	data.Body,
		}

		_, err = b.Put(int(valUint), db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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


// POST REQUESTS

func CreateBlog(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var b CreateBlogStruct

		err := json.NewDecoder(r.Body).Decode(&b)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		newUsr, err := GetUser(r,db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		newblog := models.Blog{
			Author: *newUsr,
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