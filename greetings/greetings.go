package greetings // declare a greetings package to collect related functions

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name") // Your caller will check the second value to see if an error occurred
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil // nil means no error
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) // initialize a map : make(map[key-type]value-type)

	// for后参数：the index of the current item in the loop and a copy of the item's value
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

// lowercase function name
// only to code in its own package (in other words, it's not exported)
func randomFormat() string {
	formats := []string{ // slice (like an array but its size changes dynamically as you add and remove items)
		"Hi, %v. Welcome!",
		"great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
