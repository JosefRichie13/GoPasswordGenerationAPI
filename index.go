package main

import (
	_ "fmt"
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	request := gin.Default()
	request.GET("/", landingPage)
	request.GET("/generateAPassword", generateAPassword)
	request.GET("/generateLeetPassword", generateALeetPassword)
	request.Run(":8083")

}

func landingPage(c *gin.Context) {
	c.JSON(200, "Hello, Welcome to Password Generation API")
}

// Function to generate a password
func generateAPassword(c *gin.Context) {

	// We are specifying 4 query parameters, capLetters, smallLetters, numbers, specialChar from the API.
	noOfCapLetters := c.Query("capLetters")
	noOfSmallLetters := c.Query("smallLetters")
	noOfNumbers := c.Query("numbers")
	noOfSpecialChar := c.Query("specialChar")

	// Checking if any query param is empty. If any one is empty, reject with 400
	if noOfCapLetters == "" || noOfSmallLetters == "" || noOfNumbers == "" || noOfSpecialChar == "" {
		c.JSON(400, gin.H{"status": "Please provide all required parameters"})
		return
	}

	// Checking if any query param is equal to 0. If any one is 0, reject with 400
	if noOfCapLetters == "0" || noOfSmallLetters == "0" || noOfNumbers == "0" || noOfSpecialChar == "0" {
		c.JSON(400, gin.H{"status": "All parameters should be above 0"})
		return
	}

	// Checking if any query param are whole numbers. If any one is not, reject with 400
	if !checkForNumber(noOfCapLetters) || !checkForNumber(noOfSmallLetters) || !checkForNumber(noOfNumbers) || !checkForNumber(noOfSpecialChar) {
		c.JSON(400, gin.H{"status": "All required parameters should be whole numbers"})
		return
	}

	// Generate a string using 4 functions
	generatedPwd := getRandomCapitalLetters(noOfCapLetters) + getRandomSmallLetters(noOfSmallLetters) + getRandomNumbers(noOfNumbers) + getRandomSpecialChar(noOfSpecialChar)

	// Shuffle the generated string
	shuff := []rune(generatedPwd)
	rand.Shuffle(len(generatedPwd), func(i, j int) {
		shuff[i], shuff[j] = shuff[j], shuff[i]
	})

	// Return the shuffled string as a password
	c.JSON(200, gin.H{"password": string(shuff)})
}

// Defining the Query Param body using a struct. This will accept the '?word=<QWERTY>' query param from the URL
type WordToConvert struct {
	Word string `form:"word" binding:"required"`
}

// Function to generate a Leet Password
func generateALeetPassword(c *gin.Context) {

	// Defining a map of Leets
	leetCharacterMap := map[string]string{"A": "@", "B": "8", "C": "{", "D": "[)", "E": "3", "F": "|=", "G": "6", "H": "[-]", "I": "!", "J": "._|", "K": "|<", "L": "1",
		"M": "|V|", "N": "|^|", "O": "0", "P": "9", "Q": "(_,)", "R": "12", "S": "5", "T": "7", "U": "|_|", "V": "\\//", "W": "|_|_|",
		"X": "><", "Y": "Y", "Z": "2"}

	// Creating an instance of the struct, WordToConvert
	var wordToConvert WordToConvert

	// Bind to the struct's members. If any member is invalid, binding does not happen and an error will be returned. Then its rejected with 400
	if c.Bind(&wordToConvert) != nil {
		c.JSON(400, gin.H{"status": "Please provide all required parameters"})
		return
	}

	// Checking if the word is atleast 4 characters long. If not, reject with 400
	if len(wordToConvert.Word) < 4 {
		c.JSON(400, gin.H{"status": "Please provide atleast 4 letters"})
		return
	}

	// Checking if the word contains any numbers or special characters. If it contains any of it, reject with 400
	if !wordChecker(wordToConvert.Word) {
		c.JSON(400, gin.H{"status": "Please provide only Alphabets, no numbers, special characters or spaces"})
		return
	}

	// Capitalizing the word
	capitalizedWord := strings.ToUpper(wordToConvert.Word)

	// Splliting the word into a string array
	splitWord := strings.Split(capitalizedWord, "")

	// Defining an empty slice to hold the leeted characters
	leetMapSlice := []string{}

	// Loop through the string array and convert all the characters to leet, except the 1st and the last
	for point := 1; point < len(splitWord)-1; point++ {
		leetedCharMap := leetCharacterMap[splitWord[point]]
		leetMapSlice = append(leetMapSlice, leetedCharMap)
	}

	// Putting it altogether
	generatedLeetPassword := splitWord[0] + strings.Join(leetMapSlice, "") + splitWord[len(splitWord)-1]

	// Return the Leeted password
	c.JSON(200, gin.H{"inputPassword": wordToConvert.Word, "leetPassword": generatedLeetPassword})

}
