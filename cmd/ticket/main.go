package main

import (
	"log"
	"net/http"
	"os"

	"ticket/internal/handlers"
	dbpkg "ticket/internal/util/config/db"
)

func main() {
	// db
	db := dbpkg.DBsetup()
	defer db.Close()
	// db.LogMode(true)

	// server
	s := handlers.NewServer(db)
	s.RegisterRoutes()

	portNum := os.Getenv("PORT")
	if portNum == "" {
		portNum = "2020"
	}
	log.Printf("Serving HTTP on port %s...\n", portNum)
	log.Fatal(http.ListenAndServe(":"+portNum, nil))
}
