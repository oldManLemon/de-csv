package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config struct to match your YAML structure
type Config struct {
	HomeFolders []string `yaml:"homeFolders"`
}

func loadConfig() (Config, error) {
	// Hardcoded directory path (replace with your actual path)
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, "Documents", "Projects", "goprojects", "de-csv")
	configFile := filepath.Join(configDir, "config")
	// fmt.Println("Documents Directory:", configDir)
	// fmt.Println("Downloads Directory:", configFile)
	// TODO get a better home for the config file.

	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	//Add homepath
	for i, folder := range config.HomeFolders {
		// fmt.Println(folder)
		newPath := filepath.Join(homeDir, folder)
		// fmt.Println(newPath)
		// fmt.Println(config.HomeFolders[i])
		config.HomeFolders[i] = newPath
	}

	return config, nil

}
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
			_, err = fmt.Fprint(tempFile, seperatorLine)
			if err != nil {
				panic(err)
			}
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
				//TODO Figure a better recovery. Maybe a copy to tmp and then delete if all is good.
				panic(err) //We done messed up now
			}
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
	// listFiles("C:\\Users\\Drew\\Documents\\Projects\\goprojects") //windows
	data, err := loadConfig()
	if err != nil {
		fmt.Println("Error Reading Config: ", err)
		return
	}
	fmt.Println("Data: ", data.HomeFolders)
	for _, folder := range data.HomeFolders {
		listFiles(folder)

	}
}
