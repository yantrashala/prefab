// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yantrashala/prefab/model"
)

var verbose bool
var cfgFile string
var projectName string
var projectDir string
var userLicense string
var tempDir string
var saveWorkDir bool

// colorizer
var colors aurora.Aurora

func setProjectName() {
	if projectName == "" {
		if verbose == true {
			fmt.Println("Creating new project ...")
		}
		prompt := promptui.Prompt{
			Label:   "Project Name",
			Default: model.CurrentProject.Name,
			Validate: func(input string) error {
				if len(input) < 3 {
					return errors.New("Project name must have at least 3 characters")
				}
				return nil
			},
		}

		pName, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		model.CurrentProject.SetProjectName(pName)
	} else {
		model.CurrentProject.SetProjectName(projectName)
	}
}

func createBuildEnvironment() {
	if verbose {
		fmt.Println("Creating build environment...")
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F449 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | blue }}",
		Selected: "Build environment: {{ .Name | green | bold}}",
		Details: `
--------- Build Environment ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Repo:" | faint }}	{{ .Repo }}`,
	}
	el, _ := model.GetBuildEnvironmentTypes()
	prompt := promptui.Select{
		Templates: templates,
		Label:     "Select the build environment type",
		Items:     el,
	}
	i, _, _ := prompt.Run()
	env := model.Environment(el[i])
	env.Type = "build"
	model.CurrentProject.AddEnvironment(env)
}

func createEnvironment() {
	if verbose {
		fmt.Println("Create new environment ...")
	}
	envNamePrompt := promptui.Prompt{
		Label: "Name of the environment",
	}
	envName, _ := envNamePrompt.Run()
	env := model.Environment{Name: envName}
	model.CurrentProject.AddEnvironment(env)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prefab",
	Short: "prefab's for your application",
	Long: `◤◣ prefab ◤◣
A tool to get prefabricated production ready code as a starter for your next adventure. 
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if saveWorkDir {
			fmt.Println("WORK=" + tempDir)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(colors.Gray("\u25E4\u25E3"), colors.Bold(colors.Blue(" prefab ")), colors.Gray("◤◣"))

		setProjectName()

		createBuildEnvironment()

		promptText := "Do you want a create a new run environment? [y/N]"

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "\U0001F449 {{ . | cyan }}",
			Inactive: "  {{ . | blue }}",
			Selected: " ",
		}

		prompt := promptui.Select{
			Templates: templates,
			Label:     promptText,
			Items:     []string{"y", "N"},
		}
		_, option, _ := prompt.Run()
		for option == "y" {
			createEnvironment()
			_, option, _ = prompt.Run()
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println(model.CurrentProject)
		if saveWorkDir == false {
			defer os.RemoveAll(tempDir)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	var err error
	tempDir, err = ioutil.TempDir("", "prefab")
	if err != nil {
		log.Fatal(err)
	}

	// persistent flags, global for the application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prefab/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "name of the project")
	rootCmd.PersistentFlags().StringVarP(&projectDir, "projectdir", "d", tempDir, "project directory eg. /Users/username/.prefab/projects")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "toogle verbose logging")
	rootCmd.PersistentFlags().Bool("nocolors", false, "toogle use of colors in cli mode")
	rootCmd.PersistentFlags().BoolVarP(&saveWorkDir, "work", "w", false, "print the name of the temporary work directory and do not delete it when exiting.")

	// rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("work", rootCmd.PersistentFlags().Lookup("work"))
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectName", rootCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("projectdir", rootCmd.PersistentFlags().Lookup("projectdir"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("nocolors", rootCmd.PersistentFlags().Lookup("nocolors"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))

	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "MIT")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".prefab" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".prefab"))
		viper.SetConfigName("config.yaml")
	}

	viper.SetEnvPrefix("fab")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	configerr := viper.ReadInConfig()

	verbose = viper.GetBool("verbose")
	noColors := viper.GetBool("nocolors")
	colors = aurora.NewAurora(!noColors)

	if configerr == nil {
		if verbose {
			fmt.Println("Using config file: ", colors.Green(viper.ConfigFileUsed()))
		}
	}
}
