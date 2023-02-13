package projectdetails

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func TestGetProjectDetailsGitHub(t *testing.T) {
	details := GetProjectDetails("git", "https://github.com/cidverse/normalizeci.git", "github", "github.com")

	assert.Equal(t, "205438004", details["NCI_PROJECT_ID"])
	assert.Equal(t, "normalizeci", details["NCI_PROJECT_NAME"])
	assert.Equal(t, "cidverse-normalizeci", details["NCI_PROJECT_SLUG"])
	assert.Equal(t, "A tool to turn the continuous integration / deployment variables into a common format for generally usable scripts without any dependencies.", details["NCI_PROJECT_DESCRIPTION"])
	assert.Equal(t, "cicd,normalization", details["NCI_PROJECT_TOPICS"])
	assert.Equal(t, "https://api.github.com/repos/cidverse/normalizeci/issues/{ID}", details["NCI_PROJECT_ISSUE_URL"])
	assert.NotEmpty(t, details["NCI_PROJECT_STARGAZERS"])
	assert.NotEmpty(t, details["NCI_PROJECT_FORKS"])
	assert.Equal(t, "https://github.com/cidverse/normalizeci", details["NCI_PROJECT_URL"])
	assert.Equal(t, "main", details["NCI_PROJECT_DEFAULT_BRANCH"])
}

func TestGetProjectDetailsGitLab(t *testing.T) {
	details := GetProjectDetails("git", "https://gitlab.com/PhilippHeuer/events4j.git", "gitlab", "gitlab.com")

	assert.Equal(t, "6364957", details["NCI_PROJECT_ID"])
	assert.Equal(t, "Events4J", details["NCI_PROJECT_NAME"])
	assert.Equal(t, "philipp-heuer-events4j", details["NCI_PROJECT_SLUG"])
	assert.Equal(t, "Java Event Dispatcher / Consumer", details["NCI_PROJECT_DESCRIPTION"])
	assert.Equal(t, "", details["NCI_PROJECT_TOPICS"])
	assert.Equal(t, "https://gitlab.com/PhilippHeuer/events4j/-/issues/{ID}", details["NCI_PROJECT_ISSUE_URL"])
	assert.NotEmpty(t, details["NCI_PROJECT_STARGAZERS"])
	assert.NotEmpty(t, details["NCI_PROJECT_FORKS"])
	assert.Equal(t, "https://gitlab.com/PhilippHeuer/events4j", details["NCI_PROJECT_URL"])
	assert.Equal(t, "master", details["NCI_PROJECT_DEFAULT_BRANCH"])
}
