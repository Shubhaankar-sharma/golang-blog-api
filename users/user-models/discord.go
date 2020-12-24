package user_models

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	//"strconv"
	//"strings"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type Discord struct {
	gorm.Model
	DiscordId    int    `gorm:"type:integer;unique_index" json:"discord_id"`
	DiscordToken string `gorm:"type:varchar(100);unique_index" json:"discord_token"`
	AvatarHash   string `gorm:"type:varchar(100);unique_index" json:"avatar_hash"`
}

var DISCORD_CLIENT_ID string = goDotEnvVariable("DISCORD_CLIENT_ID")
var DISCORD_SECRET string = goDotEnvVariable("DISCORD_SECRET")
var API_ENDPOINT string = "https://discord.com/api/v6"

func Exchange_code(code string) interface{} {

	REDIRECT_URI := "http://" + goDotEnvVariable("ADDR") + "/user/discord/connect"

	data := url.Values{}
	data.Set("client_id", DISCORD_CLIENT_ID)
	data.Set("client_secret", DISCORD_SECRET)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", REDIRECT_URI)
	data.Set("scope", "identify email connections")

	resp, err := http.PostForm(API_ENDPOINT+"/oauth2/token", data)

	if err != nil {
		log.Fatalf("An Error Occured %v", err.Error())
	}

	defer resp.Body.Close()

	var newresult map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&newresult)

	fmt.Println(newresult["access_token"])

	access_token := newresult["access_token"]

	url := API_ENDPOINT + "/users/@me"

	newreq, err := http.NewRequest("GET", url, nil)
	newreq.Header.Add("Authorization", fmt.Sprintf("Bearer %v", access_token))
	res, err := http.DefaultClient.Do(newreq)

	if err != nil {
		log.Fatalf("An Error Occured %v", err.Error())
	}

	defer res.Body.Close()
	log.Println(res.StatusCode)
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	var finalresult map[string]interface{}
	json.NewDecoder(res.Body).Decode(&finalresult)
	return finalresult
}
