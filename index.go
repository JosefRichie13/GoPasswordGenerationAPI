package main

import (
	_ "fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func main() {

	request := gin.Default()
	request.GET("/", landingPage)
	request.GET("/generateAPassword", generateAPassword)
	request.Run(":8083")

}

func landingPage(c *gin.Context) {
	c.JSON(200, "Hello, Welcome to Password Generation API")
}

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
