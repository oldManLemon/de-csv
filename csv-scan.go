package main

import (
	"fmt"
	"os"
	"strings"
)

func analyseReplace(csvPath string) {

	csv, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csv.Close()
	// scanner :=

}

func listFiles(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() {
			// fmt.Println("Dir: ", file.Name())
			if file.Name() == ".git" {
				// fmt.Println(file.Name())
			} else {
				newPath := fmt.Sprintf("%s/%s", dir, file.Name())
				listFiles(newPath) // Recursion already 🤣
			}

		} else {
			if strings.HasSuffix(file.Name(), ".csv") {
				csv := file.Name()
				csvPath := fmt.Sprintf("%s/%s", dir, csv)
				fmt.Println(csvPath)

			} //CSV filter
		}

	}

}

func main() {

	// fmt.Println("Happy")
	listFiles("/home/drew/Projects") //linux
	// listFiles("%USER%\\Documents\\Projects\\goprojects") //windows

}
