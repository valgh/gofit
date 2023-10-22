package model

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ChatGPTConfig struct {
	Apikey    string
	Endpoint  string
	Model     string
	MaxTokens int
}

func NewChatGPTConfig(apikey string, endpoint string, model string, maxtokens int) ChatGPTConfig {
	return ChatGPTConfig{
		Apikey:    apikey,
		Endpoint:  endpoint,
		Model:     model,
		MaxTokens: maxtokens,
	}
}

func (config ChatGPTConfig) FromYaml(fileName string) ChatGPTConfig {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading "+fileName+" file: %v", err)
	}

	var gptConfig ChatGPTConfig

	err_yaml := yaml.Unmarshal(yamlFile, &gptConfig)
	if err_yaml != nil {
		log.Fatalf("Error while unmarshaling "+fileName+" file: %v", err_yaml)
	}

	return gptConfig
}
