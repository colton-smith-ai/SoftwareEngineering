package endpoints

import (
	"encoding/json"
	db "global-greet-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGreeting responds with all languages
func GetGreeting(c *gin.Context) {
	// Ping database to verify heartbeat/connection
	_, err := db.GetDbClient().Ping().Result()

	// Response with error: cannot connect to database
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error_message": "database connection failed"})
		return
	}

	// Retrieve list of database keys (available languages)
	languages, err := db.GetDbClient().Keys("*").Result()

	// Response with error message, if applicable
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error_message": "no records in database"})
		return
	}

	// Response with languages in database
	c.IndentedJSON(http.StatusOK, languages)
}

// PostGreeting adds greeting from JSON received in the request body
func PostGreeting(c *gin.Context) {
	// Ping database to verify heartbeat/connection
	_, err := db.GetDbClient().Ping().Result()

	// Response with error: cannot connect to database
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error_message": "database connection failed"})
		return
	}

	// Create greeting to parse user request
	var newGreet db.Greeting

	// Call BindJSON to bind the received JSON to newGreet
	if err := c.BindJSON(&newGreet); err != nil {
		return
	}

	// Marshall request into json, to insert into database
	jsonGreet, err := json.Marshal(newGreet)

	// Check for json marshalling errors
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error_message": "failed to marshall json from request"})
		return
	}

	// Insert json into database
	err = db.GetDbClient().Set(newGreet.Language, jsonGreet, 0).Err()

	// Check for insertion error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error_message": "failed to insert record into database"})
		return
	}

	// Response with inserted greeting
	c.IndentedJSON(http.StatusCreated, newGreet)
}

// GetGreetingByLanguage locates greeting whose language value matches the parameter
// sent by the client, then returns that greeting as a response
func GetGreetingByLanguage(c *gin.Context) {
	// Ping database to verify heartbeat/connection
	_, err := db.GetDbClient().Ping().Result()

	// Response with error: cannot connect to database
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error_message": "database connection failed"})
		return
	}

	// Retrieve language from request
	lang := c.Param("language")

	// Extract greeting from database
	greetGuest, err := db.GetDbClient().Get(lang).Result()

	// Verify if language in database
	if err != nil {
		// Response 404 with message
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "language not available :/"})
		return
	}

	// Convert json string from database into Greeting obj
	greetingResponse := db.Greeting{}
	err = json.Unmarshal([]byte(greetGuest), &greetingResponse)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error_message": "failed unmarshall json from database"})
		return
	}

	// Language exists, return greeting
	c.IndentedJSON(http.StatusOK, greetingResponse)
}
