package fight

import (
	"Projet-Red/equipment"
	"Projet-Red/monster"
	"Projet-Red/wallet"
	"fmt"
)

func hasItem(inventory []string, itemName string) bool {
	for _, item := range inventory {
		if item == itemName {
			return true
		}
	}
	return false
}

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
		fmt.Printf("\n--- TOUR DE %s (PV: %d/%d | Mana: %d/%d) ---\n", c.Name, c.CurrentHP, c.MaxHP, c.CurrentMana, c.MaxMana)
		fmt.Println("1. Attaquer")
		fmt.Println("2. Inventaire")
		fmt.Print("Choix : ")

		var choice int
		fmt.Scanln(&choice)

		if choice == 1 {
			fmt.Println("\n--- CHOISIS TON SORT ---")
			fmt.Println("1. Coup de poing (8 dégâts | Coût: 10 Mana)")
			fmt.Println("2. Boule de feu (18 dégâts | Coût: 25 Mana - Requiert Livre de Sort)")
			fmt.Println("0. Retour")
			fmt.Print("Choix du sort : ")

			var spellChoice int
			fmt.Scanln(&spellChoice)

			if spellChoice == 0 {
				continue
			}

			if spellChoice == 1 {
				manaCost := 10
				if c.CurrentMana < manaCost {
					fmt.Println("❌ Mana insuffisant !")
					continue
				}

				c.CurrentMana -= manaCost
				damage := 8
				g.CurrentHP -= damage
				if g.CurrentHP < 0 {
					g.CurrentHP = 0
				}

				fmt.Printf("\n🥊 %s donne un Coup de poing et inflige %d dégâts à %s !\n", c.Name, damage, g.Name)
				fmt.Printf("PV de %s : %d / %d\n", g.Name, g.CurrentHP, g.MaxHP)
				break

			} else if spellChoice == 2 {
				if !hasItem(c.Inventory, "Livre de Sort : Boule de feu") {
					fmt.Println("❌ Tu dois posséder le 'Livre de Sort : Boule de feu' dans ton inventaire !")
					continue
				}

				manaCost := 25
				if c.CurrentMana < manaCost {
					fmt.Println("❌ Mana insuffisant pour lancer Boule de feu !")
					continue
				}

				c.CurrentMana -= manaCost
				damage := 18
				g.CurrentHP -= damage
				if g.CurrentHP < 0 {
					g.CurrentHP = 0
				}

				fmt.Printf("\n🔥 %s lance une Boule de feu et inflige %d dégâts à %s !\n", c.Name, damage, g.Name)
				fmt.Printf("PV de %s : %d / %d\n", g.Name, g.CurrentHP, g.MaxHP)
				break
			} else {
				fmt.Println("Sort invalide.")
			}

		} else if choice == 2 {
			if len(c.Inventory) == 0 {
				fmt.Println("Ton inventaire est vide !")
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
					fmt.Printf("\nVous utilisez %s. PV : %d / %d\n", selectedItem, c.CurrentHP, c.MaxHP)
					break

				} else if selectedItem == "Potion de mana" {
					c.CurrentMana += 25
					if c.CurrentMana > c.MaxMana {
						c.CurrentMana = c.MaxMana
					}
					c.Inventory = append(c.Inventory[:itemChoice-1], c.Inventory[itemChoice:]...)
					fmt.Printf("\n🧪 Vous utilisez %s. Mana : %d / %d\n", selectedItem, c.CurrentMana, c.MaxMana)
					break

				} else {
					fmt.Printf("Impossible d'utiliser %s en combat !\n", selectedItem)
				}
			} else {
				fmt.Println("Choix invalide.")
			}
		} else {
			fmt.Println("Option invalide.")
		}
	}
}

func TrainingFight(c *equipment.Character, w *wallet.Wallet) {
	g := monster.InitGoblin()
	turn := 1

	fmt.Printf("\n⚔️ --- COMBAT CONTRE %s --- ⚔️\n", g.Name)

	for c.CurrentHP > 0 && g.CurrentHP > 0 {
		fmt.Printf("\n==================== TOUR %d ====================\n", turn)

		CharTurn(c, &g)

		if g.CurrentHP <= 0 {
			fmt.Printf("\n🎉 Victoire ! Tu as vaincu %s !\n", g.Name)
			w.AddGold(20)
			fmt.Println("Tu as récupérer 20 pièce d'or ! ")
			c.CurrentHP = c.MaxHP
			c.CurrentMana = c.MaxMana
			fmt.Printf("💖 Tes PV et ton Mana ont été entièrement restaurés (%d/%d PV) !\n", c.CurrentHP, c.MaxHP)
			break
		}

		fmt.Println("\n--- Tour du monstre ---")
		GoblinPattern(&g, c, turn)

		if c.CurrentHP <= 0 {
			fmt.Println("\n☠️ Défaite... Tu as été vaincu par le monstre.")
			break
		}

		turn++
	}
}
