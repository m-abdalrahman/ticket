package models

import (
	"ticket/internal/util/struc"
)

func (m *Model) NewTicket(ticket *struc.Ticket) error {
	return m.db.Create(ticket).Error
}

func (m *Model) Ticket(id int) struc.Ticket {
	var ticket struc.Ticket

	m.db.Where("id = ?", id).
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		Find(&ticket)

	return ticket
}

func (m *Model) Tickets() []struc.Ticket {
	tickets := []struc.Ticket{}

	m.db.Preload("User").Find(&tickets)

	return tickets
}

func (m *Model) TicketByUser(id int) []struc.Ticket {
	tickets := []struc.Ticket{}

	m.db.
		Where("deleted = ?", false).
		Where("user_id = ?", id).
		Find(&tickets)

	return tickets
}

func (m *Model) UpdateByID(id int, title, description string) error {
	return m.db.
		Model(&struc.Ticket{}).
		Where("id = ?", id).
		Updates(struc.Ticket{Title: title, Description: description}).Error
}

func (m *Model) DeleteTicketByUser(id int) error {
	return m.db.
		Model(&struc.Ticket{}).
		Where("id = ?", id).
		Updates(struc.Ticket{Deleted: true, Status: "Deleted"}).Error
}

func (m *Model) DeleteTicketByAdmin(id int) error {
	return m.db.
		Where("id = ?", id).
		Delete(struc.Ticket{}).Error
}
