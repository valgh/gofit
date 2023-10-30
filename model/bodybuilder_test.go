package model

import (
	"testing"
)

var tdTest = NewWorkoutPlan(3, "better shape", "cardio", "I want to exercise 60 minutes each day")
var expected = NewBodyBuilder(20, 180, 80, "Male", tdTest)

func TestFromYAML(t *testing.T) {
	parsedBB := expected.FromYaml("test_res/workoutConfig_test.yaml")
	if parsedBB != expected {
		t.Fatalf(`Output = %q, want match for %#q`, parsedBB, expected)
	}
}
