package main

import (
	"fmt"
	"os"

	gh "github.com/sethvargo/go-githubactions"
)

var Version = "dev"
var Commit = "unknown"
var CommitDate = "unknown"
var TreeState = "unknown"

func main() {
	gh.Infof("Ensignia Action Version: %s", Version)

	apiKey := gh.GetInput("api_key")
	if apiKey == "" {
		gh.Fatalf("api_key input param is required")
	}

	bin := gh.GetInput("binary")
	gh.Infof("Binary path: %s", bin)

	provenancePath := gh.GetInput("attestation")
	gh.Infof("Provenance path: %s", bin)

	fi, err := os.Stat(provenancePath)
	if err != nil {
		gh.Fatalf("Failed to stat provenance file: %s", err)
	}

	gh.Infof("Provenance file size: %d", fi.Size())

	provFile, err := os.ReadFile(provenancePath)
	if err != nil {
		gh.Fatalf("Failed to read provenance file: %s", err)
	}

	gh.Infof("Provenance file contents: %s", string(provFile)[:100])

	// entries, err := os.ReadDir("./")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, e := range entries {
	// 	fmt.Println(e.Name())
	// }

	setOutput("url", "https://console.ensignia.dev/")
}

func setOutput(key, value string) {
	fmt.Printf("%s=%s >> $GITHUB_OUTPUT\n", key, value)
}
