package controllers

import (
	"net/http"

	"forum/utils"
)
 type User struct{
	Email string
	Password string
 }
func Login(w http.ResponseWriter, r *http.Request) {
	page := []string{"views/pages/login.html"}
	var user User 
	if r.Method == "POST"{
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
	}

	utils.ExecuteTemplate(w, page, user)
}
