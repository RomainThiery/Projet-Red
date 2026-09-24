package personnage

import "fmt"

func Dead (p *Personnage) {
	if p.PvActuel <= 0 {
		fmt.Println("Vous etes mort !")
		
		p.PvActuel = p.PvMax / 2
		
		fmt.Println(
			"Vous ressuciiter avec",
			p.PvActuel, 
			"PV.")
	}
}