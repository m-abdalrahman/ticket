package db

import (
	"log"

	"ticket/internal/util/struc"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	fileName = "store.db"
)

func DBsetup() *gorm.DB {
	db, err := gorm.Open("sqlite3", fileName)
	if err != nil {
		panic("failed to connect database")
	}

	if !db.HasTable("users") || !db.HasTable("tickets") || !db.HasTable("comments") {
		db.CreateTable(&struc.User{}, &struc.Ticket{}, &struc.Comment{})
		log.Println("database tables has been created")
	}

	return db
}
