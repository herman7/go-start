package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Set properties of the predefined Logger, including the log entry prefix 
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // a flag to disable printing the time, source file, and line number.

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	// message, err := greetings.Hello("Gladys")
	if err != nil {
		// If an error was returned, print it to the console and exit the programs
		log.Fatal(err)
	}
	
	fmt.Println(messages)
}