package controllers

import (
	"net/http"

	"forum/utils"
)

type info struct {
	Name      string
	Username  string
	Email     string
	Password  string
	Password2 string
}

func Registration(w http.ResponseWriter, r *http.Request) {
	page := []string{"views/pages/register.html"}
	var info info
	if r.Method == "POST" {
		info.Name = r.FormValue("name")
		info.Username = r.FormValue("username")
		info.Email = r.FormValue("email")
		info.Password = r.FormValue("password")
		info.Password2 = r.FormValue("password2")
	}
	utils.ExecuteTemplate(w, page, &info)
}
