package main

import "github.com/nightvalley/tammy/internal/commandline"

func main() {
	// t := time.Now()
	// duration := time.Since(t)

	flags := commandline.Flags{}
	flags.Launch()

	// log.Infof("\nExecution time: %v", duration)
}
