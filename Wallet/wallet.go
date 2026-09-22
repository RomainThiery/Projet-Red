package wallet

import "fmt"

type Wallet struct {
	GoldCoins int
}

func (w Wallet) DisplayBalance() {
	fmt.Printf("\033[33m Monnaie: %d pièces d'or \033[0m\n", w.GoldCoins)
}

func (w Wallet) AddGold(amount int) {
	if amount > 0 {
		w.GoldCoins += amount
		fmt.Printf("%d pièces d'or ajoutées.\n", amount)
	}
}

func (w *Wallet) RemoveGold(amount int) bool {
	if amount <= 0 {
		return false
	}
	if w.GoldCoins >= amount {
		w.GoldCoins -= amount
		fmt.Printf("%d pièces d'or retirées.\n", amount)
		return true
	}
	fmt.Println("solde insuffisant.")
	return false
}
