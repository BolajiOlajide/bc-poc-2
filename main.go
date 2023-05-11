package main

import (
	"flag"
	"log"

	"github.com/google/go-github/v52/github"
)

var githubToken string
var repo string
var branch string

func main() {
	flag.StringVar(&githubToken, "token", "", "mandatory github token")
	flag.StringVar(&repo, "repo", "", "repository to run PoC against")
	flag.StringVar(&branch, "branch", "bc-poc", "branch name")
	flag.Parse()

	if githubToken == "" {
		log.Fatal("github token is required")
	}

	// create a new private repository named "foo"
	repo := &github.Repository{
		Name:    github.String(repo),
		Private: github.Bool(false),
	}
	repo.
		fmt.Println("token is ", githubToken)
}
