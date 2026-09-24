package equipment

import "fmt"

type Equipment struct {
	Head  string
	Torso string
	Legs  string
	Feet  string
}
type Character struct {
	Name        string
	MaxHP       int
	CurrentHP   int
	MaxMana     int
	CurrentMana int
	Equipment   Equipment
	Inventory   []string
	Stage       int
}

func InitCharacter(name string) Character {
	return Character{
		Name:        name,
		MaxHP:       100,
		CurrentHP:   100,
		MaxMana:     100,
		CurrentMana: 100,
		Inventory:   []string{"Potion de vie"},
		Stage:       1,
	}
}

func removeSliceItem(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
func (c *Character) EquipItem(itemName string) {
	itemIndex := -1
	for i, item := range c.Inventory {
		if item == itemName {
			itemIndex = i
			break
		}
	}
	if itemIndex == -1 {
		fmt.Printf("Tu ne possèdes pas %s dans ton inventaire.\n", itemName)
		return
	}
	switch itemName {
	case "Helmet":
		if c.Equipment.Head != "" {
			c.Inventory = append(c.Inventory, c.Equipment.Head)
			c.MaxHP -= 20
			fmt.Printf("Tu retire %s.\n", c.Equipment.Head)
		}
		c.Equipment.Head = itemName
		c.MaxHP += 20
		fmt.Printf("Casque équiper ! (+ 20 PV max)")

	case "Chestplate":
		if c.Equipment.Torso != "" {
			c.Inventory = append(c.Inventory, c.Equipment.Torso)
			c.MaxHP -= 40
			fmt.Printf("Tu retire %s.\n", c.Equipment.Torso)
		}
		c.Equipment.Torso = itemName
		c.MaxHP += 40
		fmt.Printf("Plastron équiper ! (+ 40 PV max)")

	case "Leggings":
		if c.Equipment.Legs != "" {
			c.Inventory = append(c.Inventory, c.Equipment.Legs)
			c.MaxHP -= 25
			fmt.Printf("Tu retire %s.\n", c.Equipment.Legs)
		}
		c.Equipment.Legs = itemName
		c.MaxHP += 25
		fmt.Printf("Pentalon équiper ! (+ 25 PV max)")

	case "Boots":
		if c.Equipment.Feet != "" {
			c.Inventory = append(c.Inventory, c.Equipment.Feet)
			c.MaxHP -= 15
		}
		c.Equipment.Feet = itemName
		c.MaxHP += 15
		fmt.Printf("Botte équiper ! (+ 15 PV max)")

	default:
		fmt.Printf("Cette objet ne peut pas être équiper.")
		return
	}
	c.Inventory = removeSliceItem(c.Inventory, itemIndex)
}

func (c *Character) OpenEquipmentMenu() {
	for {
		fmt.Println("\n=== 🎒 INVENTAIRE ===")
		if len(c.Inventory) == 0 {
			fmt.Println("(Aucun objet dans l'inventaire)")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		}
		fmt.Println("0. Quitter")
		fmt.Print("Quel objet veux-tu équiper ? : ")

		var choice int
		fmt.Scanln(&choice)
		if choice == 0 {
			fmt.Println("Fermeture de l'inventaire.")
			break
		} else if choice > 0 && choice <= len(c.Inventory) {
			selectedItem := c.Inventory[choice-1]
			c.EquipItem(selectedItem)
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}
