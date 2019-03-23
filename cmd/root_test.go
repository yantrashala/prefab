// Copyright © 2019 Publicis Sapient <EMAIL ADDRESS>
//

package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yantrashala/prefab/model"
)

func TestRootCmd_NameFlag(t *testing.T) {
	assert := require.New(t)

	projectName = ""
	rootCmd.SetArgs([]string{"-n", ".testprojectx"})
	rootCmd.Execute()
	assert.Equal(".testprojectx", projectName, "-n argument sets the projectname property")
	assert.Equal(".testprojectx", model.CurrentProject.Name, "-n argument sets the CurrentProject.Name property")

	defer os.RemoveAll(path.Join(model.CurrentProject.LocalDirectory, ".testprojectx"))

	projectName = ""
	rootCmd.SetArgs([]string{"--name", ".testprojectxyz"})
	rootCmd.Execute()
	assert.Equal(".testprojectxyz", projectName, "--name argument sets the projectname property")
	assert.Equal(".testprojectxyz", model.CurrentProject.Name, "--name argument sets the CurrentProject.Name property")

	defer os.RemoveAll(path.Join(model.CurrentProject.LocalDirectory, ".testprojectxyz"))

	/* TODO: get this test case to pass

	os.Setenv("FAB_PROJECTNAME", "envprojectx")
	projectName = ""
	rootCmd.SetArgs([]string{})
	rootCmd.Execute()
	assert.Equal("envprojectx", projectName, "FAB_PROJECTNAME sets the projectname property")
	assert.Equal("envprojectx", model.CurrentProject.Name, "FAB_PROJECTNAME sets the CurrentProject.Name property")
	*/
}
