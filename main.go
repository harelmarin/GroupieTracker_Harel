package main

import (
	"Groupie/routeur"
	InitTemplate "Groupie/templates"
)

func main() {
	InitTemplate.InitTemplate()
	routeur.Serveur()
}
