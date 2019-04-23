package model

import "fmt"

// User strucure
type User struct {
	Username string
	APIToken string
}

// SCM holds the parameters required for connecting to github, gogs, gitlab etc.,
type SCM struct {
	Name  string
	Type  string
	Users []User
	URL   string
}

// GetSCMTypes gets the list of the configured SCM's
func GetSCMTypes() []string {
	//configured servers
	cs := Config["git"]
	fmt.Print(cs)
	return []string{"gogs", "github"}
}
