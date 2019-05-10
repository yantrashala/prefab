// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>

package cmd

import (
	"fmt"
	"strings"

	"github.com/yantrashala/prefab/model"

	"github.com/spf13/cobra"
)

var cName string

// versionCmd represents the version command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "manage confguration values for the prefab tool",
	Run: func(cmd *cobra.Command, args []string) {
		//prefabInfo := prefab.NewInfo("")
		fmt.Println(colors.Green(BuildVersionString()))
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list the configuration values for prefab",
	Run: func(cmd *cobra.Command, args []string) {
		// prefabInfo := prefab.NewInfo("")
		fmt.Print(model.GetConfigAsYaml())
		fmt.Println("")
	},
}

func setGitConfiguration(args []string) {
	parts := strings.Split(args[0], ".")
	if parts[0] != "git" {
		return
	}

	if len(args) > 1 {
		if len(parts) < 2 {
			fmt.Println("Invalid key. usage prefab config git.kind")
		}
	}
}

var setCmd = &cobra.Command{
	Use:   "set key [value]",
	Short: "set a configuration value",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var rootName = strings.Split(args[0], ".")[0]
		switch rootName {
		case "git":
			setGitConfiguration(args)
		default:
			fmt.Println("Unknown configuration key: ", args[0])
		}

	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a configuration value",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(listCmd)

	configCmd.AddCommand(setCmd)

	configCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
