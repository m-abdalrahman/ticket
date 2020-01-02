package views

import (
	"log"
	"net/http"

	"ticket/internal/util/struc"
)

func TicketForm(w http.ResponseWriter, PageTitle string, value struc.Ticket) {
	data := make(map[string]interface{})

	if PageTitle == "New Ticket" {
		data["PageTitle"] = PageTitle
		data["Button"] = "New Ticket"
	} else {
		data["PageTitle"] = PageTitle
		data["Button"] = "Edit"
		data["Title"] = value.Title
		data["Description"] = value.Description
	}

	err := ParseTmpl(w, data, "main", "ticket_form")
	if err != nil {
		log.Println(err)
	}
}

func Ticket(w http.ResponseWriter, PageTitle string, value struc.Ticket) {
	data := make(map[string]interface{})

	data["PageTitle"] = PageTitle
	data["Title"] = value.Title
	data["Description"] = value.Description
	data["FullName"] = value.User.FullName
	data["Comments"] = value.Comments

	err := ParseTmpl(w, data, "main", "ticket")
	if err != nil {
		log.Println(err)
	}
}

func Tickets(w http.ResponseWriter, tickets []struc.Ticket, admin bool) {
	data := map[string]interface{}{"PageTitle": "Tickets", "Tickets": tickets}

	if admin {
		data["Admin"] = "Admin"
	}

	err := ParseTmpl(w, data, "main", "tickets")
	if err != nil {
		log.Println(err)
	}
}
