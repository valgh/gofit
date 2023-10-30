package builder

import (
	"gofit/model"
	"gofit/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

var message = `Hello ChatGPT! Can you please give me a workout plan?
I'm a XXX years-old XXX, my height is XXX cm and my weight is XXX kg.
I want to work out XXX times a week.
My goal is XXX and I would like a workout type XXX. Some additional information: XXX`

const configFile = "workoutConfig.yaml"
const PLACEHOLDER = "XXX"

func BuildMessageFromFile(filename string) string {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading "+filename+" file: %v", err)
	}

	var bodyBuilder model.BodyBuilder

	err_yaml := utils.UnmarshalYAML(yamlFile, &bodyBuilder)
	if err_yaml != nil {
		log.Fatalf("Error while unmarshaling "+filename+" file: %v", err_yaml)
	}

	message = BuildMessage(bodyBuilder)
	return message
}

func BuildMessage(bodyBuilder model.BodyBuilder) string {
	message = strings.Replace(message, PLACEHOLDER, strconv.Itoa(bodyBuilder.GetAge()), 1)
	message = strings.Replace(message, PLACEHOLDER, bodyBuilder.GetGender(), 1)
	message = strings.Replace(message, PLACEHOLDER, strconv.Itoa(bodyBuilder.GetHeight()), 1)
	message = strings.Replace(message, PLACEHOLDER, strconv.Itoa(bodyBuilder.GetWeight()), 1)
	trainingDetails := bodyBuilder.GetTrainingDetails()
	message = strings.Replace(message, PLACEHOLDER, strconv.Itoa(trainingDetails.GetDays()), 1)
	message = strings.Replace(message, PLACEHOLDER, trainingDetails.GetGoal(), 1)
	message = strings.Replace(message, PLACEHOLDER, trainingDetails.GetWorkoutType(), 1)
	message = strings.Replace(message, PLACEHOLDER, trainingDetails.GetAdditionalInfo(), 1)
	return message
}
