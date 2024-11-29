package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, pages []string, data any) {
	pages = append(pages, "views/base.html")
	tmpl, err := template.ParseFiles(pages...)
	if err != nil {
		fmt.Println(err, "we can't excute this page")
		return
	}
	err = tmpl.ExecuteTemplate(w,"base",data) 
	if err != nil {
		fmt.Println(err, "we can't excute this page")
		return
	}
}
