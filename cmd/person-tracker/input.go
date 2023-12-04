package main

import (
	"bufio"
	"fmt"
	"person-tracker/internal/model"
	"strings"
)

func GetUserInput(reader *bufio.Reader) model.Person {


	name := GetUserInputName(reader)

	context := GetUserInputContext(reader)

	return model.Person{
		Name:    name,
		Context: context,
	}
}


func GetUserInputName(reader *bufio.Reader) string {

	fmt.Println("What's the name of the person you met ?")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return name
}


func GetUserInputContext(reader *bufio.Reader) string {

	fmt.Println("Where did you meet this person ?")
	context, _ := reader.ReadString('\n')
	context = strings.TrimSpace(context)

	return context

}