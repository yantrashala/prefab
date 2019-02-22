// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var projectName string
var projectDir string
var userLicense string

// colorizer
var colors aurora.Aurora

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prefab",
	Short: "prefab's for your application",
	Long: `◤◣ fab ◤◣
A tool to get prefabricated production ready code as a starter for your next adventure. 
`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose := viper.GetBool("verbose")

		fmt.Println(colors.Gray("\u25E4\u25E3"), colors.Bold(colors.Blue(" prefab ")), colors.Gray("◤◣"))

		prompt := promptui.Prompt{
			Label: "Project Name",
			Validate: func(input string) error {
				if len(input) < 3 {
					return errors.New("Project name must have at least 3 characters")
				}
				return nil
			},
		}

		if projectName == "" {
			pName, err := prompt.Run()
			if err != nil {
				fmt.Println(colors.Red("!! error -"), err)
				os.Exit(1)
			}
			projectName = pName
		}

		if verbose == true {
			fmt.Printf("Creating project %q\n", colors.Bold(colors.Green(projectName)))
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("!! error -", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// persistent flags, global for the application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prefab.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "name of the project")
	rootCmd.PersistentFlags().StringVarP(&projectDir, "projectdir", "d", "", "project directory eg. github.com/acme/project")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().Bool("verbose", false, "toogle verbose logging")
	rootCmd.PersistentFlags().Bool("nocolors", false, "toogle use of colors in cli mode")
	// rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectdir", rootCmd.PersistentFlags().Lookup("projectdir"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("nocolors", rootCmd.PersistentFlags().Lookup("nocolors"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "MIT")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	noColors := viper.GetBool("nocolors")
	colors = aurora.NewAurora(!noColors)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(colors.Red("!! error -"), err)
			os.Exit(1)
		}

		// Search config in home directory with name ".prefab" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".prefab")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("fab")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") == true {
			fmt.Println(">> Using config file:", viper.ConfigFileUsed())
		}
	} else {
		fmt.Println(colors.Red("!! error -"), err)
	}
}
