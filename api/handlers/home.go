package handlers

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello From API")
	if err != nil {
		log.Panicln(color.Info.Render(err.Error()))
	}
}
