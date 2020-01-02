package struc

type Comment struct {
	ID       int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Reply    string `gorm:"NOT NULL"`
	TicketID int
	UserID   int
	User     User
}

// db comment table name
func (Comment) TableName() string {
	return "comments"
}
