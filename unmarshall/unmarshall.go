package unmarshall

import (
	"encoding/json"
	"log"

	"gopkg.in/yaml.v2"
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

// Unmarshals a string to the provided type 'V'
func YamlUnmarshall[T any](message string) (T, error) {
	var msg T
	err := yaml.Unmarshal([]byte(message), &msg)
	if err != nil {
		log.Println("Invalid message:" + err.Error())
		log.Println("jsonUnmarshall error:" + message)
	}
	return msg, err
}
