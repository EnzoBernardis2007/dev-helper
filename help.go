package main

import (
	"fmt"
	"os"
)

func HandleHelp() {
	if len(os.Args) < 3 {
		return
	}

	arg := os.Args[2]

	switch arg {
	case "introduction":
		Introduction()
	}
}

func Introduction() {
	var flag string
	if len(os.Args) >= 4 {
		flag = os.Args[3]
	}

	if flag == "--basic" {
		fmt.Println("---> Tasks")
		fmt.Println(">> The 'tasks' command helps you manage a simple command-line task system, making it easy to organize your project and keep track of what needs to be done.")
		// when i add more commands, i insert then here
		return
	}

	// when the user types 'introduction' without any flag
	fmt.Println(`    ___                        _                 
   /   \_____   __   /\  /\___| |_ __   ___ _ __ 
  / /\ / _ \ \ / /  / /_/ / _ \ | '_ \ / _ \ '__|
 / /_//  __/\ V /  / __  /  __/ | |_) |  __/ |   
/___,' \___| \_/   \/ /_/ \___|_| .__/ \___|_|   
                                |_|               `)
	fmt.Println(">> Dev Helper aims to help you organize your projects by monitoring tasks, ...")
	fmt.Println(">> This command can remind you of Dev Helper features and configure commands for you to make your own reminders about a specific project!")
	fmt.Println(">> Run 'help introduction --basics' to see the basic commands")
	fmt.Println(">> Run 'help <command>' to see more about a specific command")
}