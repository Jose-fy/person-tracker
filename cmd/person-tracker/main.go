// cmd/person-tracker/main.go
package main

import (
	"fmt"
	"log"
	"person-tracker/internal/db"
)

func main() {
	db.InitDB("mydb.db")

	person := GetUserInput()

	fmt.Printf("%+v\n", person)
    db.InsertPerson(person)

	// Rest of your code...
	people, err := db.QueryAllPeople()

	if err != nil {
		log.Fatal("Error")
	}

	if len(people) != 0 {
		for _, person := range people {
			fmt.Printf("%+v\n", person)
		}
	}
}
