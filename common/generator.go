package prefab

import (
	"math/rand"
	"strings"
	"time"
)

// GenerateName using nouns & prnouns
func GenerateName(alliterative bool) string {
	return generateName(AdjectivesList, NounList, alliterative)
}

func generateName(adjectives []string, nouns []string, alliterative bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := adjectives[r.Intn(len(adjectives))]
	var fn []string
	if alliterative {
		fn = getAlliterativeMatches(nouns, string(a[0]))
	} else {
		fn = nouns
	}
	return a + "-" + fn[r.Intn(len(fn))]
}

func getAlliterativeMatches(list []string, letter string) []string {
	check := strings.ToLower(letter)
	b := list[:0]
	for _, x := range list {
		if check == string(x[0]) {
			b = append(b, x)
		}
	}
	return b
}
