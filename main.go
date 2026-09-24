package main

import (
	"projet-red/menu"
	"projet-red/personnage"
)

func main() {
	p1 := personnage.CharCreation()
	menu.Start(&p1)
}
