/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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

const configFileName = "workoutConfig.yaml"
const gptFileName = "gptConfig.yaml"

var defaultWorkoutPlan = model.NewWorkoutPlan(3, "", "")

var defaultBodyBuilder = model.NewBodyBuilder(20, 180, 75, "Male", defaultWorkoutPlan)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create a default workout config file or edit an existing one.",
	Long:  `Create a default workout config file, or edit an existing one.`,
	Run: func(cmd *cobra.Command, args []string) {
		isEdit, err := cmd.Flags().GetBool("edit")
		isGPTConfig, err_ := cmd.Flags().GetBool("gpt")
		if err != nil || err_ != nil {
			return
		}

		if isGPTConfig == true {
			if utils.FileExists(gptFileName) {
				log.Fatalf("GPT Config file already exists. Please remember to add your API Key by editing such file.")
				return
			} else {

				utils.GenerateYamlFile(model.ChatGPTConfig{
					Apikey:    "",
					Endpoint:  "https://api.openai.com/v1/chat/completions",
					Model:     "gpt-3.5-turbo",
					MaxTokens: 300,
				}, gptFileName)
				return
			}
		}

		if isEdit == true && utils.FileExists(configFileName) {
			fmt.Printf("Config file exists. Entering edit mode...")
			generateYAMLFomUserInput(configFileName)
		} else if isEdit == false {
			if !utils.FileExists(configFileName) {
				utils.GenerateYamlFile(defaultBodyBuilder, configFileName)
			} else {
				fmt.Printf("Your config file already exists. To modify it, please run 'gofit config -e' or 'gofit config --edit'")
			}
		} else {
			fmt.Printf("Please generate a config file. Run 'gofit config' command.")
		}
	},
}

func generateYAMLFomUserInput(filename string) {
	var bodyBuilder model.BodyBuilder = model.BodyBuilder{}.FromYaml(filename)
	if (bodyBuilder == model.BodyBuilder{}) {
		log.Fatalf("YAML file corrupted, or malformed.")
		return
	}
	editStructFromUserInput(&bodyBuilder)
	utils.GenerateYamlFile(bodyBuilder, configFileName)
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

	bodyBuilder.SetTrainingDetails(trainingDetails)

}

func init() {
	configCmd.Flags().BoolP("gpt", "g", false, "Generate ChatGPT Config file.")
	configCmd.Flags().BoolP("edit", "e", false, "Edit existing config file.")
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
