package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

func printDirectoryContents(ctx context.Context, client *github.Client, owner, repo, dir string) error {
	opt := &github.RepositoryContentGetOptions{}
	_, entries,  _, err := client.Repositories.GetContents(ctx, owner, repo, dir, opt)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if *entry.Type == "dir" {
			fmt.Printf("Directory: %s\n", *entry.Path)
			err := printDirectoryContents(ctx, client, owner, repo, *entry.Path)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("File: %s\n", *entry.Path)
		}
	}

	return nil
}

func main() {
	ctx := context.Background()

	// Set up the GitHub client with your personal access token
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Environment variable GITHUB_TOKEN is not set")
		return
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Replace "OWNER" and "REPO" with the repository owner and name
	owner := "shahab-github"
	repo := "services"

	err := printDirectoryContents(ctx, client, owner, repo, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}