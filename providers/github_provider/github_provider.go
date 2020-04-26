package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/accexs/github-microservice/clients/restclient"
	"github.com/accexs/github-microservice/domain/github_domain"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthHeader(accessRoken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessRoken)
}

func CreateRepo(accessToken string, request github_domain.CreateRepoRequest) (*github_domain.CreateRepoResponse, *github_domain.GithubErrorResponse) {
	//Authorization 266146a5fa8eba28bfc8e9d9bd72ad96f9b85d07
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthHeader(accessToken))
	response, err := restclient.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Printf(fmt.Sprintf("Error trying to create a new repo in github: %s\n", err.Error()))
		return nil, &github_domain.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github_domain.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github_domain.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github_domain.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github_domain.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf(fmt.Sprintf("Error trying to unmarshal successfull response: %s\n", err.Error()))
		return nil, &github_domain.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error trying to unmarshal github create repo response",
		}
	}

	return &result, nil
}
