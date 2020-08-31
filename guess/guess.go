package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please think of a number between 1 and 100")
	fmt.Println("Press ENTER when ready")
	scanner.Scan()
	//
	guess := 50
	for {
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high?")
		fmt.Println("(b) too low?")
		fmt.Println("(c) correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			guess--
		} else if response == "b" {
			guess++
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Invalid Response, Try Again.")
		}

	}
}
