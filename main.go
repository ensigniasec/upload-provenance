package main

import (
	"context"
	"fmt"
	"time"

	gh "github.com/google/go-github/v52/github"
	gha "github.com/sethvargo/go-githubactions"
	retry "github.com/sethvargo/go-retry"
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

	provenanceName := gha.GetInput("provenance-name")
	if provenanceName == "" {
		gha.Fatalf("provenance-name input param is required (e.g. 'needs.build.outputs.provenance-name')")
	}

	ghContext, err := gha.Context()
	if err != nil {
		return err
	}

	owner, repo := ghContext.Repo()
	runID := ghContext.RunID

	client := gh.NewTokenClient(ctx, ghToken)
	var list *gh.ArtifactList

	fn := func(ctx context.Context) error {
		results, _, err := client.Actions.ListWorkflowRunArtifacts(ctx, owner, repo, runID, &gh.ListOptions{})
		if err != nil {
			return retry.RetryableError(err)
		}

		if results.GetTotalCount() == 0 {
			gha.Warningf("Nothing found, retrying...")

			return retry.RetryableError(fmt.Errorf("no artifacts found"))
		}

		gha.Infof("Found %d artifacts", results.GetTotalCount())

		list = results
		return nil
	}

	bo := retry.WithMaxRetries(10, retry.NewExponential(time.Second))
	err = retry.Do(ctx, bo, fn)
	if err != nil {
		return err
	}

	gha.Infof("Found %d artifacts", list.GetTotalCount())

	for _, artifact := range list.Artifacts {
		if artifact.GetName() == provenanceName {
			gha.Infof("Found Artifact: Attestation %s", artifact.GetName())
		}

		gha.Infof("Artifact: %s", artifact.GetName())
	}

	// setOutput("url", "https://console.ensignia.dev/")
	return nil
}

func setOutput(key, value string) {
	fmt.Printf("%s=%s >> $GITHUB_OUTPUT\n", key, value)
}
