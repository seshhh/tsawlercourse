package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(rw http.ResponseWriter, tmpl string){
	parsedTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemp.Execute(rw, nil)
	if err != nil{
		fmt.Println("error parsing template", err)
		return
	}
}
