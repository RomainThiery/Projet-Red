package personnage

import "fmt"

func Init() Personnage {
	var nom string

	fmt.Print("Entrez le nom de votre personnage")
	fmt.Scan(&nom)

	p := Personnage{
		Nom:      "Harmony",
		Classe:   "Elfe",
		Niveau:   1,
		PvMax:    100,
		PvActuel: 40,
		Inventaire: []string{
			"potion de vie",
			"potion de vie",
			"potion de vie",
		},

		Skill: []string{"Coup de poing"},
	}
	return p
}
