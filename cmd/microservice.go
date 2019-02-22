// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>

package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// interactiveCmd represents the interactive mode of pre-fab command
var microserviceCmd = &cobra.Command{
	Annotations: map[string]string{"sequence": "3", "chainable": "false"},
	Use:         "microservice",
	Short:       "prefab microservice template",
	Long:        "Microservice template",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Microservice Template",
			Items: []string{"Spring-Boot", "Micronaut", "Barebones"},
		}
		// Todo: Add code to pull out the microservice code templates from git repo.,
		// build, deploy onto docker container and test!
		prompt.Run()
	},
}

func init() {
	rootCmd.AddCommand(microserviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// microserviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// microserviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
