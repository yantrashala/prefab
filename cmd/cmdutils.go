// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// General purpose method to filter commands based on custom filters & do something in success
func filterAndExecCommands(cmds []*cobra.Command,
	filterFunc func(ccmd *cobra.Command) bool,
	doFunc func(ccmd *cobra.Command) bool) []*cobra.Command {
	//
	cmdCounter := 0
	var filteredCommands []*cobra.Command
	for _, cmd := range cmds {
		if filterFunc(cmd) {
			filteredCommands = append(filteredCommands, cmd)
			doFunc(cmd)
			cmdCounter++
		}
	}
	if viper.GetBool("verbose") == true {
		fmt.Println(">>", cmdCounter, " commands executed!")
	}
	return filteredCommands
}

// General purpose method find and execute a command by name
func runCommandByName(cmdName string, cmds []*cobra.Command, args []string) {
	for _, cmd := range cmds {
		if strings.Compare(cmdName, cmd.Name()) == 0 {
			if viper.GetBool("verbose") {
				fmt.Println(">> cmd: ", cmd.Name(), " - ", cmd.Annotations)
			}
			cmd.Run(cmd, args)
		}
	}
	return
}
