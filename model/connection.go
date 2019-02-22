package model

// Connection holds the parameters required for connecting to github, AWS, kubernetes etc.,
type Connection struct {
	Name  string
	Type  string
	Token string
}
