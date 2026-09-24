package personnage

import "fmt"

func TakePot(p *Personnage) {
	for i, objet := range p.Inventaire {
		if objet == "potion de vie" {
			p.PvActuel += 50
			if p.PvActuel > p.PvMax {
				p.PvActuel = p.PvMax
			}
			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			fmt.Println("Vous avez utilisé une potion de vie.")
			fmt.Println("PV :", p.PvActuel, "/", p.PvMax)
			return
		}
	}
	fmt.Println("Vous n'avez plus de potion de vie.")
}
