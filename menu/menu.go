package menu

import (
	"fmt"
	"projet-red/personnage"
)

func Start(p *personnage.Personnage) {
	for {

		fmt.Println("=== MENU ===")
		fmt.Println("1 - Information du personnage")
		fmt.Println("2 - Inventaire")
		fmt.Println("3 - Marchand")
		fmt.Println("4 - Quitter")

		var choix int
		fmt.Print("votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
		 	personnage.DisplayInfo(p)
		case 2:
			personnage.AccessInventory(p)
		case 3:
			personnage.Merchant(p)
		case 4:
			fmt.Println("Au revoir !")
			return
		
		default:
			fmt.Println("Choix invalide")
		}
	}
}
