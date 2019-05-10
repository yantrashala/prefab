// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	prefab "github.com/yantrashala/prefab/common"
	"github.com/yantrashala/prefab/model"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/client"
	githttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/yaml.v2"
)

var verbose bool
var cfgFile string
var projectName string
var projectDir string
var userLicense string
var tempDir string
var saveWorkDir = true

// colorizer
var colors aurora.Aurora

func promptForYN(question string) (string, error) {
	promptText := question + " [y/N]"
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

	_, option, err := prompt.Run()

	return option, err
}

func setProjectName() {
	if projectName == "" {
		if verbose == true {
			fmt.Println(colors.Gray("Creating new project ..."))
		}
		prompt := promptui.Prompt{
			Label:   "Project Name: ",
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

func setProjectDir() {
	if projectDir == "" {

	}

	if verbose == true {
		fmt.Println(colors.Gray("Project directory: "), colors.Green(projectDir))
	}

	model.CurrentProject.SetLocalDirectory(projectDir)
}

func promptForConfigValues(dir string, configValues map[string]string) error {

	content, ferr := ioutil.ReadFile(path.Join(dir, ".prefab", "config.yaml"))
	if os.IsNotExist(ferr) {
		return nil
	}

	config := make(map[string]map[string]string)

	err := yaml.Unmarshal([]byte(content), &config)

	for k := range config {
		cdata := config[k]
		configValues[k] = cdata["default"]
		cprompt := promptui.Prompt{
			Label:   cdata["prompt"],
			Default: cdata["default"],
		}

		if cvalue, cerr := cprompt.Run(); cerr == nil {
			configValues[k] = cvalue
		} else {
			fmt.Println(cerr)
		}

	}

	return err

}

func createBuildEnvironment(p bool) error {

	var err error
	var option string

	if p {
		if option, err = promptForYN("Do you want a create a new build environment?"); err != nil {
			return err
		}

		if option == "N" {
			return nil
		}
	}

	if verbose {
		fmt.Println(colors.Gray("Creating build environment..."))
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F449 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | blue }}",
		Selected: "Build environment: {{ .Name | green | bold}}",
		Details: `
--------- Build Environment ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Repo:" | faint }}	{{ .Repo }}
{{ .Summary }}`,
	}
	el, _ := model.GetBuildEnvironmentTypes()
	sprompt := promptui.Select{
		Templates: templates,
		Label:     "Select the build environment type",
		Items:     el,
	}
	i, _, serr := sprompt.Run()

	if serr != nil {
		return serr
	}

	env := model.Environment(el[i])
	env.Name = "build"
	env.Type = "build"
	model.CurrentProject.AddEnvironment(env)

	envDir := model.CurrentProject.Environments["build"].LocalDirectory

	cerr := promptForConfigValues(envDir, model.CurrentProject.Environments["build"].Config)

	if cerr != nil {
		return cerr
	}

	createRepo()

	return nil
}

func createEnvironment() {
	if verbose {
		fmt.Println("Create new environment ...")
	}
	envNamePrompt := promptui.Prompt{
		Label: "Name of the environment: ",
	}
	envName, _ := envNamePrompt.Run()
	env := model.Environment{Name: envName}

	model.CurrentProject.AddEnvironment(env)
}

func createRuntimeEnvironments(p bool, l bool) error {
	var err error
	var option string

	if p {
		if option, err = promptForYN("Do you want a create a new run environment?"); err != nil {
			return err
		}

	}

	for option == "y" {
		createEnvironment()
		if !l {
			return nil
		}
		if option, err = promptForYN("Do you want a create a new run environment?"); err != nil {
			return err
		}
	}
	return nil
}

func createRepo() model.SCM {
	stypes := model.GetSCMTypes()
	prompt := promptui.Select{
		Label: "Select the GIT server",
		Items: stypes,
	}
	i, _, _ := prompt.Run()
	scmType := stypes[i]
	fmt.Print(scmType)
	scm := model.SCM{
		Type: scmType,
	}
	return scm
}
func createApplication() {
	atypes, _ := model.GetApplicationTypes()
	prompt := promptui.Select{
		Label: "Select the Application type",
		Items: atypes,
	}
	i, _, _ := prompt.Run()
	appType := atypes[i]
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F449 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | blue }}",
		Selected: "Build environment: {{ .Name | green | bold}}",
		Details: `
--------- New Application ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Repo:" | faint }}	{{ .Repo }}
{{ .Summary }}`,
	}
	el := model.GetApplications(appType)
	prompt = promptui.Select{
		Templates: templates,
		Label:     "Select the Application",
		Items:     el,
	}
	i, _, _ = prompt.Run()
	app := model.Application(el[i])

	label := appType + " Application Name"

	promptName := promptui.Prompt{
		Label:   label,
		Default: prefab.GetSimpleName(),
		Validate: func(input string) error {
			if len(input) < 3 {
				return errors.New("Application name must have at least 3 characters")
			}
			return nil
		},
	}

	pName, perr := promptName.Run()
	if perr != nil {
		fmt.Println(perr)
		os.Exit(1)
	}
	app.Name = pName
	model.CurrentProject.AddApplication(app)

	appDir := model.CurrentProject.Applications[pName].LocalDirectory

	cerr := promptForConfigValues(appDir, model.CurrentProject.Applications[pName].Config)

	if cerr != nil {
		log.Fatal(cerr)
	}
}

