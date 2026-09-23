package merchant

import (
	"Projet-Red/equipment"
	"Projet-Red/wallet"
	"fmt"
)

type Item struct {
	Name  string
	Price int
}

func OpenMerchantMenu(w *wallet.Wallet, c *equipment.Character) {
	items := []Item{
		{Name: "Potion de vie", Price: 3},
		{Name: "Potion de mana", Price: 3},
		{Name: "Potion de poison", Price: 6},
		{Name: "Livre de Sort : Boule de feu", Price: 25},
		{Name: "Fourrure de Loup", Price: 4},
		{Name: "Peau de Troll", Price: 7},
		{Name: "Cuir de Sanglier", Price: 3},
		{Name: "Plume de Corbeau", Price: 1},
	}

	for {
		fmt.Println("\n\033[92m=== 🧙‍♂️ BOUTIQUE DU MARCHAND 🧙‍♂️===\033[0m")
		w.DisplayBalance()

		for i, item := range items {
			fmt.Printf("%d. %s (%d pièces)\n", i+1, item.Name, item.Price)
		}
		fmt.Println("0. Sortir")
		fmt.Print("Choisir un article (0-7) : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			fmt.Println("Le marchand : « Bon voyage, aventurier ! »")
			break
		}

		if choice > 0 && choice <= len(items) {
			selectedItem := items[choice-1]

			if w.RemoveGold(selectedItem.Price) {
				c.Inventory = append(c.Inventory, selectedItem.Name)
				fmt.Printf("\nTu as acheté : %s !\n", selectedItem.Name)
			} else {
				fmt.Printf("\nLe marchand : « Tu n'as pas assez d'or ! Il te faut %d pièces pour %s. »\n", selectedItem.Price, selectedItem.Name)
			}
		} else {
			fmt.Println("Option invalide.")
		}
	}
}
