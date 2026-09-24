package personnage

import "fmt"

func Merchant(p *Personnage) {
    fmt.Println("=== MARCHAND ===")
    fmt.Println("1 - Potion de vie")
    fmt.Println("2 - Potion de poison")
    fmt.Println("3 - Livre de Sort : Boule de Feu")
    fmt.Println("0 - Retour")

    var choix int

    fmt.Print("Votre choix : ")
    fmt.Scan(&choix)

    switch choix {
    case 1:
        AddInventory(p, "potion de vie")

    case 2:
        AddInventory(p, "potion de poison")

    case 3:
        AddInventory(p, "Livre de Sort : Boule de Feu")

    case 0:
        return

    default:
        fmt.Println("Choix invalide")
    }
}
 