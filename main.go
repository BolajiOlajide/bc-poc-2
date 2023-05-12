package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/go-github/v52/github"
)

var (
	token  = flag.String("token", "", "mandatory github token")
	repo   = flag.String("repo", "sourcegraph-testing/markdowns", "repository to run git operations on")
	branch = flag.String("branch", "bc-poc", "branch name")
	owner  = flag.String("owner", "", "the Github username of the account to use")
)

func main() {
	flag.Parse()

	if *token == "" {
		log.Fatal("github token is required")
	}

	if *owner == "" {
		log.Fatal("owner is required")
	}

	ctx := context.Background()
	client := github.NewTokenClient(ctx, *token)

	path := "file.md"
	content := []byte("# Hola\n* This is a file committed via Github's API\n")
	message := "hopefully signed commit"

	branchRef, err := createBranch(ctx, client, *branch)
	if err != nil {
		log.Fatalf("error creating branch %s", err.Error())
	}

	_, _, err = client.Repositories.CreateFile(ctx, *owner, *repo, path, &github.RepositoryContentFileOptions{
		Message: github.String(message),
		Content: content,
		Branch:  github.String(*branch),
	})
	if err != nil {
		fmt.Println(err)
	}
}

func createBranch(ctx context.Context, client *github.Client, branch string) (*github.Reference, error) {
	// Fetching the latest commit on the master branch
	ref, _, err := client.Git.GetRef(ctx, "<owner>", "<repo>", "refs/heads/master")
	if err != nil {
		return nil, err
	}

	// Creating a new branch
	newRef, _, err := client.Git.CreateRef(ctx, "<owner>", "<repo>", &github.Reference{
		Ref:    github.String(fmt.Sprintf("refs/heads/%s", branch)),
		Object: &github.GitObject{SHA: ref.Object.SHA},
	})
	if err != nil {
		return nil, err
		fmt.Println(err)
	}
	return newRef, nil
}
