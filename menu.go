package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printHelp() {
	helpText :=
		"Supported requests:\n" +
			" - enter a natural number to know its properties;\n" +
			" - enter two natural numbers to obtain the properties of the list:\n" +
			"   * the first parameter represents a starting number;\n" +
			"   * the second parameter shows how many consecutive numbers are to be processed;\n" +
			" - two natural numbers and properties to search for;\n" +
			" - a property preceded by minus must not be present in numbers;\n" +
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

func isInputsValid(inputs []string) (bool, string) {
	switch inputLen := len(inputs); {
	case inputLen > 2:
		inputProperties := inputs[2:]
		ok, exlusivePairs := checkMutualExlusivity(inputProperties)
		if ok {
			message := "The request contains mutually exclusive properties: " + exlusivePairs + "\n"
			message += "There are no numbers with these properties."
			return false, message
		}

		ok, message := checkProperties(inputProperties)
		if !ok {
			return false, message
		}
		fallthrough

	case inputLen == 2:
		if !isNaturalNumber(inputs[1]) {
			return false, "The second parameter should be a natural number."
		}
		fallthrough

	case inputLen == 1:
		if !isNaturalNumber(inputs[0]) {
			return false, "The first parameter should be a natural number or zero."
		}
	}
	return true, ""
}

func checkProperties(ss []string) (bool, string) {
	var wrongProperties []string
	var message, property string
	for _, s := range ss {
		if s[0] == '-' {
			property = s[1:]
		} else {
			property = s
		}
		_, ok := propertiesMap[property]
		if !ok {
			wrongProperties = append(wrongProperties, strings.ToUpper(property))
		}
	}
	if len(wrongProperties) > 1 {
		message = fmt.Sprintf("The properties %v are wrong.\n", wrongProperties)
	}
	if len(wrongProperties) == 1 {
		message = fmt.Sprintf("The property %v is wrong.\n", wrongProperties)
	}
	if len(wrongProperties) > 0 {
		properties := getAllProperties()
		message += fmt.Sprintf("Available properties: %v", properties)
		return false, message
	}

	return true, ""
}
