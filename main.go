package main

import (
	"fmt"
	"log"
	"os"

	gh "github.com/sethvargo/go-githubactions"
)

func main() {
	bin := gh.GetInput("binary")
	gh.Infof("Binary path: %s", bin)

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}

	setOutput("url", "https://console.ensignia.dev/")
}

func setOutput(key, value string) {
	fmt.Printf("%s=%s >> $GITHUB_OUTPUT\n", key, value)
}
