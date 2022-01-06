/*
	Author : Colton Smith
	Title  : Data Science Engineer
	Date   : December 2021
	Email  : colton.smith.ai@gmail.com
	GitHub : https://github.com/colton-smith-ai
*/

package main

import (
	"bytes"
	"encoding/json"
	"global-greet-client/apiconsumer"
	"global-greet-client/simpleio"
	"global-greet-client/user"
	"io"
	"net/http"
	"strings"
)

func main() {
	for true {
		// endless loop :?
		entrypoint()
	}
}

func entrypoint() {
	// get user input and instantiate user struct
	client := user.MakeUser()

	// get list of languages
	availableLang := getRequest(
		apiconsumer.RequestBuilder("/greeting"))

	// check if language is available
	if strings.Contains(string(availableLang), client.Native) {

		// retrieve api text
		apiGreeting := getRequest(
			apiconsumer.RequestBuilder("/greeting/" + client.Native))

		// greet user in their native language
		simpleio.Printer(client.Greet(apiGreeting))

	} else {
		// notify user their language is unknown
		simpleio.Printer("We do not have that language in our Database :/")

		// default greeting to English
		simpleio.Printer("Defaulting to English ...")
		apiGreeting := getRequest(
			apiconsumer.RequestBuilder("/greeting/English"))
		simpleio.Printer(client.Greet(apiGreeting))

		// ask user to update database with their language
		simpleio.Printer("Please help us by adding a new greeting for your language: " + client.Native)
		input := simpleio.UserInput("Will you help us? Enter either 'yes' or 'no': ")

		// see if user will help
		if strings.Contains(input, "yes") {

			// ask user input and send post request to API server
			success := postRequest(client)

			// display response message from post request
			simpleio.Printer("Success! New language in database, thank you!")
			simpleio.Printer("New language added: " + string(success))

			// retrieve api text
			apiGreeting := getRequest(
				apiconsumer.RequestBuilder("/greeting/" + client.Native))

			// greet user in their native language
			simpleio.Printer(client.Greet(apiGreeting))

		} else {
			simpleio.Printer("Goodbye, " + client.Name)
		}
	}
}

func getRequest(url string) []byte {
	// make request to server
	response, err := http.Get(url)

	// check if error in request
	if err != nil {
		simpleio.HandleErrors(err)
	}

	// close response body when finished
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			simpleio.HandleErrors(err)
		}
	}(response.Body)

	// extract body as byte array
	body, err := io.ReadAll(response.Body)

	// handle errors in readying body
	if err != nil {
		simpleio.HandleErrors(err)
	}

	// return body text
	return body
}

func postRequest(client user.User) []byte {
	// retrieve user's language greeting
	hola, bienvenido := userForm(client.Native)

	// marshal json body for post request
	requestBody, err := json.Marshal(map[string]string{
		"language": client.Native,
		"hello":    hola,
		"welcome":  bienvenido,
	})

	// check error in marshalling
	if err != nil {
		simpleio.HandleErrors(err)
	}

	// make post request to API to update language greeting
	resp, err := http.Post(
		apiconsumer.RequestBuilder("/greeting"),
		"application/json",
		bytes.NewBuffer(requestBody),
	)

	// check error in request
	if err != nil {
		simpleio.HandleErrors(err)
	}

	// close response body when finished
	defer resp.Body.Close()

	// extract body as byte array
	body, err := io.ReadAll(resp.Body)

	// handle errors in readying body
	if err != nil {
		simpleio.HandleErrors(err)
	}

	// return body text
	return body
}

func userForm(language string) (hello string, welcome string) {
	// ask for hello in users language
	hello = simpleio.UserInput("How do you say 'Hello' in " + language + ": ")
	// ask for welcome in users language
	welcome = simpleio.UserInput("How do you say 'Welcome' in " + language + ": ")
	return hello, welcome
}
