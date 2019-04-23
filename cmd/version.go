// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>

package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	prefab "github.com/yantrashala/prefab/common"
)

var (
	versionNumber float32 = 0.2
	patchLevel            = 0
	versionSuffix         = "-DEV"
	commitHash            = " "
	buildDate             = " "
)

// CurrentVersion represents the current build version.
// This should be the only one.
var CurrentVersion prefab.Version

// BuildVersionString creates a version string. This is what you see when
// running "prefab version".
func BuildVersionString() string {
	program := "prefab "

	version := "v" + CurrentVersion.String()
	if commitHash != "" {
		version += "-" + strings.ToUpper(commitHash)
	}

	osArch := runtime.GOOS + "/" + runtime.GOARCH

	date := buildDate
	if date == "" {
		date = "unknown"
	}

	return fmt.Sprintf("%s %s %s BuildDate: %s", program, version, osArch, date)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of the prefab tool",
	Long:  `version of the prefab tool, in x.y.z format.`,
	Run: func(cmd *cobra.Command, args []string) {
		//prefabInfo := prefab.NewInfo("")
		fmt.Println(colors.Green(BuildVersionString()))
	},
}

func init() {
	CurrentVersion = prefab.Version{
		Number:     versionNumber,
		PatchLevel: patchLevel,
		Suffix:     versionSuffix,
	}
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
