package user_handler

import (
	"encoding/json"
	"fmt"
	usermodels "github.com/Shubhaankar-sharma/golang-blog-api/users/user-models"
	"github.com/Shubhaankar-sharma/golang-blog-api/users/utils"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/peterwade153/ivents/api/responses"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello From API")
	if err != nil {
		log.Panicln(color.Info.Render(err.Error()))
	}
}

func SignUp(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp = map[string]interface{}{"status": "success", "message": "Registered successfully"}
		user := &usermodels.User{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		err = json.Unmarshal(body, &user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		usr, _ := user.GetUser(db)
		if usr != nil {
			resp["status"] = "failed"
			resp["message"] = "User already registered, please login"
			responses.JSON(w, http.StatusBadRequest, resp)
			return
		}
		user.Prepare()
		err = user.Validate("")

		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		err = user.BeforeSave()
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		userCreated, err := user.SaveUser(db)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		resp["user"] = userCreated
		responses.JSON(w, http.StatusCreated, resp)
		return
	}
}

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp = map[string]interface{}{"status": "success", "message": "logged in"}

		user := &usermodels.User{}

		body, err := ioutil.ReadAll(r.Body) // read user input from request
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		err = json.Unmarshal(body, &user)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		user.Prepare() // here strip the text of white spaces

		err = user.Validate("login") // fields(email, password) are validated
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		usr, err := user.GetUser(db)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		if usr == nil { // user is not registered
			resp["status"] = "failed"
			resp["message"] = "Login failed, please signup"
			responses.JSON(w, http.StatusBadRequest, resp)
			return
		}
		err = usermodels.CheckPasswordHash(user.Password, usr.Password)
		if err != nil {
			resp["status"] = "failed"
			resp["message"] = "Login failed, please try again"
			responses.JSON(w, http.StatusForbidden, resp)
			return
		}
		token, err := utils.EncodeAuthToken(usr.ID)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		resp["token"] = token
		responses.JSON(w, http.StatusOK, resp)
		return

	}
}
