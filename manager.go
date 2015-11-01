package main

import (
	"fmt"
	"strings"
)

func add(domain, ip, hostname string) {

}

func rm(domain string) {

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
