package json

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/katasec/utils/yaml"
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

func JsonToYaml[T any](message string) (string, error) {

	myStruct, _ := JsonUnmarshall[T](message)
	yamlString, err := yaml.YamlMarshall(myStruct)
	if err != nil {
		fmt.Printf("There were errors, this message will not be processed. Message: %s\n", message)
	}

	return yamlString, err
}
