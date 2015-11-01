package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func main() {
	arg := "help"
	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}

	switch arg {
	case "list":
		onList()
	case "view":
		onView()
	case "add":
		onAdd()
	case "rm":
		onRm()
	case "help":
		fmt.Println("list                             Affiche la liste des domaines enregistrés dans le fichier host")
		fmt.Println("view                             Affiche le contenu brut du fichier host")
		fmt.Println("add domain [ip] [nameserver]     Ajoute (ou remplace) un nouveau domaine au fichier host")
		fmt.Println("rm domain                        Supprime un domaine du ficher host")
	}
}
