package model

type WorkoutPlan struct {
	Days           int
	Goal           string
	WorkoutType    string
	AdditionalInfo string
}

func NewWorkoutPlan(days int, goal string, workoutType string, additionalInfo string) WorkoutPlan {
	return WorkoutPlan{
		Days:           days,
		Goal:           goal,
		WorkoutType:    workoutType,
		AdditionalInfo: additionalInfo,
	}
}

func (w WorkoutPlan) GetDays() int {
	return w.Days
}

func (w *WorkoutPlan) SetDays(days int) {
	w.Days = days
}

func (w WorkoutPlan) GetGoal() string {
	return w.Goal
}

func (w *WorkoutPlan) SetGoal(goal string) {
	w.Goal = goal
}

func (w WorkoutPlan) GetWorkoutType() string {
	return w.WorkoutType
}

func (w *WorkoutPlan) SetWorkoutType(workoutType string) {
	w.WorkoutType = workoutType
}

func (w WorkoutPlan) GetAdditionalInfo() string {
	return w.AdditionalInfo
}

func (w *WorkoutPlan) SetAdditionalInfo(additionalInfo string) {
	w.AdditionalInfo = additionalInfo
}
