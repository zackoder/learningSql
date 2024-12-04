package controllers

import (
	"encoding/json"
	"net/http"

	"forum/utils"
)



func Registration(w http.ResponseWriter, r *http.Request) {
	page := []string{"views/pages/register.html"}
	var info utils.Info
	if r.Method == "POST" {
		info.Name = r.FormValue("name")
		info.Username = r.FormValue("username")
		info.Email = r.FormValue("email")
		info.Password = r.FormValue("password")
		info.Password2 = r.FormValue("password2")
	}

	utils.ExecuteTemplate(w, page, &info)
}

func Test(w http.ResponseWriter, r *http.Request) {
	var info utils.Info
	json.NewEncoder(w).Encode(&info)
}
