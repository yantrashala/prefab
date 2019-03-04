package prefab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateNames(t *testing.T) {
	assert := require.New(t)
	a := []string{"aa"}
	n := []string{"nn", "na", "an"}

	r := generateName(a, n, false)
	assert.Equal(r, a[0]+"-"+n[0])

	r = generateName(a, n, true)
	assert.Equal(r, a[0]+"-"+n[2])
}
