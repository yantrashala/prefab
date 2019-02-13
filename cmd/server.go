// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"fmt"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts the web UI",
	Long:  `starts a web UI in the specified port (default 9876) for a interactive prefab configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prefab confuration server running at ", Bold(Green("http://localhost:9876")))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().Uint16P("port", "p", 9876, "port for the prefab ui")

	viper.BindPFlag("port", serverCmd.Flags().Lookup("port"))
	viper.SetDefault("port", 9876)
}
