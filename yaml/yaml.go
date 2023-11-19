package yaml

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

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

// Marshals a struct of type 'V' to a yaml string
func YamlMarshall[T any](message T) (string, error) {
	// Convert message into yaml
	b, err := yaml.Marshal(message)
	if err != nil {
		fmt.Println("Could not covert request to yaml config data")
		log.Printf("yamlMarshall error: %v\n", message)
	}

	return string(b), err
}
