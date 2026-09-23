package blacksmith

import (
	"Projet-Red/equipment"
	"Projet-Red/wallet"
	"fmt"
)

type Recipe struct {
	EquipmentKey string
	DisplayName  string
	Price        int
	Materials    map[string]int
}

func hasMaterials(inventory []string, required map[string]int) bool {
	counts := make(map[string]int)
	for _, item := range inventory {
		counts[item]++
	}
	for mat, qty := range required {
		if counts[mat] < qty {
			return false
		}
	}
	return true
}
func removeMaterials(inventory []string, required map[string]int) []string {
	toRemove := make(map[string]int)
	for mat, qty := range required {
		toRemove[mat] = qty
	}

	newInventory := []string{}
	for _, item := range inventory {
		if toRemove[item] > 0 {
			toRemove[item]--
		} else {
			newInventory = append(newInventory, item)
		}
	}
	return newInventory
}
func OpenBlacksmithMenu(w *wallet.Wallet, c *equipment.Character) {
	// Définition des recettes avec les matériaux requis
	recipes := map[int]Recipe{
		1: {
			EquipmentKey: "Helmet",
			DisplayName:  "Chapeau de l'aventurier",
			Price:        5,
			Materials:    map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1},
		},
		2: {
			EquipmentKey: "Chestplate",
			DisplayName:  "Tunique de l'aventurier",
			Price:        5,
			Materials:    map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1},
		},
		3: {
			EquipmentKey: "Leggings",
			DisplayName:  "Jambières de l'aventurier",
			Price:        5,
			Materials:    map[string]int{"Peau de Troll": 1, "Cuir de Sanglier": 1},
		},
		4: {
			EquipmentKey: "Boots",
			DisplayName:  "Bottes de l'aventurier",
			Price:        5,
			Materials:    map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1},
		},
	}

	for {
		fmt.Println("\n\033[91m=== 🛠️ FORGERIE ROYALE 🛠️ ===\033[0m")
		w.DisplayBalance()
		fmt.Println("1. Chapeau de l'aventurier (5 pièces + 1 Plume de Corbeau, 1 Cuir de Sanglier)")
		fmt.Println("2. Tunique de l'aventurier (5 pièces + 2 Fourrure de Loup, 1 Peau de Troll)")
		fmt.Println("3. Jambières de l'aventurier (5 pièces + 1 Peau de Troll, 1 Cuir de Sanglier)")
		fmt.Println("4. Bottes de l'aventurier (5 pièces + 1 Fourrure de Loup, 1 Cuir de Sanglier)")
		fmt.Println("5. Sortir")
		fmt.Print("Choisir une fabrication (1-5) : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Le forgeron : « À la prochaine ! »")
			break
		}

		recipe, exists := recipes[choice]
		if !exists {
			fmt.Println("Option invalide.")
			continue
		}
		if w.GoldCoins < recipe.Price {
			fmt.Printf("\n❌ Le forgeron : « Tu n'as pas assez d'or ! Il me faut %d pièces. »\n", recipe.Price)
			continue
		}
		if !hasMaterials(c.Inventory, recipe.Materials) {
			fmt.Println("\n❌ Le forgeron : « Il te manque des ressources pour fabriquer cet équipement ! »")
			fmt.Println("Ressources requises :")
			for mat, qty := range recipe.Materials {
				fmt.Printf(" - %s x%d\n", mat, qty)
			}
			continue
		}
		w.RemoveGold(recipe.Price)
		c.Inventory = removeMaterials(c.Inventory, recipe.Materials)
		c.Inventory = append(c.Inventory, recipe.EquipmentKey)

		fmt.Printf("\n🔨 ✨ Succès ! Le forgeron a fabriqué : %s !\n", recipe.DisplayName)
	}
}
