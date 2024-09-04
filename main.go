package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	AddTask("task1,1,tomorrow")
}

func AddTask(task string) {
	file, err := os.Open("./tasks.csv")
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	w := csv.NewWriter(file)

	err = w.Write([]string{"\n" + task})

	if err != nil {
		fmt.Println("Error could not write to file| Error: %v", err)
	}

	defer
		w.Flush()

		if err = w.Error(); err != nil {
			fmt.Println("Error could flush filewriter| Error: %v", err)
		}

		file.Close()
}