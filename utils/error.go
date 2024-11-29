package utils

import (
	"net/http"
)

type ErrorResponse struct {
	ErrorNum  int
	ErrorType string
}

func MessageError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	msg_err := &ErrorResponse{ErrorNum: code, ErrorType: msg}
	page := []string{"views/pages/errr.html"}
	ExecuteTemplate(w, page, msg_err)
}
