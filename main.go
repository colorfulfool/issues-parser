package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	issues := make(map[string][]string)

	var lastHeading string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "## ") {
			cleanHeading := strings.Split(line[3:], "(")[0]
			lastHeading = cleanHeading
			if len(issues[cleanHeading]) > 0 {
				issues[cleanHeading][0] = line
			} else {
				issues[cleanHeading] = append(issues[cleanHeading], line)
			}
		} else if strings.HasPrefix(line, "- [") {
			issues[lastHeading] = append(issues[lastHeading], line)
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	inputScanner := bufio.NewScanner(os.Stdin)
	for inputScanner.Scan() {
		line := inputScanner.Text()
		if strings.HasPrefix(line, "## ") {
			cleanHeading := strings.Split(line[3:], "(")[0]
			lastHeading = cleanHeading
			if len(issues[cleanHeading]) > 0 {
				issues[cleanHeading][0] = line
			} else {
				issues[cleanHeading] = append(issues[cleanHeading], line)
			}
		} else if strings.HasPrefix(line, "- [") {
			found := false
			for _, item := range issues[lastHeading] {
				if item[6:] == line[6:] {
					found = true
					break
				}
			}
			if !found {
				issues[lastHeading] = append(issues[lastHeading], line)
			}
		}
	}

	for _, value := range issues {
		for _, item := range value {
			fmt.Println(item)
		}
		fmt.Println("")
	}
}
