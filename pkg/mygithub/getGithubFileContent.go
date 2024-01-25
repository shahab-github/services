
package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

func getFileContent(client *github.Client, owner, repo, path string) {
	ctx := context.Background()

	// Get the content of the file
	fileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Decode and print the content
	content, err := base64.StdEncoding.DecodeString(*fileContent.Content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File Content:\n%s\n", content)
	fmt.Println("----")
}

func listFilesAndFoldersRecursively(client *github.Client, owner, repo, path string) {
	ctx := context.Background()

	// List the contents of the specified path in the repository
	_, contents, _, err := client.Repositories.GetContents(ctx, owner, repo, path, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, content := range contents {
		fmt.Printf("Path: %s\n", *content.Path)
		fmt.Printf("Type: %s\n", *content.Type)

		if *content.Type == "file" {
			// If it's a file, display its content
			getFileContent(client, owner, repo, *content.Path)
		} else if *content.Type == "dir" {
			// If it's a directory, recursively list its contents
			listFilesAndFoldersRecursively(client, owner, repo, *content.Path)
		}

		fmt.Println("----")
	}
}

func listFilesAndFolders(owner, repo string) {
	ctx := context.Background()

	// Replace "YOUR-GITHUB-TOKEN" with your actual GitHub token.
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "your-Token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Start listing from the root of the repository
	listFilesAndFoldersRecursively(client, owner, repo, "")
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <owner> <repo>")
		os.Exit(1)
	}

	owner := os.Args[1]
	repo := os.Args[2]

	listFilesAndFolders(owner, repo)
}
