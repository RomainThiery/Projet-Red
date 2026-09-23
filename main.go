package main

import (
	"Projet-Red/equipment"
	"Projet-Red/menu"
	"Projet-Red/wallet"
)

func main() {
	myWallet := wallet.Wallet{GoldCoins: 100}
	player := equipment.Character{
		Name:      "Héros",
		MaxHP:     100,
		CurrentHP: 100,
	}
	menu.StartMainMenu(&myWallet, &player)
}
