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

		// multiple input check
		inputs := strings.Split(input, " ")
		switch len(inputs) {
		case 1:
			if !isNaturalNumber(inputs[0]) {
				fmt.Println("The first parameter should be a natural number or zero.")
				continue
			}

			number, _ := strconv.Atoi(inputs[0])

			fmt.Println("Properties of", number)
			fmt.Printf("\teven: %v\n", isEvenNumber(number))
			fmt.Printf("\todd: %v\n", isOddNumber(number))
			fmt.Printf("\tbuzz: %v\n", isBuzzNumber(number))
			fmt.Printf("\tduck: %v\n", isDuckNumber(number))
			fmt.Printf("\tpalindromic: %v\n", isPalindromicNumber(number))

			continue

		case 2:
			if !isNaturalNumber(inputs[0]) {
				fmt.Println("The first parameter should be a natural number or zero.")
				continue
			}
			if !isNaturalNumber(inputs[1]) {
				fmt.Println("The second parameter should be a natural number.")
				continue
			}

			number1, _ := strconv.Atoi(inputs[0])
			number2, _ := strconv.Atoi(inputs[1])
			fmt.Println(getConsecutiveProperties(number1, number2))

			continue

		case 3:
			if !isNaturalNumber(inputs[0]) {
				fmt.Println("The first parameter should be a natural number or zero.")
				continue
			}
			if !isNaturalNumber(inputs[1]) {
				fmt.Println("The second parameter should be a natural number.")
				continue
			}
			_, ok := propertiesMap[inputs[2]]
			if !ok {
				fmt.Printf("The property [%s] is wrong.\n", inputs[2])
				properties := ""
				for k := range propertiesMap {
					properties += k + ", "
				}
				properties = properties[:len(properties)-2]
				fmt.Printf("Available properties: %v\n", properties)

				continue
			}
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
