package github

import (
	"log"

	"github.com/LER0ever/AOSC-Bot/models"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func InitClient() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: models.CurConfig.GhToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("", nil)
	if err != nil {
		log.Fatalf("Error accessing repos: %s", err.Error())
	}
	log.Printf("%v", repos)
}
