package struc

type User struct {
	ID       int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	FullName string    `gorm:"NOT NULL"`
	Email    string    `gorm:"NOT NULL;UNIQUE"`
	Password string    `gorm:"NOT NULL"`
	Admin    bool      `gorm:"DEFAULT:false"`
	Tickets  []Ticket  `gorm:"FOREIGNKEY:UserID"`
	Comments []Comment `gorm:"FOREIGNKEY:UserID"`
}

// db user table name
func (User) TableName() string {
	return "users"
}
