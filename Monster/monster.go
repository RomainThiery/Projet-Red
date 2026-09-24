package monster

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
}

func InitGoblin() Monster {
	return Monster{
		Name:      "Gobelin",
		MaxHP:     40,
		CurrentHP: 40,
		Attack:    5,
	}
}

func InitOgre() Monster {
	return Monster{
		Name:      "Ogre",
		MaxHP:     100,
		CurrentHP: 100,
		Attack:    10,
	}
}

func InitDragon() Monster {
	return Monster{
		Name:      "Dragon",
		MaxHP:     200,
		CurrentHP: 200,
		Attack:    20,
	}
}
