package personnage

import (
    "fmt"
    "strings"

)

func CharCreation() Personnage {

    var nom string
    var classe string
    var pvMax int

    fmt.Print("Choisis ton nom : ")
    fmt.Scan(&nom)

    nom = strings.ToLower(nom)
    nom = strings.ToUpper(nom[:1]) + nom[1:]

    for {
        fmt.Print("Choisis ta classe (Humain, Elfe, Nain) : ")
        fmt.Scan(&classe)
        
		classe = strings.ToLower(classe)

        switch classe {
        case "humain":
            classe = "Humain"
            pvMax = 100
        case "elfe":
            classe = "Elfe"
            pvMax = 80
        case "nain":
            classe = "Nain"
            pvMax = 120
        default:
            fmt.Println("Classe invalide, recommence.")
            continue
        }
        break
    }

    return Personnage{
        Nom:      nom,
        Classe:   classe,
        Niveau:   1,
        PvMax:    pvMax,
        PvActuel: pvMax / 2,
        Inventaire: []string{
            "potion de vie",
            "potion de vie",
            "potion de vie",
        },
        Skill: []string{"Coup de poing"},
    }
}
 