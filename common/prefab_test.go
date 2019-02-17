package prefab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefabInfo(t *testing.T) {
	assert := require.New(t)

	prefabInfo := NewInfo("")

	assert.Equal(CurrentVersion.Version(), prefabInfo.Version())
	assert.IsType(VersionString(""), prefabInfo.Version())
	assert.Equal(commitHash, prefabInfo.CommitHash)
	assert.Equal(buildDate, prefabInfo.BuildDate)
	assert.Equal("production", prefabInfo.Environment)
}
