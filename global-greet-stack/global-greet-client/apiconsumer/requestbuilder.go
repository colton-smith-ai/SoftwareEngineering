package apiconsumer

import "os"

const (
	protocol        = "http://"
	socketConnector = ":"
)

func RequestBuilder(requestParameters string) string {
	domain := os.Getenv("API_CONTAINER_NAME")
	port := os.Getenv("API_CONTAINER_PORT")
	return protocol + domain + socketConnector + port + requestParameters
}
