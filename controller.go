package main

import (
	"fmt"
	"os"
)

func onList() {

}

func onView() {

}

func onAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Veuillez entrer un nom de domaine")
		return
	}

	domain := os.Args[2]

	ip := "127.0.0.1"
	if len(os.Args) >= 4 {
		ip = os.Args[3]
	}

	hostname := ""
	if len(os.Args) >= 5 {
		hostname = os.Args[4]
	}

	fmt.Println("add", domain, ip, hostname)
}

func onRm() {

}
