package personnage

import "fmt"

func SpellBook(p *Personnage) bool {
   for _, skill := range p.Skill {
       if skill == "Boule de feu" {
           fmt.Println("Tu connais déjà ce sort !")
           return false
       }
   }
  
   p.Skill = append(p.Skill, "Boule de feu")
   
   fmt.Println("Tu as appris Boule de feu !")
   return true
}