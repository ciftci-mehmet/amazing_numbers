package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Amazing Numbers!")
	fmt.Println()
	fmt.Println("Supported requests:")
	fmt.Println(" - enter a natural number to know its properties;")
	fmt.Println(" - enter 0 to exit.")


	var input string
	for{
		fmt.Println()
		fmt.Println("Enter a request:")
		fmt.Scanln(&input)
		fmt.Println()

		if input == "0" {
			fmt.Println("Goodbye!")
			return
		}
	
		if !isNaturalNumber(input){
			fmt.Println("The first parameter should be a natural number or zero.")
			continue
		}
		number, _ := strconv.Atoi(input)
	
		fmt.Println("Properties of", number)
		fmt.Printf("\teven: %v\n",isEvenNumber(number))
		fmt.Printf("\todd: %v\n",isOddNumber(number))
		fmt.Printf("\tbuzz: %v\n",isBuzzNumber(number))
		fmt.Printf("\tduck: %v\n",isDuckNumber(number))
		fmt.Printf("\tpalindromic: %v\n",isPalindromicNumber(number))
	}

}

func isNaturalNumber(s string) bool{
	number, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if number < 1 {
		return false
	}

	return true
}

func isEvenNumber(n int) bool{
	return n % 2 == 0
}

func isOddNumber(n int) bool{
	return n %2 == 1
}

func isBuzzNumber(n int) bool{
	if n % 7 == 0{
		return true
	}

	if n % 10 == 7{
		return true
	}
	
	return false
}

func isDuckNumber(n int) bool{
	s := strconv.Itoa(n)

	for i := 0; i < len(s); i++ {
		if s[i] == '0'{
			return true
		}
	}

	return false
}

func isPalindromicNumber(n int) bool{
	s := strconv.Itoa(n)

	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}

	return true
}

