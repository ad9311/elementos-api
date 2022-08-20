package server

import (
	"fmt"
	"net/http"
)

const (
	loginTemplate = "login.template.html"
)

func getLogin(w http.ResponseWriter, r *http.Request) {
	err := writeTemplate(w, loginTemplate)
	if err != nil {
		fmt.Println(err)
	}
}
