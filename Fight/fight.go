package fight

import (
	"Projet-Red/equipment"
	"Projet-Red/monster"
	"fmt"
)

func GoblinPattern(g *monster.Monster, c *equipment.Character, turn int) {
	damage := g.Attack
	if turn%3 == 0 {
		damage = g.Attack * 2
		fmt.Printf("\n⚡ Le %s prépare une attaque puissante ! (Tour %d)\n", g.Name, turn)
	}
	c.CurrentHP -= damage
	if c.CurrentHP < 0 {
		c.CurrentHP = 0
	}
	fmt.Printf("%s inflige à %s %d dégâts\n", g.Name, c.Name, damage)
	fmt.Printf("PV de %s : %d / %d\n", c.Name, c.CurrentHP, c.MaxHP)
}
func CharTurn(c *equipment.Character, g *monster.Monster) {
	for {
		fmt.Println("\n--- TOUR DU JOUEUR ---")
		fmt.Println("1. Attaquer ")
		fmt.Println("2. Inventaire")
		fmt.Print("choix : ")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 {
			damage := 5
			g.CurrentHP -= damage
			if g.CurrentHP < 0 {
				g.CurrentHP = 0
			}
			fmt.Printf("\n%s utilise Attaque basique et inflige %d dégâts à %s\n", c.Name, damage, g.Name)
			fmt.Printf("PV de %s : %d / %d\n", g.Name, g.CurrentHP, g.MaxHP)
			break
		} else if choice == 2 {
			if len(c.Inventory) == 0 {
				fmt.Println("Ton inventaire est vide ! ")
				continue
			}
			fmt.Println("\n--- Inventaire ---")
			for i, item := range c.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			fmt.Println("0. Retour")
			fmt.Print("Choisir un objet à utiliser : ")
			var itemChoice int
			fmt.Scanln(&itemChoice)
			if itemChoice == 0 {
				continue
			}
			if itemChoice > 0 && itemChoice <= len(c.Inventory) {
				selectedItem := c.Inventory[itemChoice-1]
				if selectedItem == "Potion de vie" {
					c.CurrentHP += 50
					if c.CurrentHP > c.MaxHP {
						c.CurrentHP = c.MaxHP
					}
					c.Inventory = append(c.Inventory[:itemChoice-1], c.Inventory[itemChoice:]...)
					fmt.Printf("\nVous utilisez %s.\n", selectedItem)
					fmt.Printf("PV de %s : %d / %d\n", c.Name, c.CurrentHP, c.MaxHP)
					break
				} else {
					fmt.Printf("Impossible d'utiliser %s en combat !\n", selectedItem)
				}
			} else {
				fmt.Printf("Choix invalide.")
			}
		} else {
			fmt.Printf("Option invalide.")
		}
	}
}
func TrainingFight(c *equipment.Character) {
	g := monster.InitGoblin()
	turn := 1
	fmt.Printf("\n⚔️ --- COMBAT CONTRE %s --- ⚔️\n", g.Name)
	for c.CurrentHP > 0 && g.CurrentHP > 0 {
		fmt.Printf("\n==================== TOUR %d ====================\n", turn)
		CharTurn(c, &g)
		if g.CurrentHP <= 0 {
			fmt.Printf("\n🎉 Victoire ! Tu as vaincu %s !\n", g.Name)
			break
		}
		fmt.Println("\n--- Tour du monstre ---")
		GoblinPattern(&g, c, turn)
		if c.CurrentHP <= 0 {
			fmt.Printf("\n☠️ Défaite... Tu as été vaincu par le monstre.")
			break
		}
		turn++
	}
}
