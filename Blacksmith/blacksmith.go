package blacksmith

import (
	"Projet-Red/equipment"
	"Projet-Red/wallet"
	"fmt"
)

func OpenBlacksmithMenu(w *wallet.Wallet, c *equipment.Character) {

	for {
		fmt.Println("\n\033[91m=== 🛠️ FORGERIE ROYALE 🛠️ ===\033[0m")
		w.DisplayBalance()
		fmt.Println("1. Casque (5 pièces)")
		fmt.Println("2. Plastron (5 pièces)")
		fmt.Println("3. Jambière (5 pièces)")
		fmt.Println("4. Botte (5 pièces)")
		fmt.Println("5. Sortir")
		fmt.Print("Choisir un équipement (1-5) : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Le forgeron : « À la prochaine ! »")
			break
		}

		var itemKey string
		var displayName string

		switch choice {
		case 1:
			itemKey = "Helmet"
			displayName = "Casque"
		case 2:
			itemKey = "Chestplate"
			displayName = "Plastron"
		case 3:
			itemKey = "Leggings"
			displayName = "Jambière"
		case 4:
			itemKey = "Boots"
			displayName = "Botte"
		default:
			fmt.Println("Option invalide.")
			continue
		}

		if w.RemoveGold(5) {
			c.Inventory = append(c.Inventory, itemKey)
			fmt.Printf("Tu as acheté : %s !\n", displayName)
		} else {
			fmt.Println("Le forgeron : « Tu n'as pas assez d'or, il me faut 5 pièces ! »")
		}
	}
}
