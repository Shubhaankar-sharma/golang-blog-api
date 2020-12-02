package app

import (
	"fmt"
	"github.com/Shubhaankar-sharma/golang-blog-api/api/models"
	usermodels "github.com/Shubhaankar-sharma/golang-blog-api/users/user-models"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
func (app *App) Init() {
	r := mux.NewRouter().StrictSlash(true)
	//config
	config := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		goDotEnvVariable("DB_HOST"),
		goDotEnvVariable("DB_PORT"),
		goDotEnvVariable("DB_USERNAME"),
		goDotEnvVariable("DB_NAME"),
		goDotEnvVariable("DB_PASSWORD"))

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	app.DB = db
	app.Router = r
}
func (app *App) Migrate() {
	err := app.DB.Debug().AutoMigrate(&models.Blog{})
	if err != nil {
		log.Fatalln(color.FgRed.Render("An Error Occurred " + err.Error()))
	}
	err = app.DB.Debug().AutoMigrate(&usermodels.User{})
	if err != nil {
		log.Fatalln(color.FgRed.Render("An Error Occurred " + err.Error()))
	}
	log.Println(color.Info.Render("Migration Complete"))
}
func (app *App) Run() {
	log.Println(color.Info.Sprintf("Server has started on %s Press Ctrl C to Exit.", goDotEnvVariable("ADDR")))

	err := http.ListenAndServe(goDotEnvVariable("ADDR"), app.Router)
	if err != nil {
		log.Fatalln(color.FgRed.Render("An Error Occurred " + err.Error()))
	}
}
