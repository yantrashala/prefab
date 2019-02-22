// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>

package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// interactiveCmd represents the interactive mode of pre-fab command,
// which chains the commands one after other to provide a yeoman like experience,
// ex: ./prefab interactive
// two modes are supported for sequencing the commands:
// a. using command annotations in code  (more control and predictable) --> "sequence":"<number>", "chainable":"<true/false>"
// b. using .prefab.yaml via configuration (more flexibility) --> "cmdsequence: 'command-1,command-2,...,command-n'"
var interactiveCmd = &cobra.Command{
	Aliases:     []string{"im"},
	Annotations: map[string]string{"sequence": "3", "chainable": "false"},
	Example:     "./prefab interactive \n ./prefab im",
	Long:        "prefab command execution in interactive mode",
	Run: func(cmd *cobra.Command, args []string) {
		// Chain the other commands, enabling an interactive experience
		cmds := rootCmd.Commands()
		cseq := viper.GetString("cmdsequence")
		// Going by the configurations...
		if cseq != "" {
			cmdCounter := 0
			// Chain and Execute the commands as per the config sequence
			for _, cmdName := range strings.Split(cseq, ",") {
				runCommandByName(cmdName, cmds, args)
				cmdCounter++
			}
			if viper.GetBool("verbose") == true {
				fmt.Println(">>", cmdCounter, " commands executed!")
			}
		} else {
			// Chain and Execute the commands (sorted by sequence number) as per the annotations
			sort.SliceStable(cmds, func(srcIndex, destIndex int) bool {
				sseq, _ := strconv.Atoi(cmds[srcIndex].Annotations["sequence"])
				dseq, _ := strconv.Atoi(cmds[destIndex].Annotations["sequence"])
				//
				return (sseq < dseq)
			})
			// Filter and Execute the chainable commands
			filterAndExecCommands(
				cmds,
				func(ccmd *cobra.Command) bool {
					return (strings.Compare(ccmd.Annotations["chainable"], "true") == 0)
				},
				func(ccmd *cobra.Command) bool {
					ccmd.Run(ccmd, args)
					return true
				})
		}
	},
	Short: "prefab interactive mode",
	Use:   "interactive",
}

func init() {
	interactiveCmd.PersistentFlags().StringP("Interactive", "i", "", "interactive mode")
	rootCmd.AddCommand(interactiveCmd)
	// Define your flags and configuration settings.
}
