package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	arg := "help"
	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}

	switch arg {
	case "ls":
		onList()
	case "view":
		onView()
	case "add":
		onAdd()
	case "rm":
		onRm()
	case "help":
		fmt.Println("ls                               Affiche la liste des domaines enregistrés dans le fichier host")
		fmt.Println("view                             Affiche le contenu brut du fichier host")
		fmt.Println("add domain [ip] [nameserver]     Ajoute (ou remplace) un nouveau domaine au fichier host")
		fmt.Println("rm domain                        Supprime un domaine du ficher host")
	}
}
