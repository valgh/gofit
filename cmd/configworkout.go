package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gofit/model"
	"gofit/utils"

	"github.com/spf13/cobra"
)

var defaultWorkoutPlan = model.NewWorkoutPlan(3, "", "", "")

var defaultBodyBuilder = model.NewBodyBuilder(20, 180, 75, "Male", defaultWorkoutPlan)

const workoutFileName = "workoutConfig.yaml"

var configworkoutCmd = &cobra.Command{
	Use:   "cfgwkt",
	Short: "Create a default workout config file or edit an existing one.",
	Long:  `Create a default workout config file, or edit an existing one.`,
	Run: func(cmd *cobra.Command, args []string) {
		isEdit, err := cmd.Flags().GetBool("edit")

		if err != nil {
			return
		}

		if isEdit == true && utils.FileExists(workoutFileName) {
			fmt.Printf("Config file exists. Entering edit mode...")
			generateYAMLFomUserInput(workoutFileName)
		} else if isEdit == false {
			if !utils.FileExists(workoutFileName) {
				bodyBuilder := generateFromCommand(cmd)
				if bodyBuilder != (model.BodyBuilder{}) {
					utils.GenerateYamlFile(bodyBuilder, workoutFileName)
				}
			} else {
				fmt.Printf("Your config file already exists. To modify it, please run 'gofit config -e' or 'gofit config --edit'")
			}
		} else {
			fmt.Printf("Please generate a config file. Run 'gofit configworkout' command.")
		}
	},
}

func generateFromCommand(cmd *cobra.Command) model.BodyBuilder {
	age, errAG := cmd.Flags().GetInt("age")
	height, errHG := cmd.Flags().GetInt("height")
	weight, errWG := cmd.Flags().GetInt("weight")
	gender, errGD := cmd.Flags().GetString("gender")
	days, errDA := cmd.Flags().GetInt("days")
	goal, errGL := cmd.Flags().GetString("goal")
	wkType, errWT := cmd.Flags().GetString("type")
	info, errAI := cmd.Flags().GetString("info")

	if errAG != nil || errHG != nil || errWG != nil ||
		errGD != nil || errDA != nil || errGL != nil ||
		errWT != nil || errAI != nil {
		log.Fatalf("Invalid or missing parameter. Please re-enter the command with valid input.")
		return model.BodyBuilder{}
	}

	workoutPlan := model.WorkoutPlan{
		Days:           days,
		Goal:           goal,
		WorkoutType:    wkType,
		AdditionalInfo: info,
	}

	bodyBuilder := model.BodyBuilder{
		Age:             age,
		Height:          height,
		Weight:          weight,
		Gender:          gender,
		TrainingDetails: workoutPlan,
	}

	return bodyBuilder
}

func generateYAMLFomUserInput(filename string) {
	var bodyBuilder model.BodyBuilder = model.BodyBuilder{}.FromYaml(filename)
	if (bodyBuilder == model.BodyBuilder{}) {
		log.Fatalf("YAML file corrupted, or malformed.")
		return
	}
	editStructFromUserInput(&bodyBuilder)
	utils.GenerateYamlFile(bodyBuilder, workoutFileName)
}

func editStructFromUserInput(bodyBuilder *model.BodyBuilder) {
	var numberInput int
	var charInput string

	fmt.Printf("Enter your Age (current value: %v):  ", bodyBuilder.GetAge())
	fmt.Scanf("%d\n", &numberInput)
	bodyBuilder.SetAge(numberInput)

	fmt.Printf("Enter your Height (current value: %v):  ", bodyBuilder.GetHeight())
	fmt.Scanf("%d\n", &numberInput)
	bodyBuilder.SetHeight(numberInput)

	fmt.Printf("Enter your Weight (current value: %v):  ", bodyBuilder.GetWeight())
	fmt.Scanf("%d\n", &numberInput)
	bodyBuilder.SetWeight(numberInput)

	fmt.Printf("Enter your Gender (current value: %v):  ", bodyBuilder.GetGender())
	fmt.Scanf("%v\n", &charInput)
	bodyBuilder.SetGender(charInput)

	trainingDetails := bodyBuilder.GetTrainingDetails()

	fmt.Printf("Enter how many days per week you are going to workout (current value: %v):  ", trainingDetails.GetDays())
	fmt.Scanf("%d\n", &numberInput)
	trainingDetails.SetDays(numberInput)

	fmt.Printf("Enter your goal for this workout (current value: %v):  ", trainingDetails.GetGoal())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	goal := scanner.Text()
	trainingDetails.SetGoal(goal)

	fmt.Printf("Enter what kind of workout you wish - cardio, strength... (current value: %v):  ", trainingDetails.GetWorkoutType())
	scanner.Scan()
	wkType := scanner.Text()
	trainingDetails.SetWorkoutType(wkType)

	fmt.Printf("Enter any additional information to be requested for your workout (current value %v): ", trainingDetails.GetAdditionalInfo())
	scanner.Scan()
	additionalInfo := scanner.Text()
	trainingDetails.SetAdditionalInfo(additionalInfo)

	bodyBuilder.SetTrainingDetails(trainingDetails)

}

func init() {
	configworkoutCmd.Flags().IntP("age", "a", 20, "User Age.")
	configworkoutCmd.Flags().IntP("height", "c", 170, "User height.")
	configworkoutCmd.Flags().IntP("weight", "w", 70, "User weight.")
	configworkoutCmd.Flags().StringP("gender", "g", "Male", "User gender.")
	configworkoutCmd.Flags().IntP("days", "d", 3, "Days user will workout per week.")
	configworkoutCmd.Flags().StringP("goal", "o", "", "User Goal.")
	configworkoutCmd.Flags().StringP("type", "t", "cardio", "User workout type.")
	configworkoutCmd.Flags().StringP("info", "i", "", "Additional information to be provided in message.")

	configworkoutCmd.Flags().BoolP("edit", "e", false, "Edit existing config file.")
	rootCmd.AddCommand(configworkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
