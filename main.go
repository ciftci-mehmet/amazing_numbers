package main

import (
	"bufio"
	"fmt"
	"os"
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
		if !isValid{
			fmt.Println(message)
			continue
		}

		switch len(inputs) {
		case 1:
			number, _ := strconv.Atoi(inputs[0])

			fmt.Println("Properties of", number)
			fmt.Printf("\teven: %v\n", isEvenNumber(number))
			fmt.Printf("\todd: %v\n", isOddNumber(number))
			fmt.Printf("\tbuzz: %v\n", isBuzzNumber(number))
			fmt.Printf("\tduck: %v\n", isDuckNumber(number))
			fmt.Printf("\tpalindromic: %v\n", isPalindromicNumber(number))

			continue

		case 2:
			number1, _ := strconv.Atoi(inputs[0])
			number2, _ := strconv.Atoi(inputs[1])
			fmt.Println(getConsecutiveProperties(number1, number2))

			continue

		case 3:
			number1, _ := strconv.Atoi(inputs[0])
			number2, _ := strconv.Atoi(inputs[1])
			fmt.Println(getNumbersWithProperty(number1, number2, inputs[2]))

			continue
		}

	}

}

func printHelp() {
	helpText :=
		"Supported requests:\n" +
			" - enter a natural number to know its properties;\n" +
			" - enter two natural numbers to obtain the properties of the list:\n" +
			"   * the first parameter represents a starting number;\n" +
			"   * the second parameter shows how many consecutive numbers are to be processed;\n" +
			" - two natural numbers and a property to search for;\n" +
			" - two natural numbers and two properties to search for;\n" +
			" - separate the parameters with one space\n" +
			" - enter 0 to exit."
	fmt.Println(helpText)
}

func requestInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Printf("Enter a request: ")
	scanner.Scan()

	return scanner.Text()
}

func isInputsValid(inputs []string) (bool, string){
	switch len(inputs) {
	case 4:
		_, ok := propertiesMap[inputs[3]]
		if !ok {
			
			message := fmt.Sprintf("The property [%s] is wrong.\n", inputs[3])
			properties := ""
			for k := range propertiesMap {
				properties += k + ", "
			}
			properties = properties[:len(properties)-2]
			message += fmt.Sprintf("Available properties: %v", properties)

			return false, message
		}
		fallthrough

	case 3:
		_, ok := propertiesMap[inputs[2]]
		if !ok {
			
			message := fmt.Sprintf("The property [%s] is wrong.\n", inputs[2])
			properties := ""
			for k := range propertiesMap {
				properties += k + ", "
			}
			properties = properties[:len(properties)-2]
			message += fmt.Sprintf("Available properties: %v", properties)

			return false, message
		}
		fallthrough
		
	case 2:
		if !isNaturalNumber(inputs[1]) {
			return false, "The second parameter should be a natural number."
		}
		fallthrough

	case 1:
		if !isNaturalNumber(inputs[0]) {
			return false, "The first parameter should be a natural number or zero."
		}
	}
	return true, ""
}