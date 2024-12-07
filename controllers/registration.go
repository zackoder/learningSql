package controllers

import (
	"fmt"
	"net/http"

	"forum/utils"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	html := `<script>alert("%s")</script> `

	page := []string{"views/pages/register.html"}
	var info utils.Info
	if r.Method == "POST" {
		info.Name = r.FormValue("name")
		info.Username = r.FormValue("username")
		info.Email = r.FormValue("email")
		info.Password = r.FormValue("password")
		info.Password2 = r.FormValue("password2")
	}
	if info.Password != info.Password2 {
		msj := fmt.Sprintf(html, "password error")
		fmt.Fprint(w, msj)
	}
	// if !isValidEmail(info.Email) {
	// 	msj2 := fmt.Sprintf(html, "incorect email format")
	// 	fmt.Fprint(w, msj2)
	// }
	utils.ExecuteTemplate(w, page, &info)
}

// func isValidEmail(email string) bool {
// 	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
// 	return re.MatchString(email)
// }
