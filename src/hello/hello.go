package main

import (
	"awesomeProject/src/greetings"
	"awesomeProject/src/letters"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	name := letters.MyString("Sam Anderson")
	var v letters.VowelsFinder
	v = name
	fmt.Printf("Vowels are %c\n", v.FindVowels())
}
