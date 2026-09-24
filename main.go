package main

import (
	"Projet-Red/equipment"
	"Projet-Red/menu"
	"Projet-Red/wallet"
)

func main() {
	myWallet := wallet.Wallet{GoldCoins: 100}

	player := equipment.InitCharacter("Héros")

	menu.StartMainMenu(&myWallet, &player)
}
