package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Task struct {
	Task string
	Priority int
	Deadline string
}

func main() {
	AddTask("Do the dishes", 1, "in progress", "2021-09-01")
	ReadTasks()
}

func ReadTasks(){
	// file, err := os.OpenFile("./tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file, err := os.Open("./tasks.csv")
	if err != nil {
		fmt.Println("Error opening file: %v", err)
	}

	r := csv.NewReader(file)

	records, err := r.ReadAll()

	if err != nil {
		fmt.Println("Error reading file: %v", err)
	}

	for _, row := range records {
		fmt.Println(row)
	}
}

func AddTask(task string, priority int, status string, deadline string) {
	file, err := os.OpenFile("./tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error opening file: %v", err)
	}

	w := csv.NewWriter(file)
	if err != nil {
		fmt.Println("Error creating filewriter: %v", err)
	}

	err = w.Write([]string{"\n" + task, fmt.Sprintf("%d", priority), status, deadline})

	if err != nil {
		fmt.Println("Error writing to file: %v", err)
	}

	w.Flush()

	file.Close()
}

// func AddTask(task string) {
// 	file, err := os.Open("../tasks.csv")
// 	if err != nil {
// 		fmt.Println("Error: %v", err)
// 	}

// 	w := csv.NewWriter(file)

// 	err = w.Write([]string{"\n" + task})

// 	if err != nil {
// 		fmt.Println("Error could not write to file| Error: %v", err)
// 	}

// 	defer
// 		w.Flush()

// 		if err = w.Error(); err != nil {
// 			fmt.Println("Error could flush filewriter| Error: %v", err)
// 		}

// 		file.Close()
// }