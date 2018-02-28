package main

import (
	"fmt"
	"log"

	"github.com/gobuffalo/pop"
	"github.com/jubalh/db2-test/models"
)

func main() {
	fmt.Println("This is a Go database test program")

	db, err := pop.Connect("development")
	if err != nil {
		log.Panic(err)
	}

	user := models.User{Firstname: "Vincent", Lastname: "Vega"}
	_, err = db.ValidateAndSave(&user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Added %s to database\n", user.Firstname)

	user = models.User{Firstname: "Jules", Lastname: "Winnfield"}
	_, err = db.ValidateAndSave(&user)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Added %s to database\n", user.Firstname)
}
