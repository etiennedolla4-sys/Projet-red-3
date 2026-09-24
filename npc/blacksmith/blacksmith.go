package npc

import (
	"fmt"

	"Projet-red-3/characters"
	"Projet-red-3/item"
)

const equipmentPrice = 20

func Blacksmith(p *characters.Character) {
	for {
		fmt.Println()
		fmt.Println("========== FORGERON ==========")
		fmt.Println("Classe :", p.Class)
		fmt.Println("Gold :", p.Gold)
		fmt.Println()

		equipmentList := GetEquipmentList(p.Class)

		if len(equipmentList) == 0 {
			fmt.Println("Aucun équipement disponible pour cette classe.")
			fmt.Println("0. Retour")
			fmt.Print("Votre choix : ")

			var choice int
			fmt.Scan(&choice)

			if choice == 0 {
				return
			}

			continue
		}

		for i, equipment := range equipmentList {
			fmt.Printf(
				"%d. %s - %d Gold\n",
				i+1,
				equipment.Name(),
				equipmentPrice,
			)
		}

		fmt.Println("0. Retour")
		fmt.Print("Votre choix : ")

		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			return
		}

		if choice < 1 || choice > len(equipmentList) {
			fmt.Println("Choix invalide.")
			continue
		}

		equipment := equipmentList[choice-1]

		if p.Gold < equipmentPrice {
			fmt.Println()
			fmt.Println("Vous n'avez pas assez de Gold !")
			continue
		}

		if len(p.Inventory) >= p.MaxInventory {
			fmt.Println()
			fmt.Println("Votre inventaire est plein !")
			continue
		}

		p.Gold -= equipmentPrice
		p.Inventory = append(p.Inventory, equipment)

		ApplyEquipmentBonus(p, equipment)

		fmt.Println()
		fmt.Println("========== ÉQUIPEMENT ==========")
		fmt.Println("Vous avez obtenu :", equipment.Name())
		fmt.Println("L'équipement a été ajouté à votre inventaire.")

		fmt.Println()
		fmt.Println("========== BONUS ==========")

		if equipment.AttackBonus > 0 {
			fmt.Printf("+%d Attaque\n", equipment.AttackBonus)
		}

		if equipment.DefenseBonus > 0 {
			fmt.Printf("+%d Défense\n", equipment.DefenseBonus)
		}

		if equipment.InitiativeBonus > 0 {
			fmt.Printf("+%d Initiative\n", equipment.InitiativeBonus)
		}

		if equipment.HPBonus > 0 {
			fmt.Printf("+%d PV maximum\n", equipment.HPBonus)
		}

		fmt.Println()
		fmt.Println("Gold restant :", p.Gold)
	}
}

func GetEquipmentList(class string) []item.BasicItem {
	switch class {

	case "Gobelin":
		return []item.BasicItem{
			{
				ItemName:        "Hache du Gobelin",
				ItemDescription: "Une petite hache adaptée aux Gobelins.",
				ItemType:        "Équipement",
				AttackBonus:     8,
			},
			{
				ItemName:        "Armure légère du Gobelin",
				ItemDescription: "Une armure légère qui protège sans ralentir.",
				ItemType:        "Équipement",
				DefenseBonus:    5,
			},
			{
				ItemName:        "Bottes rapides du Gobelin",
				ItemDescription: "Des bottes permettant de se déplacer rapidement.",
				ItemType:        "Équipement",
				InitiativeBonus: 5,
			},
		}

	case "Vampire":
		return []item.BasicItem{
			{
				ItemName:        "Cape du Vampire",
				ItemDescription: "Une cape sombre adaptée aux vampires.",
				ItemType:        "Équipement",
				InitiativeBonus: 4,
				HPBonus:         10,
			},
			{
				ItemName:        "Armure sanguinaire",
				ItemDescription: "Une armure imprégnée d'énergie sanguinaire.",
				ItemType:        "Équipement",
				DefenseBonus:    7,
				HPBonus:         15,
			},
			{
				ItemName:        "Bottes nocturnes",
				ItemDescription: "Des bottes silencieuses adaptées à la nuit.",
				ItemType:        "Équipement",
				InitiativeBonus: 8,
			},
		}

	case "Berserker":
		return []item.BasicItem{
			{
				ItemName:        "Grande Hache du Berserker",
				ItemDescription: "Une énorme hache conçue pour infliger de lourds dégâts.",
				ItemType:        "Équipement",
				AttackBonus:     15,
			},
			{
				ItemName:        "Armure du Berserker",
				ItemDescription: "Une armure lourde protégeant son porteur.",
				ItemType:        "Équipement",
				DefenseBonus:    10,
				HPBonus:         20,
			},
			{
				ItemName:        "Gants de Rage",
				ItemDescription: "Des gants renforçant la puissance des attaques.",
				ItemType:        "Équipement",
				AttackBonus:     10,
				InitiativeBonus: 3,
			},
		}

	default:
		return []item.BasicItem{}
	}
}

func ApplyEquipmentBonus(
	p *characters.Character,
	equipment item.BasicItem,
) {
	p.BaseAttack += equipment.AttackBonus
	p.BaseDefense += equipment.DefenseBonus
	p.Initiative += equipment.InitiativeBonus

	if equipment.HPBonus > 0 {
		p.MaxHP += equipment.HPBonus
		p.HP += equipment.HPBonus
	}
}
