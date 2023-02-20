package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Amazing Numbers!")
	fmt.Println()
	helpText := 
		"Supported requests:\n" +
		" - enter a natural number to know its properties;\n" +
		" - enter two natural numbers to obtain the properties of the list:\n" +
		"   * the first parameter represents a starting number;\n" +
		"   * the second parameter shows how many consecutive numbers are to be processed;\n" +
		" - separate the parameters with one space\n" +
		" - enter 0 to exit."
	fmt.Println(helpText)
	
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Println()
		fmt.Println("Enter a request:")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println()

		if input == "0" {
			fmt.Println("Goodbye!")
			return
		}

		if input == "" {
			fmt.Println(helpText)
			continue
		}

		// multi number input check
		numbers := strings.Split(input, " ")
		if len(numbers) == 2 && isNaturalNumber(numbers[0]) && isNaturalNumber(numbers[1]){
			number1, _ := strconv.Atoi(numbers[0])
			number2, _ := strconv.Atoi(numbers[1])
			fmt.Println(getConsecutiveProperties(number1, number2))
			continue
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

func isGapfulNumber(n int) bool{
	if n < 100{
		return false
	}

	s := strconv.Itoa(n)
	first, _ := strconv.Atoi(string(s[0]) )
	last := n%10
	concatenation := (first * 10) + last

	return n % concatenation == 0
}

func getProperties(n int) string{
	s := strconv.Itoa(n)
	s += " is "

	if isEvenNumber(n) {
		s += "even" + ", "
	}
	if isOddNumber(n) {
		s += "odd" + ", "
	}
	if isDuckNumber(n) {
		s += "duck" + ", "
	}
	if isBuzzNumber(n) {
		s += "buzz" + ", "
	}
	if isPalindromicNumber(n) {
		s += "palindromic" + ", "
	}
	if isGapfulNumber(n) {
		s += "gapful" + ", "
	}

	// trim last 2 chars ", " from the string
	s = s[:len(s)-2]
	return s
}

func getConsecutiveProperties(m, n int) string{
	var s string

	for i := m; i < m+n; i++ {
		s += "\t\t" + getProperties(i) + "\n"
	}

	// trim last newline char from the string
	s = s[:len(s)-1]
	return s
}
