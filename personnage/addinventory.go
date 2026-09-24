package personnage

import "fmt"

func AddInventory(p *Personnage, objet string) {
   if len(p.Inventaire) >= 10 {
       fmt.Println("Inventaire plein !")
       return
   }
   
   p.Inventaire = append(p.Inventaire, objet)
   fmt.Println(objet, "ajouté à l'inventaire.")
   
}