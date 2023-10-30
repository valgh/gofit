package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func UnmarshalYAML(yamlFile []byte, data interface{}) error {
	err_yaml := yaml.Unmarshal(yamlFile, data)
	return err_yaml
}

func marshalYAML(data interface{}) ([]byte, error) {
	yamlData, err_yaml := yaml.Marshal(data)

	if err_yaml != nil {
		log.Fatalf("Error while marshaling object into YAML file: %v", err_yaml)
	}
	return yamlData, err_yaml
}

func GenerateYamlFile(data interface{}, fileName string) {

	yamlData, err := marshalYAML(data)

	if err != nil {
		return
	}

	fmt.Println("\n\n=========== YAML generated ==========")
	fmt.Println(string(yamlData))

	yamlFile, err_file := os.Create(fileName)

	if err_file != nil {
		log.Fatalf("Error while creating the config file: %v", err_file)
		return
	}

	_, err_write := yamlFile.Write(yamlData)

	if err_write != nil {
		log.Fatalf("Error while writing to config file: %v", err_write)
		return
	}

	fmt.Printf("\nYAML config file generated successfully!")
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return true
}
