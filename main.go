package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func addIssue(issues map[string][]string, line string, lastHeading *string) {
	cleanHeading := strings.TrimSpace(strings.Split(line[3:], "(")[0])
	*lastHeading = cleanHeading
	if len(issues[cleanHeading]) > 0 {
		issues[cleanHeading][0] = line
	} else {
		issues[cleanHeading] = append(issues[cleanHeading], line)
	}
}

func hasTodo(issues map[string][]string, key string, line string) bool {
	for _, item := range issues[key] {
		if item[6:] == line[6:] {
			return true
		}
	}
	return false
}

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
			addIssue(issues, line, &lastHeading)
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
			addIssue(issues, line, &lastHeading)
		} else if strings.HasPrefix(line, "- [") {
			if !hasTodo(issues, lastHeading, line) {
				issues[lastHeading] = append(issues[lastHeading], line)
			}
		}
	}

	for _, value := range issues {
		for index, item := range value {
			if index == 0 {
				fmt.Println(strings.Replace(item, "Redesign: ", "", 1))
			} else {
				fmt.Println(item)
			}
		}
		fmt.Println("")
	}
}
