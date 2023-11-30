package main

import (
	"bufio"
	"fmt"
	"os"
	"person-tracker/internal/model"
	"strings"
)

func GetUserInput() model.Person {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What's the name of the person you met ?")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Where did you meet this person ?")
	context, _ := reader.ReadString('\n')
	context = strings.TrimSpace(context)

	return model.Person{
		Name:    name,
		Context: context,
	}
}
