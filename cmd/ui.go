// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>

package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// interactiveCmd represents the interactive mode of pre-fab command
var uiCmd = &cobra.Command{
	Annotations: map[string]string{"sequence": "3", "chainable": "false"},
	Long:        "UI template",
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "UI Template",
			Items: []string{"Angular", "React", "Vue"},
		}
		prompt.Run()
	},
	Short: "prefab ui template",
	Use:   "ui",
}

func init() {
	rootCmd.AddCommand(uiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
