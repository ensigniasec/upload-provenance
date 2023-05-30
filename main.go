package main

import (
	"fmt"

	gh "github.com/sethvargo/go-githubactions"
)

func main() {
	bin := gh.GetInput("binary")
	gh.Infof("Binary path", bin)

	setOutput("url", "https://console.ensignia.dev/")
}

func setOutput(key, value string) {
	fmt.Printf("%s=%s >> $GITHUB_OUTPUT\n", key, value)
}
