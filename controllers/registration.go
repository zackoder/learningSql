package controllers

import (
	"net/http"
	"forum/utils"
)

func Registration(w http.ResponseWriter, r http.Request) {
	page := []string{"views/pages/register.html"}
	var info struct{
		Name string
		Username string
		Email string
		Password string
		Password2 string
	}
	if r.Method == "POST"{
		 info.Name = r.PostFormValue("name")
		 info.Username= r.PostFormValue("username")
		 info.Email= r.PostFormValue("email")
		 info.Password = r.PostFormValue("password")
		 info.Password2 = r.PostFormValue("password2")
	}
	utils.ExecuteTemplate( w , page, &info)
}
