package main

import (
	"fmt"
	"os"
)

func HandleTasksCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Insuffici")
	}

	switch os.Args[2] {
	case "init":
		InitTasks()
	}
}

func InitTasks() error {
	path := "./dh"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}

		fmt.Println("Dev Helper initialized")
	} else {
		fmt.Println("Dev Helper already been initialized")
	}

	return nil
}