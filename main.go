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

	createBranch

	// create the branch and get it's reference
	client.Git.CreateRef(ctx, *owner, *repo, &github.Reference{})

	// 	owner: 'OWNER',
	//   repo: 'REPO',
	//   message: 'my commit message',
	//   author: {
	//     name: 'Mona Octocat',
	//     email: 'octocat@github.com',
	//     date: '2008-07-09T16:13:30+12:00'
	//   },
	//   parents: [
	//     '7d1b31e74ee336d15cbd21741bc88a537ed063a0'
	//   ],
	//   tree: '827efc6d56897b048c772eb4087f854f46256132',
	//   signature: '-----BEGIN PGP SIGNATURE-----\n\niQIzBAABAQAdFiEESn/54jMNIrGSE6Tp6cQjvhfv7nAFAlnT71cACgkQ6cQjvhfv\n7nCWwA//XVqBKWO0zF+bZl6pggvky3Oc2j1pNFuRWZ29LXpNuD5WUGXGG209B0hI\nDkmcGk19ZKUTnEUJV2Xd0R7AW01S/YSub7OYcgBkI7qUE13FVHN5ln1KvH2all2n\n2+JCV1HcJLEoTjqIFZSSu/sMdhkLQ9/NsmMAzpf/iIM0nQOyU4YRex9eD1bYj6nA\nOQPIDdAuaTQj1gFPHYLzM4zJnCqGdRlg0sOM/zC5apBNzIwlgREatOYQSCfCKV7k\nnrU34X8b9BzQaUx48Qa+Dmfn5KQ8dl27RNeWAqlkuWyv3pUauH9UeYW+KyuJeMkU\n+NyHgAsWFaCFl23kCHThbLStMZOYEnGagrd0hnm1TPS4GJkV4wfYMwnI4KuSlHKB\njHl3Js9vNzEUQipQJbgCgTiWvRJoK3ENwBTMVkKHaqT4x9U4Jk/XZB6Q8MA09ezJ\n3QgiTjTAGcum9E9QiJqMYdWQPWkaBIRRz5cET6HPB48YNXAAUsfmuYsGrnVLYbG+\nUpC6I97VybYHTy2O9XSGoaLeMI9CsFn38ycAxxbWagk5mhclNTP5mezIq6wKSwmr\nX11FW3n1J23fWZn5HJMBsRnUCgzqzX3871IqLYHqRJ/bpZ4h20RhTyPj5c/z7QXp\neSakNQMfbbMcljkha+ZMuVQX1K9aRlVqbmv3ZMWh+OijLYVU2bc=\n=5Io4\n-----END PGP SIGNATURE-----\n',
	//   headers: {
	//     'X-GitHub-Api-Version': '2022-11-28'
	//   }
	// 	client.Git.CreateCommit(ctx, "BolajiOlajide", *repo, &github.Commit{
	// Parents: [,
	// 	})
	// client.PullRequests.Create(ctx, "BolajiOlajide", *repo, &github.NewPullRequest{})

	// create a new private repository named "foo"
	// repo := &github.Repository{
	// 	Name:    github.String(repo),
	// 	Private: github.Bool(false),
	// }
	// repo.
	fmt.Println("token is ", *token)
	fmt.Println("repo is ", repo)
	fmt.Println("branch is ", branch, client)
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
