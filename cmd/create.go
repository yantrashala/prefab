// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>

package cmd

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create using the prefab tool",
	Long:  `create apps, environments and git servers using the prefab tool`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var createAppCmd = &cobra.Command{
	Use:   "app",
	Short: "create a new application",
	Run: func(cmd *cobra.Command, args []string) {
		createApplication()
	},
}

var envType *string
var createEnvCmd = &cobra.Command{
	Use:   "env",
	Short: "create a new application",
	Run: func(cmd *cobra.Command, args []string) {

		envOption := *envType
		if *envType == "" {
			promptText := "Type of environment to create?"
			templates := &promptui.SelectTemplates{
				Label:    "{{ . }}?",
				Active:   "\U0001F449 {{ . | cyan }}",
				Inactive: "  {{ . | blue }}",
				Selected: " ",
			}

			prompt := promptui.Select{
				Templates: templates,
				Label:     promptText,
				Items:     []string{"build", "runtime"},
			}
			_, envOption, _ = prompt.Run()
		}

		switch envOption {
		case "build":
			createBuildEnvironment(false)
		case "runtime":
			createRuntimeEnvironments(false, false)
		default:
			log.Fatalln("Unknown build type")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createAppCmd, createEnvCmd)

	envType = createEnvCmd.Flags().String("type", "", "type of environment (build/runtime...)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
