package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"ticket/internal/util/struc"
	"ticket/internal/views"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("test"))

func (s *Server) newUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth_user")
	if !session.IsNew {
		admin := session.Values["Admin"].(bool)
		if !admin {
			http.Redirect(w, r, "/ticket/new", 302)
			return
		}

	}

	switch r.Method {
	case http.MethodGet:
		views.NewUser(w)
	case http.MethodPost:
		fullname := strings.TrimSpace(r.FormValue("fullname"))
		password := strings.TrimSpace(r.FormValue("password"))
		email := strings.TrimSpace(r.FormValue("email"))

		if len(fullname) > 0 && len(password) > 0 && len(email) > 0 {
			user := &struc.User{
				FullName: fullname,
				Email:    email,
				Password: password,
			}

			if err := s.model.NewUser(user); err != nil {
				//error massage
				fmt.Fprintln(w, err)
			}

			http.Redirect(w, r, "/login", 302)
		} else {
			fmt.Fprintln(w, "plase fill all fields")
		}
	}
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth_user")
	if !session.IsNew {
		admin := session.Values["Admin"].(bool)
		if admin {
			http.Redirect(w, r, "/tickets", 302)
			return
		}

		http.Redirect(w, r, "/ticket/new", 302)
		return
	}

	switch r.Method {
	case http.MethodGet:
		views.Login(w)
	case http.MethodPost:
		email := strings.TrimSpace(r.FormValue("email"))
		password := strings.TrimSpace(r.FormValue("password"))

		if isEmpty(email) || isEmpty(password) {
			//error massage
			fmt.Fprintln(w, "Please fill all fields")
		} else {
			if validEmail(email) {
				u, err := s.model.GetUserByEmailPass(email, password)
				if err != nil {
					//error massage
					fmt.Fprintln(w, err)
				} else {
					session.Values["UserID"] = u.ID
					session.Values["Admin"] = u.Admin
					session.Save(r, w)

					http.Redirect(w, r, "/tickets", 302)
				}
			} else {
				//error massage
				fmt.Fprintln(w, "Please enter a valid email address")
			}
		}
	}
}

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth_user")
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
}
