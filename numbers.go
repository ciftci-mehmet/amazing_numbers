package main

import (
	"fmt"
	"math"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var propertiesMap = map[string]func(int) bool{
	"even":        isEvenNumber,
	"odd":         isOddNumber,
	"buzz":        isBuzzNumber,
	"duck":        isDuckNumber,
	"palindromic": isPalindromicNumber,
	"gapful":      isGapfulNumber,
	"spy":         isSpyNumber,
	"square":      isSquareNumber,
	"sunny":       isSunnyNumber,
	"jumping":     isJumpingNumber,
	"happy":       isHappyNumber,
	"sad":         isSadNumber,
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

func isSquareNumber(n int) bool {
	x := math.Sqrt(float64(n))

	return x == float64(int64(x))
}

func isSunnyNumber(n int) bool {
	return isSquareNumber(n + 1)
}

func isJumpingNumber(n int) bool {
	var i int
	for i = n; i > 9; i /= 10 {
		if math.Abs(float64((i%10)-((i/10)%10))) != 1 {
			return false
		}
	}
	return true
}

func isHappyNumber(n int) bool {
	var sum, digit int
	lastNumber := n
	checkedNumbers := make(map[int]bool)
	for {
		if _, ok := checkedNumbers[lastNumber]; ok {
			return false
		}
		if sum == 1 {
			return true
		}
		sum = 0
		for i := lastNumber; i > 0; i /= 10 {
			digit = i % 10
			sum += digit * digit
		}
		checkedNumbers[lastNumber] = true
		lastNumber = sum
	}
}

func isSadNumber(n int) bool {
	return !isHappyNumber(n)
}

func getProperties(n int) string {
	s := formatNumber(n)
	s += " is "

	for k, f := range propertiesMap {
		if f(n) {
			s += k + ", "
		}
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

func getNumbersWithMultiProperties(m, n int, properties []string) string {
	var s string
	found := 0

	for i := m; found < n; i++ { // loop until found enough numbers
		satisfyConditions := true

		for j := 0; j < len(properties); j++ { // loop filtered properties

			if properties[j][0] == '-' {
				numberFunc := propertiesMap[properties[j][1:len(properties[j])]]
				if numberFunc(i) {
					satisfyConditions = false
					break
				}
			} else {
				numberFunc := propertiesMap[properties[j]]
				if !numberFunc(i) {
					satisfyConditions = false
					break
				}
			}

		}

		if satisfyConditions {
			s += "\t\t" + getProperties(i) + "\n"
			found++
		}
	}

	// trim last newline char from the string
	s = s[:len(s)-1]
	return s
}

func getSingleNumberProperties(n int) string {
	var s string

	s += fmt.Sprintf("Properties of %v\n", formatNumber(n))

	for k, f := range propertiesMap {
		s += fmt.Sprintf("\t%v: %v\n", k, f(n))
	}

	// trim last newline char from the string
	s = s[:len(s)-1]
	return s
}

func getAllProperties() string {
	properties := ""
	for k := range propertiesMap {
		properties += k + ", "
	}
	properties = properties[:len(properties)-2]
	return properties
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func checkMutualExlusivity(inputProperties []string) (bool, string) {
	if contains(inputProperties, "odd") && contains(inputProperties, "even") {
		return true, "[ODD, EVEN]"
	}
	if contains(inputProperties, "-odd") && contains(inputProperties, "-even") {
		return true, "[-ODD, -EVEN]"
	}
	if contains(inputProperties, "-odd") && contains(inputProperties, "odd") {
		return true, "[-ODD, ODD]"
	}
	if contains(inputProperties, "-even") && contains(inputProperties, "even") {
		return true, "[-EVEN, EVEN]"
	}

	if contains(inputProperties, "square") && contains(inputProperties, "sunny") {
		return true, "[SQUARE, SUNNY]"
	}
	if contains(inputProperties, "-square") && contains(inputProperties, "square") {
		return true, "[-SQUARE, SQUARE]"
	}
	if contains(inputProperties, "-sunny") && contains(inputProperties, "sunny") {
		return true, "[-SUNNY, SUNNY]"
	}

	if contains(inputProperties, "duck") && contains(inputProperties, "spy") {
		return true, "[DUCK, SPY]"
	}
	if contains(inputProperties, "-duck") && contains(inputProperties, "duck") {
		return true, "[-DUCK, DUCK]"
	}
	if contains(inputProperties, "-spy") && contains(inputProperties, "spy") {
		return true, "[-SPY, SPY]"
	}

	if contains(inputProperties, "happy") && contains(inputProperties, "sad") {
		return true, "[HAPPY, SAD]"
	}
	if contains(inputProperties, "-happy") && contains(inputProperties, "-sad") {
		return true, "[-HAPPY, -SAD]"
	}
	if contains(inputProperties, "-happy") && contains(inputProperties, "happy") {
		return true, "[-HAPPY, HAPPY]"
	}
	if contains(inputProperties, "-sad") && contains(inputProperties, "sad") {
		return true, "[-SAD, SAD]"
	}
	return false, ""
}

func formatNumber(n int) string {
	np := message.NewPrinter(language.English)
	s := np.Sprintf("%d", n)
	return s
}
