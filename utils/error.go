package utils

import (
	"html/template"
	"net/http"
)


type ErrorResponse struct{
	ErrorNum int
	ErrorType string
}
func MessageError(w http.ResponseWriter, r *http.Request, code int , msg string){
	msg_err := &ErrorResponse{ErrorNum: code , ErrorType:msg}
	tmpl := template.Must(template.ParseGlob("views/*.html"))
	w.WriteHeader(code)
	tmpl.ExecuteTemplate(w,"error.html",msg_err)
}