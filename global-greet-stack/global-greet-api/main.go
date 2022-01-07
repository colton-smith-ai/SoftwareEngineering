/*
	Author : Colton Smith
	Email  : colton.smith.ai@gmail.com
	Github : https://github.com/colton-smith-ai
	Date   : January 2022
*/

package main

import (
	"fmt"
	db "global-greet-api/database"
	ep "global-greet-api/endpoints"
	"net/http"
	"os"

	"github.com/colton-smith-ai/SoftwareEngineering/watermark"
	"github.com/gin-gonic/gin"
)

func main() {

	// @colton.smith.ai
	watermark.Sign("January 2022")

	// Check connection to database
	errMessage, err := db.EstablishRedis()
	if err != nil {
		fmt.Println(errMessage)
		fmt.Println(err)
		return
	} else {
		fmt.Println(errMessage)
	}

	// Route API endpoints
	router := gin.Default()
	router.GET("/", getAbout)
	router.GET("/greeting", ep.GetGreeting)
	router.GET("/greeting/:language", ep.GetGreetingByLanguage)
	router.POST("/greeting", ep.PostGreeting)
	router.GET("/math", ep.GetMath)
	router.PUT("/math/:x/:y", ep.PutMath)

	// Expose server port
	router.Run(":" + os.Getenv("API_PORT"))
}

// Home page with readme and help to use API calls
func getAbout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"about": "Welcome to this API about greetings in various languages.",

		"help": "This API serves 3 endpoints: " +
			"1) GET /greeting = returns list of available languages to query API " +
			"2) POST /greeting = adds a new greeting from JSON received in the request body " +
			"3) GET /greeting/<language> = returns greeting for specified language",

		"easterEgg": "There is a hidden endpoint that takes two numbers and adds them together. " +
			"See if you can solve this MATH problem ;)",
	})
}
