package personnage

type Personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	PvMax      int
	PvActuel   int
	Inventaire []string
	Skill      []string
}
