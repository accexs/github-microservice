package config

import "os"

const apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"

var (
	gitHubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGitHubAccessToken() string {
	return gitHubAccessToken
}
