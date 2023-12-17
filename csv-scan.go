package main

import (
	"fmt"
	"os"
)

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
			fmt.Println(file.Name())
		}

	}

}

func main() {

	// fmt.Println("Happy")
	listFiles("/home/drew/Projects")

}
