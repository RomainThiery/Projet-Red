package menu

import (
	"Projet-Red/blacksmith"
	"Projet-Red/equipment"
	"Projet-Red/fight"
	"Projet-Red/merchant"
	"Projet-Red/wallet"
	"fmt"
)

func StartMainMenu(w *wallet.Wallet, c *equipment.Character) {
	for {
		fmt.Println("\n=== 🏰 MENU PRINCIPAL 🏰 ===")
		fmt.Println("1. Forgerie  2. Marchand  3. Inventaire  4. Quitter")
		fmt.Print("Choix : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			blacksmith.OpenBlacksmithMenu(w, c)
		case 2:
			merchant.OpenMerchantMenu(w, c)
		case 3:
			c.OpenEquipmentMenu()
		case 4:
			fight.TrainingFight(c)
		case 5:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
