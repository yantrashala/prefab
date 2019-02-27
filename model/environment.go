package model

// Environment describes a build or a run environment
type Environment struct {
	Name string
	Type string
}

// GetBuildEnvironmentTypes returns a list of supported environment types
func GetBuildEnvironmentTypes() ([]string, error) {
	//TODO: get from git repo
	return []string{"local(Docker)", "AWS", "GCP", "AZURE"}, nil
}
