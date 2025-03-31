package main

import (
	"fmt"
	"os"
	"strings"
)

func HandleGit() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Missing argument for HandleGit")
		return
	}

	switch os.Args[2] {
	case "config-gitignore":
		ConfigGitignore()
	default:
		fmt.Println("Unknown command:", os.Args[2])
	}
}

func ConfigGitignore() {
	filename := ".gitignore"
	content := `# Dev Helper area, don't touch!
dh/`

	data, err := os.ReadFile(filename)
	if err == nil {
		if strings.Contains(string(data), content) {
			fmt.Println("The rules are already present in .gitignore!")
			return
		}

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening .gitignore:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString("\n\n" + content)
		if err != nil {
			fmt.Println("Error writing to .gitignore:", err)
			return
		}

		fmt.Println("New rules added to .gitignore successfully!")
		return
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating .gitignore:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to .gitignore:", err)
		return
	}

	fmt.Println(".gitignore created successfully!")
}