func createApplications() {

	option, _ := promptForYN("Do you want a create a new application?")
	for option == "y" {
		createApplication()
		option, _ = promptForYN("Do you want a create a new application?")
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prefab",
	Short: "prefab's for your application",
	Long: `◤◣ prefab ◤◣
A tool to get prefabricated production ready code as a starter for your next adventure. 
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		/*
			if saveWorkDir {
				fmt.Println("WORK=" + tempDir)
			}
		*/

		if verbose {
			fmt.Println("Reading default config from: ", colors.Green(cfgFile))
		}
		err := model.LoadConfig(cfgFile)
		if err != nil {
			if verbose {
				fmt.Println("WARNING: ", err)
			}
		}

		if projectName != "" {
			model.CurrentProject.SetProjectName(projectName)
		}

		if projectDir != "" {
			fmt.Println("Setting project directory: ", colors.Green(projectDir))
			model.CurrentProject.SetLocalDirectory(projectDir)
			model.CurrentProject.LoadProject()
		}

	},
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(colors.Gray("\u25E4\u25E3"), colors.Bold(colors.Blue(" prefab"))) //, colors.Gray("◤◣"))

		setProjectName()

		setProjectDir()

		model.CurrentProject.LoadProject()

		createBuildEnvironment(true)

		createApplications()

		createRuntimeEnvironments(true, true)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var err error
		if verbose {
			fmt.Println("Writing project to: ", model.CurrentProject.GetProjectFilename())
		}
		err = model.CurrentProject.SaveProject()

		if err != nil {
			fmt.Printf("%+v", reflect.TypeOf(err))
			fmt.Println(colors.Red(err))
		}

		model.CurrentProject.ApplyValues()

		if verbose {
			fmt.Println("writing updated default config to: ", colors.Green(cfgFile))
		}
		err = model.SaveConfig(cfgFile)
		if err != nil {
			fmt.Println(colors.Red(err))
		}
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

	// Create a custom http(s) client with your config
	customClient := &http.Client{
		// accept any certificate (might be useful for testing)
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		// 15 second timeout
		Timeout: 15 * time.Second,

		// don't follow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Override http(s) default protocol to use our custom client
	client.InstallProtocol("https", githttp.NewClient(customClient))

	cobra.OnInitialize(initConfig)

	var err error
	tempDir, err = ioutil.TempDir("", "prefab")
	if err != nil {
		log.Fatal(err)
	}

	// persistent flags, global for the application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prefab/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "name of the project")
	rootCmd.PersistentFlags().StringVarP(&projectDir, "projectdir", "d", projectDir, "project directory eg. /Users/username/.prefab/projects")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "toogle verbose logging")
	rootCmd.PersistentFlags().Bool("nocolors", false, "toogle use of colors in cli mode")
	// rootCmd.PersistentFlags().BoolVarP(&saveWorkDir, "work", "w", false, "print the name of the temporary work directory and do not delete it when exiting.")

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
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".prefab" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".prefab"))
		viper.SetConfigName("config.yaml")
		cfgFile = filepath.Join(home, ".prefab", "config.yaml")
	}

	if projectDir == "" {
		projectDir = filepath.Join(home, ".prefab", "projects")
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
