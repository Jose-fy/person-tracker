// cmd/person-tracker/main.go
package main

import (
    "log"
    "person-tracker/internal/db"
    "fmt"
)

func main() {
    db.InitDB("mydb.db")
    // Rest of your code...
	people, err := db.QueryAllPeople()

    if err != nil{
        log.Fatal("Error")
    }

    if len(people) != 0 {
        for _, person := range(people){
            fmt.Printf("%+v\n", person)
        }
    }
}
