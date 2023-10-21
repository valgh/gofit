/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gofit/builder"
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
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			log.Fatalf("Invalid command or parametger: %v", err)
			return
		}

		if message == "" {
			message = buildFromConfig()
		}
		getGPTMessage(message)
	},
}

func buildFromConfig() string {
	var message string
	if utils.FileExists(configFileName) {
		message = builder.BuildMessage(configFileName)
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
	workoutCmd.Flags().StringP("message", "m", "", "Enter request message to be sent to ChatGPT endpoint.")
	rootCmd.AddCommand(workoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scheduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scheduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
