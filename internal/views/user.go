package views

import (
	"log"
	"net/http"
)

func NewUser(w http.ResponseWriter) {
	data := map[string]interface{}{"Title": "Register"}

	err := ParseTmpl(w, data, "register_login", "register")
	if err != nil {
		log.Println(err)
	}
}

func Login(w http.ResponseWriter) {
	data := map[string]interface{}{"Title": "Login"}

	err := ParseTmpl(w, data, "register_login", "login")
	if err != nil {
		log.Println(err)
	}
}
