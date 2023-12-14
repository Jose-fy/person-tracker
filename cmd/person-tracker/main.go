// cmd/person-tracker/main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"person-tracker/internal/api"
	"person-tracker/internal/db"
    "person-tracker/internal/model"

	"github.com/spf13/cobra"
)

func main() {
	db.InitDB("mydb.db")

    var rootCmd = &cobra.Command{Use: "myapp"}

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

			client := &openai.RealOpenAIClient{
				HTTPClient: &http.Client{},
			}

			message := "Say this is a test."

			result, err := client.SendMessageGPT(message)
			if err != nil {
                log.Fatal("Error inserting person: ", err)
            } else {

				fmt.Printf("%+v\n", result)
			}
		},
    }

    var cmdAskNaturalQuestion = &cobra.Command{
        Use: "find_person",
        Short: "Find the person given the context",
        Run: func (cmd *cobra.Command, args[]string)  {
            client := &openai.RealOpenAIClient{
                HTTPClient: &http.Client{},
            }

            people, err := db.QueryAllPeople()
            if err != nil {
                log.Fatal("Error querying all people: ", err)
            }

            people_s := model.SliceToString(people)

            reader := bufio.NewReader(os.Stdin)
            context := GetUserInputContext(reader)
            people_s += "\n" + "Tell me where I met a person in " + context

            fmt.Println(people_s)

            result, err := client.SendMessageGPT(people_s)
			if err != nil {
                log.Fatal("Error inserting person: ", err)
            }

            message := result.ParseOpenAIResponse()
            fmt.Printf("%+v\n", message)

            },
        }

    // Add more commands as needed
    rootCmd.AddCommand(cmdQueryAll, cmdInsert, cmdTalkOpenAI, cmdAskNaturalQuestion)

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
