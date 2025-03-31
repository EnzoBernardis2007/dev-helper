package main

import (
	"fmt"
	"os"
	"encoding/json"
	"time"
	"path/filepath"
)

type Status int

const (
	Pending Status = iota
	InProgress
	Completed
	Canceled
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status Status `json:"status"`
	CreationDate time.Time `json:"creation_date"`
	Term time.Time `json:"term"`
}

// HELP FUNCS
func GetLastID() int {
	if _, err := os.Stat("dh"); os.IsNotExist(err) {
		fmt.Println("No Dev Helper initialized, run tasks init to initialize")
		return -1
	}

	files, err := filepath.Glob("dh/tasks*.json")
	if err != nil {
		fmt.Println("Error searching files")
		return -1
	}

	return len(files)
}

func AlreadyHasName(name string) bool {
	files, err := filepath.Glob("dh/task*.json")
	if err != nil {
		fmt.Println("Error searching files:", err)
		return false
	}

	if len(files) <= 0 {
		fmt.Println("No task found")
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", file, err)
			continue
		}

		var task Task
		err = json.Unmarshal(data, &task)
		if err != nil {
			fmt.Println("Error decoding JSON in file:", file, err)
			continue
		}

		if name == task.Name {
			return true
		}
	}

	return false
}

// COMMANDS
func HandleTasksCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Insufficient")
	}

	switch os.Args[2] {
	case "init":
		InitTasks()
	case "create":
		CreateTask()
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

func CreateTask() {
	name := os.Args[3]
	termStr := os.Args[4]

	if AlreadyHasName(name) {
		fmt.Println("Already has a task with that name")
		return
	}

	term, err := time.Parse("2006-01-02", termStr)
	if err != nil {
		fmt.Println("Invalid data format, use YYYY-MM-DD")
		return
	}

	newID := GetLastID()
	newID++
	if newID == -1 {
		fmt.Println("Error retrieving last ID. Make sure 'tasks init' was executed first.")
		return
	}

	newTask := Task{
		ID: newID,
		Name: name,
		Description: "",
		Status: Pending,
		CreationDate: time.Now(),
		Term: term,
	}

	jsonData, err := json.MarshalIndent(newTask, "", " ")
	if err != nil {
		fmt.Println("Error trying to convert file to JSON")
		return
	}

	fileName := fmt.Sprintf("dh/task%d.json", newID)
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Error trying to create file:", err)
		return
	}

	fmt.Println("Task created succefully")
}