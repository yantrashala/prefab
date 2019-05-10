package prefab

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateNames(t *testing.T) {
	assert := require.New(t)
	a := []string{"aa"}
	n := []string{"nn"}

	r := generateName(a, n, false)
	assert.Equal(r, a[0]+"-"+n[0])

	ad := []string{"aa"}
	no := []string{"nn", "na", "an"}
	r = generateName(ad, no, true)
	assert.Equal(r, "aa-an")
}
