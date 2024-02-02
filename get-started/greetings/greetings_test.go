package greetings

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	assert := assert.New(t)
	name := "John"
	message, err := Hello(name)

	if err != nil {
		log.Fatal(err)
	}

	assert.Containsf(message, "John", "Should contains the name")
}

func TestHelloEmpty(t *testing.T) {
	assert := assert.New(t)
	name := ""
	_, err := Hello(name)

	assert.NotNil(err)
}

func TestHellosNames(t *testing.T) {
	assert := assert.New(t)
	names := []string{"John", "Juliana"}
	messages, err := Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	assert.Containsf(messages["John"], "John", "Should contains the name")
	assert.Containsf(messages["Juliana"], "Juliana", "Should contains the name")
}

func TestHellosEmpty(t *testing.T) {
	assert := assert.New(t)
	names := []string{"John", ""}
	_, err := Hellos(names)

	assert.NotNil(err)
}
