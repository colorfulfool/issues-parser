package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type issue struct {
	key   string
	title string
	items []string
}

func addIssue(issues *[]issue, title string, lastKey *string) {
	re := regexp.MustCompile(`\((http.+)#.+\)$`)
	key := re.FindStringSubmatch(title)[1]

	*lastKey = key

	var i = 0
	for ; i < len(*issues); i++ {
		if (*issues)[i].key == key {
			(*issues)[i].title = title
			return
		}
	}
	*issues = append(*issues, issue{title: title, key: key})
}

func addTodo(issues *[]issue, key string, line string) {
	var i int

	var found = false
	for i = 0; i < len(*issues); i++ {
		if (*issues)[i].key == key {
			found = true
			break
		}
	}

	if !found {
		panic("tried to add todo to nonexistent issue")
	}

	var issue = (*issues)[i]
	for j := 0; j < len(issue.items); j++ {
		if issue.items[j][6:] == line[6:] {
			return
		}
	}
	(*issues)[i].items = append(issue.items, line)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var issues []issue

	var lastKey string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "## ") {
			addIssue(&issues, line, &lastKey)
		} else if strings.HasPrefix(line, "- [") {
			addTodo(&issues, lastKey, line)
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	inputScanner := bufio.NewScanner(os.Stdin)
	for inputScanner.Scan() {
		line := inputScanner.Text()
		if strings.HasPrefix(line, "## ") {
			addIssue(&issues, line, &lastKey)
		} else if strings.HasPrefix(line, "- [") {
			addTodo(&issues, lastKey, line)
		}
	}

	for _, issue := range issues {
		fmt.Println(strings.Replace(issue.title, "Redesign: ", "", 1))
		for _, item := range issue.items {
			fmt.Println(item)
		}
		fmt.Println("")
	}
}
