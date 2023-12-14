package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestSliceToString(t *testing.T){
	persons := []Person{
		{Name: "John", Context: "House"},
		{Name: "Jane", Context: "Supermarket"},
	}

	result := SliceToString(persons)

	assert.Equal(t, "John | House\nJane | Supermarket\n", result)
}