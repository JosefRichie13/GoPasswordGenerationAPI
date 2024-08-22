package main

import (
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

// Checks if a string contains only words
// If yes, returns true
// Else returns false
func wordChecker(wordToCheck string) bool {
	for _, r := range wordToCheck {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Checks if a string is a number or not
// If its a number, returns true
// Else returns false
func checkForNumber(stringToCheck string) bool {
	_, err := strconv.Atoi(stringToCheck)
	if err != nil {
		return false
	} else {
		return true
	}
}

// Converts a string to a integer and returns it
func convertStringToNum(stringToConvert string) int {
	convertedNum, err := strconv.Atoi(stringToConvert)
	if err != nil {
		panic(err)
	}
	return convertedNum
}

// This function returns a random string of Captial letters
func getRandomCapitalLetters(numberOf string) string {

	// Converts the param to integer
	numberOfChar := convertStringToNum(numberOf)

	// All the capital letters in a slice
	allCapitalLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	// Create an empty slice to hold the result
	returningCapitalLetters := []string{}

	// Loop, gets a random element from the allCapitalLetters slice and appends it to the empty returningCapitalLetters slice
	for index := 0; index < numberOfChar; index++ {
		randomLetter := rand.Intn(len(allCapitalLetters))
		returningCapitalLetters = append(returningCapitalLetters, allCapitalLetters[randomLetter])
	}

	// Converts the returningCapitalLetters slice to a string and returns it.
	return strings.Join(returningCapitalLetters, "")
}

// Full explanation in getRandomCapitalLetters(), instead of Capital letters it returns random Small letters.
func getRandomSmallLetters(numberOf string) string {

	numberOfChar := convertStringToNum(numberOf)

	allSmallLetters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	returningSmallLetters := []string{}
	for index := 0; index < numberOfChar; index++ {
		randomLetter := rand.Intn(len(allSmallLetters))
		returningSmallLetters = append(returningSmallLetters, allSmallLetters[randomLetter])
	}
	return strings.Join(returningSmallLetters, "")
}

// Full explanation in getRandomCapitalLetters(), instead of Capital letters it returns random Numbers.
func getRandomNumbers(numberOf string) string {

	numberOfChar := convertStringToNum(numberOf)

	allNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	returningNumbers := []string{}
	for index := 0; index < numberOfChar; index++ {
		randomLetter := rand.Intn(len(allNumbers))
		returningNumbers = append(returningNumbers, allNumbers[randomLetter])
	}
	return strings.Join(returningNumbers, "")
}

// Full explanation in getRandomCapitalLetters(), instead of Capital letters it returns random Special Characters.
func getRandomSpecialChar(numberOf string) string {

	numberOfChar := convertStringToNum(numberOf)

	allSpecialChar := []string{"!", "@", "#", "$", "%", "^", "*", "(", ")", "-", "=", "+"}
	returningSpecialChar := []string{}

	for index := 0; index < numberOfChar; index++ {
		randomLetter := rand.Intn(len(allSpecialChar))
		returningSpecialChar = append(returningSpecialChar, allSpecialChar[randomLetter])
	}
	return strings.Join(returningSpecialChar, "")
}
