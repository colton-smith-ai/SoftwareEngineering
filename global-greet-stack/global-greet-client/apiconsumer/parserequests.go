package apiconsumer

import (
	"encoding/json"
	"global-greet-client/simpleio"
)

type GreetingResponse struct {
	Language string `json:"language"`
	Hello    string `json:"hello"`
	Welcome  string `json:"welcome"`
}

func ParseToJSON(request []byte) GreetingResponse {
	greet := GreetingResponse{}
	err := json.Unmarshal(request, &greet)
	if err != nil {
		simpleio.HandleErrors(err)
	}
	return greet
}
