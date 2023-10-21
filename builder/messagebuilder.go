package builder

import (
	"gofit/model"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

var message = `Hello ChatGPT! Can you please give me a workout plan?
I'm a XXX years-old XXX, my height is XXX cm and my weight is XXX kg.
My goal is XXX and I would like a workout type XXX.`

const configFile = "workoutConfig.yaml"
const PLACEHOLDER = "XXX"

func BuildMessage(filename string) string {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading "+filename+" file: %v", err)
	}

	var bodyBuilder model.BodyBuilder

	err_yaml := yaml.Unmarshal(yamlFile, &bodyBuilder)
	if err_yaml != nil {
		log.Fatalf("Error while unmarshaling "+filename+" file: %v", err_yaml)
	}

	message = strings.Replace(message, PLACEHOLDER, strconv.Itoa(bodyBuilder.GetAge()), 1)
	message = strings.Replace(message, PLACEHOLDER, bodyBuilder.GetGender(), 1)
	message = strings.Replace(message, PLACEHOLDER, strconv.Itoa(bodyBuilder.GetHeight()), 1)
	message = strings.Replace(message, PLACEHOLDER, strconv.Itoa(bodyBuilder.GetWeight()), 1)
	message = strings.Replace(message, PLACEHOLDER, bodyBuilder.GetTrainingDetails().GetGoal(), 1)
	message = strings.Replace(message, PLACEHOLDER, bodyBuilder.TrainingDetails.GetWorkoutType(), 1)
	return message
}
