package controllers

import (
	"fmt"
	"net/http"

	"forum/utils"
	
)
 
func Login(w http.ResponseWriter, r *http.Request) {
	page := []string{"views/pages/login.html"}
	var G  utils.User 
	if r.Method == "POST"{
		G.Email = r.FormValue("email")
		G.Password = r.FormValue("password")
	}
	utils.ExecuteTemplate(w, page, &G)
	fmt.Println(G.Email)
	fmt.Println(G.Password)
}
