package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Printf("Welcome to Amazing Numbers!\n\n")
	printHelp()

	for {
		input := requestInput()
		fmt.Println()

		if input == "0" {
			fmt.Println("Goodbye!")
			return
		}

		if input == "" {
			printHelp()
			continue
		}

		inputs := strings.Split(input, " ")

		isValid, message := isInputsValid(inputs)
		if !isValid {
			fmt.Println(message)
			continue
		}

		switch len(inputs) {
		case 0:
			printHelp()
			continue

		case 1:
			number, _ := strconv.Atoi(inputs[0])
			fmt.Println(getSingleNumberProperties(number))
			continue

		case 2:
			number1, _ := strconv.Atoi(inputs[0])
			number2, _ := strconv.Atoi(inputs[1])
			fmt.Println(getConsecutiveProperties(number1, number2))
			continue

		default:
			number1, _ := strconv.Atoi(inputs[0])
			number2, _ := strconv.Atoi(inputs[1])
			fmt.Println(getNumbersWithMultiProperties(number1, number2, inputs[2:]))
			continue
		}

	}

}
