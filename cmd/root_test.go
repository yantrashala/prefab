// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yantrashala/prefab/model"
)

func TestRootCmd_NameFlag(t *testing.T) {
	assert := require.New(t)

	projectName = ""
	rootCmd.SetArgs([]string{"-n", "projectx"})
	rootCmd.Execute()
	assert.Equal("projectx", projectName, "-n argument sets the projectname property")
	assert.Equal("projectx", model.CurrentProject.Name, "-n argument sets the CurrentProject.Name property")

	projectName = ""
	rootCmd.SetArgs([]string{"--name", "projectxyz"})
	rootCmd.Execute()
	assert.Equal("projectxyz", projectName, "--name argument sets the projectname property")
	assert.Equal("projectxyz", model.CurrentProject.Name, "--name argument sets the CurrentProject.Name property")

	/* TODO: get this test case to pass
	os.Setenv("FAB_PROJECTNAME", "envprojectx")
	projectName = ""
	rootCmd.SetArgs([]string{})
	rootCmd.Execute()
	assert.Equal("envprojectx", projectName, "FAB_PROJECTNAME sets the projectname property")
	assert.Equal("envprojectx", model.CurrentProject.Name, "FAB_PROJECTNAME sets the CurrentProject.Name property")
	*/
}
