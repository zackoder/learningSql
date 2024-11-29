package controllers

import (
	"net/http"

	"forum/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	page := []string{"views/pages/home.html"}
	utils.ExecuteTemplate(w, page, nil)

	if r.URL.Path != "/" {
		utils.MessageError(w, r, http.StatusNotFound, "Page Not Found!!!!")
		return
	}
}
