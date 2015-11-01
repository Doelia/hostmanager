package main

import (
	"fmt"
	"strings"
)

func add(domain, ip, hostname string) {
	finalContent := getContentFile()
	finalContent += domain + " " + ip + hostname + "\n"
	setContentFile(finalContent)
}

func rm(domain string) {

}

func view() {
	contentString := getContentFile()
	fmt.Print(contentString)
}

func list() {
	contentString := getContentFile()

	linesSpaces := strings.Split(contentString, "\n")
	linesTabs := strings.Split(contentString, "\t")

	var lines []string
	if len(linesSpaces) < len(linesTabs) {
		lines = linesTabs
	} else {
		lines = linesSpaces
	}

	for _, line := range lines {
		if line != "" {
			if string(line[0]) != "#" {
				fmt.Println(line)
			}
		}
	}
}
