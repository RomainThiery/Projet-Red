package fight

import (
	"Projet-Red/equipment"
	"Projet-Red/monster"
	"fmt"
)

func GoblinPattern(g *monster.Monster, c *equipment.Character, turn int) {
	damage := g.Attack
	if turn%3 == 0 
}