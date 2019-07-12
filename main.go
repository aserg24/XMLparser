package main

import (
	"XMLparser/db"
	"XMLparser/ivi"
	"log"

	"github.com/go-pg/pg"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "myShows"}
	dbConnect := pg.Connect(cfg)
	defer dbConnect.Close()
	dbc := db.DB{dbConnect}
	loader := ivi.NewIviLoader("document", "tables", &dbc)

	err := loader.Parse()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PARSED")
	err = loader.Save()
	if err != nil {
		log.Println(err)
	}
}
