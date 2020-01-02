package struc

type Ticket struct {
	ID          int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Title       string    `gorm:"NOT NULL"`
	Description string    `gorm:"NOT NULL"`
	Status      string    `gorm:"NOT NULL;DEFAULT:New"`
	Deleted     bool      `gorm:"DEFAULT:false"`
	Comments    []Comment `gorm:"FOREIGNKEY:TicketID"`
	UserID      int
	User        User
}

// db ticket table name
func (Ticket) TableName() string {
	return "tickets"
}
