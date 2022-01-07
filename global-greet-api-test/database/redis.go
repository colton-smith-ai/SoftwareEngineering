package database

import (
	"encoding/json"
	"os"

	"github.com/go-redis/redis"
)

/*
	Reference for Go + Redis
	https://tutorialedge.net/golang/go-redis-tutorial/
*/

// Greeting database object (record structure)
type Greeting struct {
	Language string `json:"language"`
	Hello    string `json:"hello"`
	Welcome  string `json:"welcome"`
}

// redisClient package var connecting to redis
var redisClient *redis.Client

// GetDbClient getter function for database connector
func GetDbClient() *redis.Client {
	return redisClient
}

// EstablishRedis initiates connection to redis and verifies connection
func EstablishRedis() (string, error) {
	// Instantiate connection to redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_CONTAINER_NAME") + ":" + os.Getenv("DB_PORT"),
		Password: "",
		DB:       0,
	})

	// Verify redis connection
	pong, err := redisClient.Ping().Result()

	// Return error if unable to connect
	if err != nil {
		return "FAIL: UNABLE TO CONNECT TO DATABASE", err
	} else {
		// Insert default language into database
		errMessage, err := insertDefaultGreeting()
		// Catch insert errors
		if err != nil {
			return errMessage, err
		}

		// Return successfully connection message
		return "PASS: CONNECTED TO DATABASE ~ " + pong, nil
	}
}

// insertDefaultGreeting first entry in database upon startup
func insertDefaultGreeting() (string, error) {
	// Create greeting as json
	jsonGreeting, err := json.Marshal(Greeting{"English", "Hello", "Welcome"})

	// Check marshalling error
	if err != nil {
		return "FAIL: UNABLE TO MARSHALL JSON GREETING", err
	}

	// Insert json into database
	/*
		redis.client.Set() method parameters take a key, value, and key expiration.
		Setting the expiration to 0 effectively sets the key to have no expiration time.
	*/
	err = redisClient.Set("English", jsonGreeting, 0).Err()

	// Catch errors from database insert
	if err != nil {
		return "FAIL: UNABLE TO INSERT INTO DATABASE", err
	} else {
		return "PASS: INSERTED NEW ENTRY INTO DATABASE", err
	}
}
