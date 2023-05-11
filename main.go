package main

import (
	"flag"
	"fmt"
	"log"
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

	fmt.Println("token is ", githubToken)
}
