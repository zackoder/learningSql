package controllers

import (
	"net/http"

	"forum/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	
	Page := []string{"views/pages/login.html"}
	utils.ExecuteTemplate(w, Page, nil)
}
