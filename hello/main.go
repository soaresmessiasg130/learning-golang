package main

import (
	"fmt"
	"log"

	"messiasspp.com/greetings"
)

func main() {
	nameOfMyFriend := greetings.ReturnName("Name of a Friend to Return")

	fmt.Println(nameOfMyFriend)

	greetings.Quote()

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, error := greetings.Hello("Fran Games")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)

	message2, error2 := greetings.Hello("")

	if error2 != nil {
		log.Fatal(error2)
	}

	fmt.Println(message2)
}
