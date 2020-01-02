package handlers

import (
	"net/http"
	"os"

	"ticket/internal/models"

	middleware "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Server ...
type Server struct {
	model *models.Model
}

// NewServer return new server with database connection
func NewServer(db *gorm.DB) *Server {
	m := models.NewModel(db)
	return &Server{model: m}
}

// RegisterRoutes ...
func (s *Server) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", s.index)

	r.HandleFunc("/register", s.newUser)
	r.HandleFunc("/login", s.login)
	r.HandleFunc("/logout", s.logout)

	r.HandleFunc("/ticket/new", s.newTicket)
	r.HandleFunc("/tickets", s.tickets)
	r.HandleFunc("/ticket/{id}", s.ticket)
	r.HandleFunc("/ticket/edit/{id}", s.editTicket)
	r.HandleFunc("/ticket/delete/{id}", s.deleteTicket).Methods(http.MethodPost)

	// static files
	r.PathPrefix("/web/").
		Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web/"))))

	http.Handle("/", middleware.LoggingHandler(os.Stdout, r))

	return r
}
