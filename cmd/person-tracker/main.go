// cmd/person-tracker/main.go
package main

import (
	"fmt"
	"log"
	"person-tracker/internal/db"
	"person-tracker/internal/api"
	"bufio"
	"os"
	"github.com/spf13/cobra"
	"net/http"


)

func main() {
	db.InitDB("mydb.db")

    var rootCmd = &cobra.Command{Use: "myapp"}

	client := &http.Client{}


    var cmdQueryAll = &cobra.Command{
        Use:   "queryall",
        Short: "Query all people from the database",
        Run: func(cmd *cobra.Command, args []string) {
            people, err := db.QueryAllPeople()
            if err != nil {
                log.Fatal("Error querying all people: ", err)
            }
            for _, person := range people {
                fmt.Printf("%+v\n", person)
            }
        },
    }

    var cmdInsert = &cobra.Command{
        Use:   "insert",
        Short: "Insert a new person into the database",
        Run: func(cmd *cobra.Command, args []string) {
            reader := bufio.NewReader(os.Stdin)
            person := GetUserInput(reader)
            err := db.InsertPerson(person)
            if err != nil {
                log.Fatal("Error inserting person: ", err)
            }
        },
    }

	var cmdTalkOpenAI = &cobra.Command{
		Use:  "openai",
		Short: "Talk with OpenAI's api (chatgpt)",
		Run: func(cmd *cobra.Command, args []string) {
			result, err := openai.SendMessageGPT(*client, "Say this is a test")
			if err != nil {
                log.Fatal("Error inserting person: ", err)
            } else {

				fmt.Printf("%+v\n", result)
			}
		},
	}

    // Add more commands as needed
    rootCmd.AddCommand(cmdQueryAll, cmdInsert, cmdTalkOpenAI)

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
