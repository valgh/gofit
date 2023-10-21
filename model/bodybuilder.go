package model

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type BodyBuilder struct {
	Age             int
	Height          int
	Weight          int
	Gender          string
	TrainingDetails WorkoutPlan
}

func NewBodyBuilder(age int, height int, weight int, gender string, trainingDetails WorkoutPlan) BodyBuilder {
	return BodyBuilder{
		Age:             age,
		Height:          height,
		Weight:          weight,
		Gender:          gender,
		TrainingDetails: trainingDetails,
	}
}

func (b BodyBuilder) GetAge() int {
	return b.Age
}

func (b *BodyBuilder) SetAge(age int) {
	b.Age = age
}

func (b BodyBuilder) GetHeight() int {
	return b.Height
}

func (b *BodyBuilder) SetHeight(height int) {
	b.Height = height
}

func (b BodyBuilder) GetWeight() int {
	return b.Weight
}

func (b *BodyBuilder) SetWeight(weight int) {
	b.Weight = weight
}

func (b BodyBuilder) GetGender() string {
	return b.Gender
}

func (b *BodyBuilder) SetGender(gender string) {
	b.Gender = gender
}

func (b BodyBuilder) GetTrainingDetails() WorkoutPlan {
	return b.TrainingDetails
}

func (b *BodyBuilder) SetTrainingDetails(trainingDetails WorkoutPlan) {
	b.TrainingDetails = trainingDetails
}

func (b BodyBuilder) FromYaml(fileName string) BodyBuilder {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error reading "+fileName+" file: %v", err)
	}

	var bodyBuilder BodyBuilder

	err_yaml := yaml.Unmarshal(yamlFile, &bodyBuilder)
	if err_yaml != nil {
		log.Fatalf("Error while unmarshaling "+fileName+" file: %v", err_yaml)
	}

	return bodyBuilder
}
