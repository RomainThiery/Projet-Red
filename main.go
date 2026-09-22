package main

import (
	"Projet-Red/blacksmith"
	"Projet-Red/equipment"
	"Projet-Red/wallet"
)

func main() {
	myWallet := wallet.Wallet{GoldCoins: 100}

	player := equipment.Character{
		Name:      "Héros",
		MaxHP:     100,
		CurrentHP: 100,
		Inventory: []string{},
	}
	blacksmith.OpenBlacksmithMenu(&myWallet, &player)
	player.OpenEquipmentMenu()
}
