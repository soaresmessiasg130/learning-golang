package greetings

import "fmt"
import "rsc.io/quote"
import "errors"

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("The parameter \"name\" is an empty string.")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}

func ReturnName(name string) string {
	return name
}

func Quote() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
}
