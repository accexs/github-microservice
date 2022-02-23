package app

import (
	"github.com/accexs/github-microservice/controllers/repositories"
	"github.com/accexs/github-microservice/controllers/status"
)

func mapUrls() {
	router.GET("/status", status.Alive )
	router.POST("/repos", repositories.CreateRepo)
}
