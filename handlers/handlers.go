package handlers

import (
	"net/http"

	"github.com/seshhh/tsawlercourse/render"
)


func Home(rw http.ResponseWriter, r *http.Request){
	render.Render(rw, "homePage.html")
}

func About(rw http.ResponseWriter, r *http.Request){
	render.Render(rw, "aboutPage.html")
}

func Goiz(rw http.ResponseWriter, r *http.Request){
	render.Render(rw, "goizPage.html")
}


