/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gofit/builder"
	"gofit/model"
	"gofit/rest"
	"gofit/utils"
	"log"

	"github.com/spf13/cobra"
)

// scheduleCmd represents the workout command
var workoutCmd = &cobra.Command{
	Use:   "workout",
	Short: "Ask ChatGPT your ideal workout plan.",
	Long: `Provide your data as a text input to this command, 
	or store them in the config.yaml file, in order to ask ChatGPT for your workout plan.`,
	Run: func(cmd *cobra.Command, args []string) {
		isFromConfig, err := cmd.Flags().GetBool("fromconfig")
		if err != nil {
			log.Fatalf("Invalid command or parametger: %v", err)
			return
		}

		var message string
		if isFromConfig {
			message = buildFromConfig()
		} else {
			message = parseCommandAndGetMessage(cmd)
		}
		getGPTMessage(message)
	},
}

func parseCommandAndGetMessage(cmd *cobra.Command) string {
	var message string
	bodyBuilder := generateFromCommand(cmd)
	if bodyBuilder != (model.BodyBuilder{}) {
		message = builder.BuildMessage(bodyBuilder)
	}
	return message
}

func buildFromConfig() string {
	var message string
	if utils.FileExists(workoutFileName) {
		message = builder.BuildMessageFromFile(workoutFileName)
	} else {
		log.Fatalf("Config file doesn't exist. Please create one through 'gofit config' command.")
	}
	return message
}

func getGPTMessage(message string) {
	if message != "" {
		GPTResponse := rest.GPTRestCall(message)
		fmt.Println(GPTResponse)
	} else {
		log.Printf("Message is empty, nothing to ask.")
	}

}

func init() {
	workoutCmd.Flags().BoolP("fromconfig", "f", false, "Send workout request to ChatGPT from workout config file.")

	workoutCmd.Flags().IntP("age", "a", 20, "User Age.")
	workoutCmd.Flags().IntP("height", "c", 170, "User height.")
	workoutCmd.Flags().IntP("weight", "w", 70, "User weight.")
	workoutCmd.Flags().StringP("gender", "g", "Male", "User gender.")
	workoutCmd.Flags().IntP("days", "d", 3, "Days user will workout per week.")
	workoutCmd.Flags().StringP("goal", "o", "", "User Goal.")
	workoutCmd.Flags().StringP("type", "t", "cardio", "User workout type.")
	workoutCmd.Flags().StringP("info", "i", "", "Additional information to be provided in message.")
	rootCmd.AddCommand(workoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scheduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scheduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
