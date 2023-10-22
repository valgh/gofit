/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"gofit/model"
	"gofit/utils"

	"github.com/spf13/cobra"
)

const gptFileName = "gptConfig.yaml"

// configCmd represents the config command
var configgptCmd = &cobra.Command{
	Use:   "cfggpt",
	Short: "Create a new ChatGPT config file or edit an existing one.",
	Long:  `Create a new ChatGPT config file, or edit an existing one.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.FileExists(gptFileName) {
			parseFlagsAndGenerateConfig(cmd)
		} else {
			log.Printf("Your config file already exists.")
		}
	},
}

func parseFlagsAndGenerateConfig(cmd *cobra.Command) {

	apiKey, errAK := cmd.Flags().GetString("apikey")
	endpoint, errEP := cmd.Flags().GetString("endpoint")
	modelGPT, errMD := cmd.Flags().GetString("model")
	maxtokens, errMT := cmd.Flags().GetInt("maxtokens")

	if errAK != nil || errEP != nil || errMD != nil || errMT != nil {
		log.Fatalf("Invalid or missing parameter. Please re-enter the command with valid input.")
		return
	}

	utils.GenerateYamlFile(model.ChatGPTConfig{
		Apikey:    apiKey,
		Endpoint:  endpoint,
		Model:     modelGPT,
		MaxTokens: maxtokens,
	}, gptFileName)
}

func init() {
	configgptCmd.Flags().StringP("apikey", "a", "myapikey", "Apikey to be used with ChatGPT.")
	configgptCmd.Flags().StringP("endpoint", "e", "https://api.openai.com/v1/chat/completions", "HTTPS ChatGPT endpoint to be called.")
	configgptCmd.Flags().StringP("model", "m", "gpt-3.5-turbo", "ChatGPT model.")
	configgptCmd.Flags().IntP("maxtokens", "t", 300, "Maximu number of tokens to be received in ChatGOT response.")

	configgptCmd.MarkFlagRequired("apikey")
	configgptCmd.MarkFlagRequired("endpoint")
	configgptCmd.MarkFlagRequired("model")
	configgptCmd.MarkFlagRequired("maxtokens")
	rootCmd.AddCommand(configgptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
