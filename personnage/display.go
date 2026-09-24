package personnage

import "fmt"

func DisplayInfo(p *Personnage) {
    fmt.Println("=== PERSONNAGE ===")

    fmt.Println("Nom :", p.Nom)
    fmt.Println("Classe :", p.Classe)
    fmt.Println("Niveau :", p.Niveau)
    fmt.Println("PV :", p.PvActuel, "/", p.PvMax)
    
	fmt.Println("=== COMPÉTENCES ===")

    for _, skill := range p.Skill {
        fmt.Println("-", skill)
    }
}
 