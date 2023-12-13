package model


type Person struct {
	Name	string;
	Context 	string;
}

func SliceToString(persons []Person) string {
	var result string
	for _, person := range persons{
		result += person.Name + " | " + person.Context + "\n"
	}
	return result
}