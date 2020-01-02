package handlers

import (
	"net/http"
)

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	users := s.model.CountUsers()

	if users == 0 {
		http.Redirect(w, r, "/register", 302)
		return
	}

	http.Redirect(w, r, "/login", 302)
}
