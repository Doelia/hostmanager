package main

import (
	"fmt"
	"strings"
)

func add(domain, ip, hostname string) {
	finalContent := getContentFile()
	finalContent += ip + " " + domain
	if hostname != "" {
		finalContent += " " + hostname
	}
	finalContent += "\n"
	setContentFile(finalContent)
}

func rm(domain string) {
	finalContent := ""
	contentString := getContentFile()
	lines := strings.Split(contentString, "\n")
	for _, line := range lines {
		if line != "" {
			if string(line[0]) != "#" {
				args := getArgumentsOfLine(line)
				thisDomain := args[1]
				if thisDomain != domain {
					finalContent += line + "\n"
				}
			} else {
				finalContent += line + "\n"
			}
		}
	}
	setContentFile(finalContent)
}

func view() {
	contentString := getContentFile()
	fmt.Print(contentString)
}

func list() {
	contentString := getContentFile()
	lines := strings.Split(contentString, "\n")
	for _, line := range lines {
		if line != "" {
			if string(line[0]) != "#" {
				fmt.Println(line)
			}
		}
	}
}
