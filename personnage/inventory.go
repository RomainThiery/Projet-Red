package personnage

import "fmt"

func AccessInventory(p *Personnage) {
    fmt.Println("=== INVENTAIRE ===")

    for i, objet := range p.Inventaire {
        fmt.Println(i+1, "-", objet)
    }

    fmt.Println("0 - Retour")

    var choix int
    fmt.Print("Votre choix : ")
    fmt.Scan(&choix)

    if choix == 0 {
        return
    }

    if choix < 1 || choix > len(p.Inventaire) {
        fmt.Println("Choix invalide")
        return
    }

    objet := p.Inventaire[choix-1]
    if objet == "potion de vie" {
        TakePot(p)

    } else if objet == "potion de poison" {
        PoisonPot(p)

        p.Inventaire = append(
            p.Inventaire[:choix-1],
            p.Inventaire[choix:]...,

        )

    } else if objet == "Livre de Sort : Boule de Feu" {
       if SpellBook(p) {
				p.Inventaire = append(
				p.Inventaire[:choix-1],
				p.Inventaire[choix:]...,
			)
	   }
    }
}
 