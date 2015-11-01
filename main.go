package main

import (
	"fmt"
	"os"
)

func getNameHostFile() string {
	return "/etc/hosts"
}

func main() {
	arg := "help"
	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}
	fmt.Println(arg)

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
		fmt.Println("list                             Affiche la liste des domaines entrées et triés")
		fmt.Println("view                             Affiche le contenu du fichier host")
		fmt.Println("add domain [ip] [nameserver]     Ajoute (ou remplace) un nouveau domaine")
		fmt.Println("rm domain                        Supprime un domaine")
	}
}
