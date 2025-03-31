package main

import (
	"fmt"
	"os"
)

func main() {
	// prevent insufficient args
	if len(os.Args) < 2 {
		fmt.Println("Insufficient arguments")
		return
	}

	command := os.Args[1]

	switch command {
	case "help":
		HandleHelp()
	case "tasks":
		HandleTasksCommand()
	case "git":
		HandleGit()
	}
}