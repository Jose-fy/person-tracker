package db
import (
    "testing"
    "person-tracker/internal/model"
    "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertPerson(t *testing.T) {
    // Create a mock database connection.
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock DB: %v", err)
    }
    defer db.Close()

    // Replace the global DB variable with the mock.
    DB = db
    defer func() {
        DB = nil // Reset the global DB variable when the test is done.
    }()

    // Create a sample person to insert.
    person := model.Person{
        Name:    "John Doe",
        Context: "Test context",
    }

    // Define the expected SQL query and mock behavior.
	mock.ExpectPrepare("INSERT INTO people").ExpectExec().WithArgs(person.Name, person.Context).WillReturnResult(sqlmock.NewResult(1, 1))


    // Call the InsertPerson function.
    err = InsertPerson(person)

    if err != nil {
        t.Errorf("InsertPerson failed: %v", err)
    }

    // Ensure that all expectations were met.
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Unfulfilled expectations: %v", err)
    }
}

// Repeat similar test functions for other db functions (e.g., QueryAllPeople, FindPeopleByName).


func TestQueryAllPeople(t *testing.T){

	db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("Error creating mock DB: %v", err)
    }
    defer db.Close()

    // Replace the global DB variable with the mock.
    DB = db
    defer func() {
        DB = nil // Reset the global DB variable when the test is done.
    }()

    mock.ExpectQuery("SELECT name, context FROM people").WillReturnRows(sqlmock.NewRows([]string{"name", "context"}))

	result, err := QueryAllPeople()

	if err != nil {
        t.Errorf("QueryAllPeople failed: %v", err)
    }

	assert.Equal(t, len(result), 0)

}
