package json

import (
	"encoding/json"
	"log"
)

// Unmarshals a string to the provided type 'V'
func JsonUnmarshall[T any](message string) (T, error) {
	var msg T
	err := json.Unmarshal([]byte(message), &msg)
	if err != nil {
		log.Println("Invalid message:" + err.Error())
		log.Println("jsonUnmarshall error:" + message)
	}
	return msg, err
}
