package services

import (
	"github.com/accexs/github-microservice/config"
	"github.com/accexs/github-microservice/domain/github_domain"
	"github.com/accexs/github-microservice/domain/repositories"
	"github.com/accexs/github-microservice/providers/github_provider"
	"github.com/accexs/github-microservice/utils/errors"
	"strings"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github_domain.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     true,
	}

	response, err := github_provider.CreateRepo(config.GetGitHubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Name: response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
