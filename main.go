package main

import (
	"context"
	"fmt"

	gh "github.com/google/go-github/v52/github"
	gha "github.com/sethvargo/go-githubactions"
)

var Version = "dev"
var Commit = "unknown"
var CommitDate = "unknown"
var TreeState = "unknown"

func main() {
	ctx := context.Background()
	if err := realMain(ctx); err != nil {
		gha.Fatalf("error: %s", err)
	}
}

func realMain(ctx context.Context) error {
	gha.Infof("Ensignia Action Version: %s", Version)

	apiKey := gha.GetInput("api-key")
	if apiKey == "" {
		gha.Fatalf("api-key input param is required")
	}

	ghToken := gha.GetInput("repo-token")
	if apiKey == "" {
		gha.Fatalf("repo-token input param is required")
	}

	ghContext, err := gha.Context()
	if err != nil {
		return err
	}

	owner, repo := ghContext.Repo()
	runID := ghContext.RunID

	client := gh.NewTokenClient(ctx, ghToken)
	list, _, err := client.Actions.ListWorkflowRunArtifacts(ctx, owner, repo, runID, &gh.ListOptions{
		PerPage: 100,
	})
	// list, _, err := client.Actions.ListArtifacts(ctx, owner, repo, &gh.ListOptions{
	// 	PerPage: 100,
	// })
	if err != nil {
		return err
	}

	gha.Infof("Found %d artifacts", list.GetTotalCount())

	for _, artifact := range list.Artifacts {
		if artifact.GetName() == gha.GetInput("attestation") {
			gha.Infof("Found Artifact: Attestation %s", artifact.GetName())
		}

		gha.Infof("Artifact: %s", artifact.GetName())
	}

	bin := gha.GetInput("binary")
	gha.Infof("Binary path: %s", bin)

	setOutput("url", "https://console.ensignia.dev/")
	return nil
}

func setOutput(key, value string) {
	fmt.Printf("%s=%s >> $GITHUB_OUTPUT\n", key, value)
}
