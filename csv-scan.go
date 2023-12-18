package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func analyseReplace(csvPath string) {

	// seperators := []string{",", ";"}
	//TODO Build a bigger list of known deliminators
	seperators := []rune{',', ';'} //use rune. https://go.dev/blog/strings bufio streams in bytes

	csv, err := os.Open(csvPath)
	if err != nil {
		fmt.Println(err)
		// return err.Error()
	}

	defer csv.Close()
	scanner := bufio.NewScanner(csv)
	scanner.Scan()
	line := scanner.Text()
	for _, delimiter := range seperators {

		if strings.Contains(line, string(delimiter)) {
			// fmt.Println("Here")
			fmt.Println(string(delimiter))
			// return string(delimiter)
			//add sep=(delimiter)
			//Reset to begginging
			csv.Seek(0, 0)
			seperatorLine := fmt.Sprintf("sep=%s\n", string(delimiter))
			tempFile, err := os.CreateTemp("", "tempCSVFile")
			//Error to create temp file
			if err != nil {
				fmt.Println("Error creating temp file:", err)
				return
			}
			defer tempFile.Close()
			//No longer want read open
			// csv.Close()
			_, _ = fmt.Fprint(tempFile, seperatorLine)
			_, err = io.Copy(tempFile, csv)
			if err != nil {
				panic(err)
			}
			csv.Close()
			tempFile.Close()

			//! DANGER
			//REMOVE THE ORIGINAL FILE AND COPY OUR TEMP FILE
			err = os.Remove(csvPath)
			if err != nil {
				panic(err)
			}
			// Rename and copy to replace original file
			err = os.Rename(tempFile.Name(), csvPath)
			if err != nil {
				panic(err) //We done messed up now
			}

			// -------------------
			//stream in new content followed by the rest of the horse
			// csvEdit, err := os.OpenFile(csvPath, os.O_APPEND, 0600)
			// if err != nil {
			// 	panic(err)
			// }
			// defer csvEdit.Close()
			// _, err = csvEdit.Seek(0, 0)
			// if err != nil {
			// 	panic(err)
			// }
			// _, err = csvEdit.WriteString(seperatorLine)
			// if err != nil {
			// 	panic(err)
			// }
			// csvEdit.Close()

		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
		// return err.Error()
	}
	// return ""

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
	// listFiles("/home/drew/Projects") //linux
	// userHomeDir, err := os.UserHomeDir()
	// if err != nil {
	// 	// return err
	// }
	listFiles("C:\\Users\\Drew\\Documents\\Projects\\goprojects") //windows

}
