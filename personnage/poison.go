package personnage

import (
	"fmt"
	"time"
)

func PoisonPot(p *Personnage) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)

		p.PvActuel -= 10

		fmt.Println("PV :", p.PvActuel, "/", p.PvMax)

		Dead(p)
	}
}
