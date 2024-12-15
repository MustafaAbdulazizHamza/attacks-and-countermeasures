package main

import (
	"log"

	goPandora "github.com/MustafaAbdulazizHamza/go-pandora"
)

func main() {

	pandora := goPandora.NewPandoraClient("https://192.168.0.111:8080/", "root", "root", "", "")

	if err := pandora.UpdateUserCredentials("mustafa", "3rgd#$d"); err != nil {
		log.Fatal(err)
	}

	if err := pandora.UpdateUserCredentials("root", "Gx#8dLpM4aB$eK7"); err != nil {
		log.Fatal(err)
	}

}
