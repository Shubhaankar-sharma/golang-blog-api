package user_handler

import (
	//"encoding/json"
	"fmt"
	usermodels "github.com/Shubhaankar-sharma/golang-blog-api/users/user-models"
	"github.com/gookit/color"
	//"gorm.io/gorm"
	//"io/ioutil"
	//	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"github.com/peterwade153/ivents/api/responses"
)

func DiscordRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://discord.com/api/oauth2/authorize?client_id=752952469654601768&redirect_uri=http%3A%2F%2F127.0.0.1%3A8000%2Fuser%2Fdiscord%2Fconnect&response_type=code&scope=identify%20email%20connections", http.StatusFound)
}

func DiscordCode(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	sb := usermodels.Exchange_code(code)
	_, err := fmt.Fprintln(w, sb)
	if err != nil {
		log.Panicln(color.Info.Render(err.Error()))
	}

}
