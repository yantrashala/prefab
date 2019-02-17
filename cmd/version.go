// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	prefab "github.com/yantrashala/prefab/common"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of the prefab tool",
	Long:  `version of the prefab tool, in x.y.z format.`,
	Run: func(cmd *cobra.Command, args []string) {
		//prefabInfo := prefab.NewInfo("")
		fmt.Println(colors.Green(prefab.BuildVersionString()))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
