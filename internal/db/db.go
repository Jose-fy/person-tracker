package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "log"
	"person-tracker/internal/model"
)

var DB *sql.DB


func InitDB(dataSourceName string){
	var err error

	DB, err = sql.Open("sqlite3", dataSourceName)

	if err != nil {

		log.Fatal(err)
	}

	createTables()
}


func createTables(){
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS people (name TEXT, context TEXT)`)

	if err != nil {
		log.Fatal(err)
	}
}

// func insertEntry(){


// }


func QueryAllPeople() ([]model.Person, error) {

	var people []model.Person

	rows, err := DB.Query(`SELECT name, context FROM people`)

	if err != nil {
        return nil, err // return an error here
    }

	defer rows.Close()

	for rows.Next() {
		var p model.Person
		err := rows.Scan(&p.Name, &p.Context)

		if err != nil {
			return nil, err // return an error here
		}
		people = append(people, p)
	}

    if err = rows.Err(); err != nil {
        return nil, err // handle any error encountered during iteration
    }

    return people, nil
}

