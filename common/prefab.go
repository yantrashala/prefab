package prefab

const (
	EnvironmentDevelopment = "development"
	EnvironmentProduction  = "production"
)

var (
	// commitHash contains the current Git revision. Use make to build to make
	// sure this gets set.
	commitHash string

	// buildDate contains the date of the current build.
	buildDate string
)

// Info contains information about the current Hugo environment
type Info struct {
	CommitHash string
	BuildDate  string

	// The build environment.
	// Defaults are "production" (prefab) and "development" (prefab server).
	// This can also be set by the user.
	// It can be any string, but it will be all lower case.
	Environment string
}

// Version returns the current version as a comparable version string.
func (i Info) Version() VersionString {
	return CurrentVersion.Version()
}

// NewInfo creates a new Hugo Info object.
func NewInfo(environment string) Info {
	if environment == "" {
		environment = EnvironmentProduction
	}
	return Info{
		CommitHash:  commitHash,
		BuildDate:   buildDate,
		Environment: environment,
	}
}
