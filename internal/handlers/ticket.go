package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"ticket/internal/util/struc"
	"ticket/internal/views"

	"github.com/gorilla/mux"
)

func (s *Server) newTicket(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth_user")
	userID := session.Values["UserID"].(int)

	switch r.Method {
	case http.MethodGet:
		views.TicketForm(w, "New Ticket", struc.Ticket{})
	case http.MethodPost:
		title := strings.TrimSpace(r.FormValue("title"))
		description := strings.TrimSpace(r.FormValue("description"))

		if len(title) > 0 && len(description) > 0 {
			ticket := &struc.Ticket{
				Title:       title,
				Description: description,
				UserID:      userID,
			}

			if err := s.model.NewTicket(ticket); err != nil {
				//error massage
				fmt.Fprintln(w, err)
			}

			http.Redirect(w, r, "/tickets", 302)
		} else {
			fmt.Fprintln(w, "plase fill all fields")
		}
	}
}

func (s *Server) tickets(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth_user")
	admin := session.Values["Admin"].(bool)
	userID := session.Values["UserID"].(int)

	var tickets []struc.Ticket
	if admin {
		tickets = s.model.Tickets()
	} else {
		tickets = s.model.TicketByUser(userID)
	}

	views.Tickets(w, tickets, admin)
}

func (s *Server) ticket(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth_user")
	userID := session.Values["UserID"].(int)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case http.MethodGet:
		ticket := s.model.Ticket(id)
		views.Ticket(w, ticket.Title, ticket)
	case http.MethodPost:
		comment := strings.TrimSpace(r.FormValue("comment"))

		if len(comment) > 0 {
			c := &struc.Comment{
				Reply:    comment,
				UserID:   userID,
				TicketID: id,
			}

			if err := s.model.NewComment(c); err != nil {
				//error massage
				fmt.Fprintln(w, err)
			}

			s.model.NewComment(c)
		} else {
			fmt.Fprintln(w, "plase fill the comment field")
		}

		http.Redirect(w, r, "/ticket/"+idStr, 302)
	}

}

func (s *Server) editTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	switch r.Method {
	case http.MethodGet:
		ticket := s.model.Ticket(id)
		views.TicketForm(w, "Edit", ticket)
	case http.MethodPost:
		title := strings.TrimSpace(r.FormValue("title"))
		description := strings.TrimSpace(r.FormValue("description"))

		s.model.UpdateByID(id, title, description)

		http.Redirect(w, r, "/ticket/edit/"+idStr, 302)
	}
}

func (s *Server) deleteTicket(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth_user")
	admin := session.Values["Admin"].(bool)

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	if admin {
		s.model.DeleteTicketByAdmin(id)
	} else {
		s.model.DeleteTicketByUser(id)
	}

	http.Redirect(w, r, "/tickets", 302)
}
