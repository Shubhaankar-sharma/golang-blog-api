package main

import (
	"github.com/Shubhaankar-sharma/golang-blog-api/api"
	"github.com/Shubhaankar-sharma/golang-blog-api/app"
	"github.com/Shubhaankar-sharma/golang-blog-api/users"
	"github.com/Shubhaankar-sharma/golang-blog-api/users/middlewares"
	_ "github.com/lib/pq"
)

func main() {
	newApp := app.App{}
	newApp.Init()

	newApp.Migrate()

	//app routers

	//newApp.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json
	ApiRouter := newApp.Router.PathPrefix("/api/").Subrouter()
	UserRouter := newApp.Router.PathPrefix("/user").Subrouter()
	ApiRouter.Use(middlewares.AuthJwtVerify)
	api.Router(ApiRouter, newApp.DB)
	users.Router(UserRouter, newApp.DB)

	//run app
	newApp.Run()
}
