package main

import (
	gh "github.com/sethvargo/go-githubactions"
)

func main() {
	bin := gh.GetInput("binary")
	gh.Infof("Binary path", bin)

	gh.SetOutput("url", "https://console.ensignia.dev/")
}
