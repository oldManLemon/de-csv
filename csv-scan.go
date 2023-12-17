package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func analyseReplace(csvPath string) {

	// seperators := []string{",", ";"}
	seperators := []rune{',', ';'} //use rune. https://go.dev/blog/strings bufio streams in bytes

	csv, err := os.Open(csvPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer csv.Close()
	scanner := bufio.NewScanner(csv)
	scanner.Scan()
	line := scanner.Text()
	for _, delimiter := range seperators {

		if strings.Contains(line, string(delimiter)) {
			fmt.Println("Here")
			fmt.Println(delimiter)
		}

	}
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println("Line: ", line)
	// }
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}

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
				analyseReplace(csvPath)

			} //CSV filter
		}

	}

}

func main() {

	// fmt.Println("Happy")
	listFiles("/home/drew/Projects") //linux
	// listFiles("%USER%\\Documents\\Projects\\goprojects") //windows

}
