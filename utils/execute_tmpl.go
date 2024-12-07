package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, pages []string, data any) {
	base_pages := []string{
		"views/components/navbar.html",
		"views/base.html",
	}
	pages = append(pages, base_pages...)
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
