package builder

import (
	"gofit/model"
	"testing"
)

var tdTest = model.NewWorkoutPlan(3, "gain muscle", "strength", "Train for 45 minutes")
var bbTest = model.NewBodyBuilder(14, 157, 54, "Male", tdTest)

var expected = `Hello ChatGPT! Can you please give me a workout plan?
I'm a 14 years-old Male, my height is 157 cm and my weight is 54 kg.
I want to work out 3 times a week.
My goal is gain muscle and I would like a workout type strength. Some additional information: Train for 45 minutes`

func TestMessageBuilder(t *testing.T) {
	outMessage := BuildMessage(bbTest)
	if outMessage != expected {
		t.Fatalf(`Output = %q, want match for %#q`, outMessage, expected)
	}
}

func TestMessageBuilderFromFile(t *testing.T) {
	outMessage := BuildMessageFromFile("test_res/workoutConfig_test.yaml")
	if outMessage != expected {
		t.Fatalf(`Output = %q, want match for %#q`, outMessage, expected)
	}
}
