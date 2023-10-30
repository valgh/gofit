package utils

import (
	"gofit/model"
	"regexp"
	"testing"
)

var tdTest = model.NewWorkoutPlan(3, "gain muscle", "strength", "Train for 45 minutes")

var bbTest = model.NewBodyBuilder(14, 157, 54, "Male", tdTest)

const yamlOut = `age: 14
height: 157
weight: 54
gender: Male
trainingdetails:
  days: 3
  goal: gain muscle
  workouttype: strength
  additionalinfo: Train for 45 minutes
`

// Test MarshalYAML function
func TestMarshalYAML(t *testing.T) {
	var yamlData, err = marshalYAML(bbTest)
	var yamlString = string(yamlData)
	want := regexp.MustCompile(yamlOut)
	if !want.MatchString(yamlString) || err != nil {
		t.Fatalf(`Marshal = %q, %v, want match for %#q, nil`, yamlString, err, want)
	}
}
