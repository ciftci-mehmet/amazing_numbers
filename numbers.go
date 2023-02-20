package main

import "strconv"

var propertiesMap = map[string]func(int) bool{
	"even":        isEvenNumber,
	"odd":         isOddNumber,
	"buzz":        isBuzzNumber,
	"duck":        isDuckNumber,
	"palindromic": isPalindromicNumber,
	"gapful":      isGapfulNumber,
	"spy":         isSpyNumber,
}

func isNaturalNumber(s string) bool {
	number, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if number < 1 {
		return false
	}

	return true
}

func isEvenNumber(n int) bool {
	return n%2 == 0
}

func isOddNumber(n int) bool {
	return n%2 == 1
}

func isBuzzNumber(n int) bool {
	if n%7 == 0 {
		return true
	}

	if n%10 == 7 {
		return true
	}

	return false
}

func isDuckNumber(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			return true
		}
	}

	return false
}

func isPalindromicNumber(n int) bool {
	s := strconv.Itoa(n)

	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}

	return true
}

func isGapfulNumber(n int) bool {
	if n < 100 {
		return false
	}

	s := strconv.Itoa(n)
	first, _ := strconv.Atoi(string(s[0]))
	last := n % 10
	concatenation := (first * 10) + last

	return n%concatenation == 0
}

func isSpyNumber(n int) bool {
	var sumOfDigits int
	for i := n; i > 0; i /= 10 {
		sumOfDigits += i % 10
	}

	productOfDigits := 1
	for i := n; i > 0; i /= 10 {
		productOfDigits *= i % 10
	}

	return sumOfDigits == productOfDigits
}

func getProperties(n int) string {
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
	if isSpyNumber(n) {
		s += "spy" + ", "
	}

	// trim last 2 chars ", " from the string
	s = s[:len(s)-2]
	return s
}

func getConsecutiveProperties(m, n int) string {
	var s string

	for i := m; i < m+n; i++ {
		s += "\t\t" + getProperties(i) + "\n"
	}

	// trim last newline char from the string
	s = s[:len(s)-1]
	return s
}

func getNumbersWithProperty(m, n int, property string) string{
	var s string
	found := 0

	numberFunction := propertiesMap[property]

	for i := m; found < n; i++ {
		if numberFunction(i) {
			s += "\t\t" + getProperties(i) + "\n"
			found++
		}
	}

	// trim last newline char from the string
	s = s[:len(s)-1]
	return s
}
