package npc

import (
	"Projet-red-3/characters"
	"Projet-red-3/item"
	"Projet-red-3/ui/text"
	"fmt"
)

var merchantFirstTime = true

func Merchant(p *characters.Character) {
	message := fmt.Sprintf(
		"En poussant la porte, une petite cloche résonne dans la boutique. Des fioles, des livres et des morceaux de créatures sont entassés sur des étagères poussiéreuses. Derrière le comptoir, un marchand te fixe quelques secondes avant de sourire.\n%s : « Bienvenue, aventurier. Ici, tout a une valeur : potions, objets, matériaux… et parfois des choses que tu ferais mieux de ne pas toucher. Si tu trouves quelque chose d'intéressant dans Veyr, ramène-le-moi. Je pourrais être généreux. »",
		text.Color("Armando", "31"),
	)

	if merchantFirstTime {
		text.PrintSlow(message)
		merchantFirstTime = false
	}

	for {
		fmt.Println("\n=== MARCHAND ===")
		fmt.Println("Gold :", p.Gold)
		fmt.Println("1. Potion de vie - 20 Gold")
		fmt.Println("2. Potion de poison - 30 Gold")
		fmt.Println("3. Livre de Sort : Boule de Feu - 25 Gold")
		fmt.Println("4. Fourrure de Loup - 4 Gold")
		fmt.Println("5. Peau de Troll - 7 Gold")
		fmt.Println("6. Cuir de Sanglier - 10 Gold")
		fmt.Println("7. Plume de Corbeau - 10 Gold")
		fmt.Println("8. Potion de réanimation - 60 Gold")
		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		fmt.Print("\033[H\033[2J")

		var object item.Item
		var price int

		switch choice {
		case 1:
			object = item.NewHealthPotion()
			price = 20
		case 2:
			object = item.NewPoisonPotion()
			price = 30
		case 3:
			object = item.NewFireballBook()
			price = 25
		case 4:
			object = item.NewWolfFur()
			price = 4
		case 5:
			object = item.NewTrollSkin()
			price = 7
		case 6:
			object = item.NewBoarLeather()
			price = 10
		case 7:
			object = item.NewCrowFeather()
			price = 10
		case 8:
			if p.Revived {
				fmt.Println("Vous avez déjà utilisé une réanimation pendant cette partie.")
				continue
			}

			hasRevivalPotion := false
			for _, inventoryItem := range p.Inventory {
				if inventoryItem.Name() == "Potion de réanimation" {
					hasRevivalPotion = true
					break
				}
			}

			if hasRevivalPotion {
				fmt.Println("Vous avez déjà une Potion de réanimation.")
				continue
			}

			object = item.NewRevivalPotion()
			price = 60
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if p.Gold < price {
			fmt.Println("Vous n'avez pas assez de Gold !")
			continue
		}

		if len(p.Inventory) >= p.MaxInventory {
			fmt.Println("Votre inventaire est plein !")
			continue
		}

		p.Gold -= price
		p.Inventory = append(p.Inventory, object)

		fmt.Println()
		fmt.Println("Vous avez acheté :", object.Name())
		fmt.Println("L'objet a été ajouté à votre inventaire.")
		fmt.Println("Gold restant :", p.Gold)
	}
}
