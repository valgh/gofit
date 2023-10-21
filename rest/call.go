package rest

import (
	"encoding/json"
	"gofit/model"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v2"
)

const configFile = "gptConfig.yaml"

var defaultConfig model.ChatGPTConfig = model.ChatGPTConfig{
	Apikey:    "none",
	Endpoint:  "https://api.openai.com/v1/chat/completions",
	Model:     "gpt-3.5-turbo",
	MaxTokens: 200,
}

func getConfigFromYaml(fileName string) model.ChatGPTConfig {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error: config file not found %v", err)
		return defaultConfig
	}

	var chatConfig model.ChatGPTConfig = model.ChatGPTConfig{}.FromYaml(fileName)

	err_yaml := yaml.Unmarshal(yamlFile, &chatConfig)
	if err_yaml != nil {
		log.Fatalf("Error while unmarshaling "+configFile+" file: %v", err_yaml)
		return defaultConfig
	}

	return chatConfig
}

func GPTRestCall(message string) interface{} {

	chatGPTConfig := getConfigFromYaml(configFile)
	restClient := resty.New()

	resp, err := restClient.R().SetAuthToken(chatGPTConfig.Apikey).
		SetHeader("Content-Type", "application/json").
		SetBody(
			map[string]interface{}{
				"model": chatGPTConfig.Model,
				"messages": []interface{}{
					map[string]interface{}{
						"role":    "system",
						"content": message,
					},
				},
				"max_tokens": chatGPTConfig.MaxTokens,
			},
		).Post(chatGPTConfig.Endpoint)

	if err != nil {
		log.Fatalf("Error while forwarding request to ChatGPT endpoint: %v", err)
	}

	body := resp.Body()
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error while unmarshaling JSON response: %v", err)
	}

	var content string
	if data["error"] != nil {
		content = data["error"].(map[string]interface{})["message"].(string)
	} else {
		content = data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	}
	return content
}
