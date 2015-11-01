package main

import (
	"io/ioutil"
	"strings"
)

func getNameHostFile() string {
	return "/etc/hosts"
}

func getContentFile() string {
	content, err := ioutil.ReadFile(getNameHostFile())
	check(err)
	contentString := string(content)
	return contentString
}

func setContentFile(content string) {
	contentBytes := []byte(content)
	err := ioutil.WriteFile(getNameHostFile(), contentBytes, 0644)
	check(err)
}

func getArgumentsOfLine(line string) []string {
	argsSpaces := strings.Split(line, " ")
	argsTabs := strings.Split(line, "\t")
	var args []string
	if len(argsSpaces) < len(argsTabs) {
		args = argsTabs
	} else {
		args = argsSpaces
	}
	return args
}
